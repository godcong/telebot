package message

import (
	"context"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	shell "github.com/godcong/go-ipfs-restapi"
	"github.com/godcong/go-trait"
	"github.com/yinhevr/seed/model"
	"golang.org/x/xerrors"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
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
	for update := range updates {
		HookMessage(update)
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

	for update := range updates {
		HookMessage(update)
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
func HookMessage(update tgbotapi.Update) {
	if update.Message == nil {
		return
	}

	var cts []tgbotapi.Chattable
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
			log.Info("private", update.Message)
			cts = append(cts, tgbotapi.NewMessage(update.Message.Chat.ID, "您好，有什么可以帮您？"))
			cts = append(cts, tgbotapi.NewMessage(update.Message.Chat.ID, help))
		}
	} else {
		// Create a new MessageConfig. We don't have text yet,
		// so we should leave it empty.
		log.Infof("%+v", update)
		switch update.Message.Command() {
		case "video", "v", "ban", "b":
			cts = Video(update.Message)
		case "list", "l":
			if update.Message.Chat.IsPrivate() == true {
				cts = List(update.Message)
			} else {
				bot := fmt.Sprintf("仅支持私聊(%s)", property.BotName)
				cts = append(cts, tgbotapi.NewMessage(update.Message.Chat.ID, bot))
			}
		case "top", "t":
			video := model.Video{}
			b, e := model.Top(&video)
			if e != nil || !b {
				cts = append(cts, tgbotapi.NewMessage(update.Message.Chat.ID, "没有找到对应资源"))
				break
			}
			photo := tgbotapi.NewPhotoUpload(update.Message.Chat.ID, "")
			e = parseVideoInfo(&photo, []*model.Video{&video})
			if e != nil {
				cts = append(cts, tgbotapi.NewMessage(update.Message.Chat.ID, "没有找到对应资源"))
				break
			}
			cts = append(cts, photo)
		case "status", "s":
			cts = append(cts, tgbotapi.NewMessage(update.Message.Chat.ID, "I'm ok"))
		case "close":
			closeMsg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
			closeMsg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
			cts = append(cts, closeMsg)
		case "down", "d":
			down := fmt.Sprintf("请认准官方下载地址:%s", property.Download)
			cts = append(cts, tgbotapi.NewMessage(update.Message.Chat.ID, down))
		case "help", "h":
			cts = append(cts, tgbotapi.NewMessage(update.Message.Chat.ID, help))
		case "fuck":
			fuck := fmt.Sprintf("你想fuck谁？可以私聊我哦（%s)", property.BotName)
			cts = append(cts, tgbotapi.NewMessage(update.Message.Chat.ID, fuck))
		default:
			return
		}
	}

	for _, ct := range cts {
		if _, err := bot.Send(ct); err != nil {
			log.Error("send message error:", err)
		}
	}
}

func extInfo(total, episode string, sharpness string) string {
	if sharpness != "" {
		sharpness = "［" + sharpness + "］"
	}

	return episode + sharpness
}

func parseVideoInfo(photo *tgbotapi.PhotoConfig, videos []*model.Video) (err error) {
	if videos == nil || len(videos) <= 0 {
		return xerrors.New("nil video")
	}
	first := videos[0]

	photo.Caption = first.Intro
	if max := len(first.Role); max > 0 {
		if max > 5 {
			max = 5
		}
		for i := 0; i < max; i++ {
			photo.Caption += " " + first.Role[i]
		}

	}
	photo.Caption = addLine(photo.Caption)
	fb, e := getFile(first.PosterHash)
	if e != nil {
		return e
	}
	photo.File = *fb
	hasVideo := false

	for i, video := range videos {
		if video.M3U8Hash == "" && video.SourceHash == "" {
			continue
		}
		hasVideo = true
		if i == 0 {
			photo.Caption += fmt.Sprintf("番号: %s", video.Bangumi)
			photo.Caption = addLine(photo.Caption)
		}

		if video.M3U8Hash != "" {
			photo.Caption += fmt.Sprintf("哈希%s:  %s", video.Episode, video.SourceHash)
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
	photo.Caption = addLine(photo.Caption)
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
