package message

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	shell "github.com/godcong/go-ipfs-restapi"
	"github.com/godcong/go-trait"
	"github.com/yinhevr/seed/model"
	"golang.org/x/xerrors"
)

const help = `输入:
/v 或 /video +番号 查询视频 如：/v ssni-334
/t 或 /top　显示推荐视频
/l 或 /list 显示列表（仅私聊）点击:@yinhe_bot，私聊获取更多信息
/d 或 /down 获取求哈希下载链接（仅私聊）
/r 或 /rule 查看群规（仅私聊）
/h 或 /help 显示帮助`

const rule = ``

const groupName = "银河VR共享总部"
const welcome = ""

// ServerURL ...
const ServerURL = "https://ipfs.io/ipfs/"

// LocalURL ...
const LocalURL = "http://localhost:8080/ipfs/"

// WhiteSpace ...
const WhiteSpace = " "

var bot *tgbotapi.BotAPI
var hasLocal = false
var log = trait.NewZapSugar()
var downloadFileID string

// BootWithGAE ...
func BootWithGAE(path string, port string) {
	e := LoadProperty(path)
	if e != nil {
		panic(e)
	}
	bot, err := tgbotapi.NewBotAPI(property.Token)
	if err != nil {
		log.Fatal(err)
	}
	//port := os.Getenv("PORT")
	if port == "" {
		port = "443"
		log.Infof("Defaulting to port %s", port)
	}
	bot.Debug = true
	response, e := bot.RemoveWebhook()
	if e != nil {
		return
	}
	log.Infof("webhook info:%+v", response)
	log.Infof("Authorized on account %s", bot.Self.UserName)
	_, err = bot.SetWebhook(tgbotapi.NewWebhook(property.Host + property.HookAddress))
	if err != nil {
		log.Fatal(err)
	}
	info, err := bot.GetWebhookInfo()
	if err != nil {
		log.Fatal(err)
	}
	if info.LastErrorDate != 0 {
		log.Infof("Telegram callback failed: %s", info.LastErrorMessage)
	}

	updates := bot.ListenForWebhook("/" + property.HookAddress)
	http.HandleFunc("/ping", func(writer http.ResponseWriter, request *http.Request) {
		log.Info("ping call")
		writer.WriteHeader(http.StatusOK)
		_, e := writer.Write([]byte("PONG"))
		if e != nil {
			log.Error(e)
		}
	})
	go func() {
		e := http.ListenAndServeTLS(fmt.Sprintf(":%s", port), "cert.pem", "key.pem", nil)
		if e != nil {
			log.Error(e)
		}
	}()
	InitBoot(bot)
	ct := make(chan tgbotapi.Chattable, 5)

	go func(c <-chan tgbotapi.Chattable) {
		for {
			select {
			case in := <-c:
				if in == nil {

				}
			}
		}
	}(ct)
	for update := range updates {
		HookMessage(update, ct)
	}
}

// BootWithUpdate ...
func BootWithUpdate(token string) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Infof("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	ct := make(chan tgbotapi.Chattable, 5)

	go func(c <-chan tgbotapi.Chattable) {
		for {
			select {
			case in := <-c:
				if in == nil {
					log.Error("nothing")
					return
				}
				if resp, err := bot.Send(in); err != nil {
					log.Error("send message error:", err)
				} else {
					if resp.Document != nil {
						downloadFileID = resp.Document.FileID
					}
				}
			}
		}
	}(ct)

	for update := range updates {
		HookMessage(update, ct)
	}
}

// InitBoot ...
func InitBoot(botapi *tgbotapi.BotAPI) {
	bot = botapi
}

func getName(user *tgbotapi.User) string {
	if user.UserName != "" {
		return user.UserName
	}
	return user.LastName + "·" + user.FirstName
}

// HookMessage ...
func HookMessage(update tgbotapi.Update, ct chan<- tgbotapi.Chattable) {
	if update.Message == nil {
		return
	}

	log.Info("users:", update.Message.NewChatMembers)
	if update.Message.NewChatMembers != nil {

		var usr []string
		for _, u := range *update.Message.NewChatMembers {
			usr = append(usr, getName(&u))
		}
		m := fmt.Sprintf(property.Welcome, strings.Join(usr, ","), property.GroupName)
		var msg tgbotapi.Message
		var err error
		if msg, err = bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, m)); err != nil {
			log.Error("send message error:", err)
		}
		go func(message tgbotapi.Message) {
			time.Sleep(30 * time.Second)
			response, e := bot.DeleteMessage(tgbotapi.NewDeleteMessage(message.Chat.ID, message.MessageID))
			if e != nil {
				log.Error(e)
			}
			log.Infof("delete resp:%+v", response)
		}(msg)
	}

	if !update.Message.IsCommand() {
		if update.Message.Chat.IsPrivate() {
			ps := update.Message.Photo
			fid := ""
			if ps != nil {
				idx := len(*ps) - 1
				if idx >= 0 {
					fid = (*ps)[idx].FileID

				}
			}

			if fid != "" {
				ct <- tgbotapi.NewMessage(update.Message.Chat.ID, "正在识别,稍后将推送结果!")
				go func() {
					a, e := Recognition(update.Message, fid)
					if e != nil {
						log.Error(e)
						ct <- tgbotapi.NewMessage(update.Message.Chat.ID, "对不起这个妹子长得太有个性,我没认出来!")
					} else {
						ct <- a
					}
				}()
			} else {
				log.Infof("private:%+v", update.Message)
				ct <- tgbotapi.NewMessage(update.Message.Chat.ID, "您好，有什么可以帮您？")
				ct <- tgbotapi.NewMessage(update.Message.Chat.ID, help)
			}
		}
	} else {
		// Create a new MessageConfig. We don't have text yet,
		// so we should leave it empty.
		log.Infof("%+v", update)
		switch update.Message.Command() {
		case "video", "v", "ban", "b":
			for _, vt := range Video(update.Message) {
				ct <- vt
			}
		case "list", "l":
			if update.Message.Chat.IsPrivate() == true {
				for _, vt := range List(update.Message) {
					ct <- vt
				}
			} else {
				bot := fmt.Sprintf("仅支持私聊(%s)", property.BotName)
				ct <- tgbotapi.NewMessage(update.Message.Chat.ID, bot)
			}
		case "top", "t":
			video := model.Video{}
			b, e := model.Top(&video)
			if e != nil || !b {
				ct <- tgbotapi.NewMessage(update.Message.Chat.ID, "没有找到对应资源")
				break
			}
			photo := tgbotapi.NewPhotoUpload(update.Message.Chat.ID, "")
			e = parseVideoInfo(&photo, []*model.Video{&video})
			if e != nil {
				ct <- tgbotapi.NewMessage(update.Message.Chat.ID, "没有找到对应资源")
				break
			}
			ct <- photo
		case "status", "s":
			ct <- tgbotapi.NewMessage(update.Message.Chat.ID, "I'm ok")
		case "close":
			closeMsg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
			closeMsg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
			ct <- closeMsg
		case "down", "d":
			ct <- Download(update.Message)
		case "help", "h":
			ct <- tgbotapi.NewMessage(update.Message.Chat.ID, help)
		case "fuck":
			fuck := fmt.Sprintf("你想fuck谁？可以私聊我哦（%s)", property.BotName)
			ct <- tgbotapi.NewMessage(update.Message.Chat.ID, fuck)
		default:
			return
		}
	}

}

func extInfo(total, episode string, sharpness string) string {
	if sharpness != "" {
		sharpness = "［" + sharpness + "］"
	}

	return episode + sharpness
}

func fixIntroName(intro string, role []string, limit int) string {
	if max := len(role); max > 0 {
		if max > limit {
			max = limit
		}
		for i := 0; i < max; i++ {
			if strings.Index(intro, role[i]) == -1 {
				intro += " " + role[i]
			}
		}
	}
	return intro
}

func parseVideoInfo(photo *tgbotapi.PhotoConfig, videos []*model.Video) (err error) {
	if videos == nil || len(videos) <= 0 {
		return xerrors.New("nil video")
	}
	first := videos[0]

	photo.Caption = fixIntroName(first.Intro, first.Role, 5)
	photo.Caption = addLine(photo.Caption)
	fb, e := getFile(first.PosterHash)
	if e != nil {
		photo.Caption += "番号信息尚未收录"
		return e
	}
	photo.File = *fb
	hasVideo := false
	ban := ""
	for _, video := range videos {
		if video.M3U8Hash == "" && video.SourceHash == "" {
			continue
		}
		hasVideo = true
		if ban == "" {
			ban = video.Bangumi
			photo.Caption += fmt.Sprintf("番号: %s", video.Bangumi)
			photo.Caption = addLine(photo.Caption)
		}

		if ban != video.Bangumi {
			//skip other video with fuzzy query
			continue
		}

		if video.M3U8Hash != "" {
			photo.Caption += fmt.Sprintf("哈希%s:  %s", video.Episode, video.M3U8Hash)
		} else {
			photo.Caption += fmt.Sprintf("哈希%s:  %s", video.Episode, video.SourceHash)
		}
		photo.Caption = addLine(photo.Caption)

	}
	if !hasVideo {
		photo.Caption += "无片源信息"
		return nil
	}
	photo.Caption += "请复制本片【番号】或【哈希】到求哈希APP,即可播放视频"
	//photo.Caption = addLine(photo.Caption)
	return nil
}

func addLine(s string) string {
	return s + "\n" + "-----------" + "\n"
}

func searchVideo(s string) []*model.Video {
	videos := new([]*model.Video)
	if err := model.DeepFind(s, videos); err != nil {
		return nil
	}
	return *videos
}

func searchVideoList(limit, start int) (videos []*model.Video, err error) {
	videos = []*model.Video{}
	err = model.DB().Where("m3u8_hash <> ?", "").OrderBy("visit desc").Limit(limit, start).Find(&videos)
	if err != nil {
		return nil, err
	}
	return videos, nil
}

func getLocalFile(name, path string) (fb *tgbotapi.FileBytes, e error) {
	if name == "" {
		name = "dhash.apk"
	}
	fb = &tgbotapi.FileBytes{
		Name: name,
	}
	file, e := os.Open(path)
	if e != nil {
		return &tgbotapi.FileBytes{}, e
	}
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return &tgbotapi.FileBytes{}, err
	}
	fb.Bytes = bytes
	return fb, nil
}

func getFile(hash string) (fb *tgbotapi.FileBytes, e error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	url := connectURL(hash)
	log.Info("connectURL:", url)
	fb = &tgbotapi.FileBytes{
		Name: hash,
	}

	request, e := http.NewRequest("GET", url, nil)
	if e != nil {
		return &tgbotapi.FileBytes{}, e
	}
	response, e := http.DefaultClient.Do(request.WithContext(ctx))

	if e != nil {
		return &tgbotapi.FileBytes{}, e
	}

	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return &tgbotapi.FileBytes{}, err
	}
	fb.Bytes = bytes
	return fb, nil
}

func url(hash string) string {
	return ServerURL + hash
}

func connectURL(hash string) string {
	if checkLocal() {
		return LocalURL + hash
	}
	return ServerURL + hash
}

func checkLocal() bool {
	if !hasLocal {
		rest := shell.NewShell("localhost:5001")
		if _, err := rest.ID(); err != nil {
			return false
		}
		hasLocal = true
	}
	return hasLocal
}
