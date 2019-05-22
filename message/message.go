package message

import (
	"context"
	"fmt"
	"github.com/girlvr/yinhe_bot/model"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	shell "github.com/godcong/go-ipfs-restapi"
	"github.com/godcong/go-trait"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
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
var photoHas = make(map[string][]byte)
var hasLocal = false
var log *zap.SugaredLogger

// BootWithGAE ...
func BootWithGAE(path string) {
	log = trait.InitGlobalZapSugar()
	e := LoadProperty(path)
	if e != nil {
		panic(e)
	}
	bot, err := tgbotapi.NewBotAPI(property.Token)
	if err != nil {
		log.Fatal(err)
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "443"
		log.Infof("Defaulting to port %s", port)
	}
	bot.Debug = true

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
		writer.Write([]byte("PONG"))
	})
	go http.ListenAndServeTLS(fmt.Sprintf(":%s", port), "cert.pem", "key.pem", nil)
	InitBoot(bot)
	for update := range updates {
		HookMessage(update)
	}
}

// BootWithUpdate ...
func BootWithUpdate(token string) {
	log = trait.ZapSugar()
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
			usr = append(usr, u.UserName)
		}
		m := fmt.Sprintf("欢迎 %s 加入%s! 请阅读置顶下置顶内容。（本消息将于30秒后销毁）", strings.Join(usr, ","), groupName)
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
				cts = append(cts, tgbotapi.NewMessage(update.Message.Chat.ID, "仅支持私聊@yinhe_bot"))
			}
		case "top", "t":
			video := model.Video{}
			b, e := model.Top(&video)
			if e != nil || !b {
				cts = append(cts, tgbotapi.NewMessage(update.Message.Chat.ID, "没有找到对应资源"))
				break
			}
			photo := tgbotapi.NewPhotoUpload(update.Message.Chat.ID, "")
			e = parseVideoInfo(&photo, &video)
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
			cts = append(cts, tgbotapi.NewMessage(update.Message.Chat.ID, "请认准官方下载地址：xxx"))
		case "help", "h":
			cts = append(cts, tgbotapi.NewMessage(update.Message.Chat.ID, help))
		case "fuck":
			cts = append(cts, tgbotapi.NewMessage(update.Message.Chat.ID, "你想fuck谁？可以私聊我哦（@yinhe_bot)"))
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

func parseVideoInfo(photo *tgbotapi.PhotoConfig, video *model.Video) (err error) {
	photo.Caption = "[" + video.Bangumi + "] " + video.Intro
	if len(video.Role) > 0 {
		photo.Caption += " " + video.Role[0]
	}
	photo.Caption = addLine(photo.Caption)
	fb, e := getFile(video.Poster)
	if e != nil {
		return e
	}
	photo.File = *fb
	hasVideo := false
	if video.VideoGroupList == nil || len(video.VideoGroupList) == 0 {
		photo.Caption += "无片源信息"
		return nil
	}

	for _, value := range video.VideoGroupList {
		if value.Object == nil || len(value.Object) == 0 {
			continue
		}
		if len(value.Object) > 0 {
			hasVideo = true
			photo.Caption += fmt.Sprintf("%s片源:", value.Sharpness)
			photo.Caption = addLine(photo.Caption)
		}
		if objS := len(value.Object); objS == 1 {
			if value.Sliced {
				photo.Caption += url(value.Object[0].Link.Hash) + "/" + value.HLS.M3U8
			} else {
				photo.Caption += url(value.Object[0].Link.Hash)
			}
		} else if objS > 1 {
			for idx, o := range value.Object {
				if idx != 0 {
					photo.Caption += "\n"
				}
				if value.Sliced {
					photo.Caption += "片段" + strconv.FormatInt(int64(idx+1), 10) + ":" + url(o.Link.Hash) + "/" + value.HLS.M3U8
				} else {
					photo.Caption += "片段" + strconv.FormatInt(int64(idx+1), 10) + ":" + url(o.Link.Hash)
				}
			}
		} else {
			// do nothing if size < 1
		}
		photo.Caption = addLine(photo.Caption)
	}
	if !hasVideo {
		photo.Caption += "无片源信息"
	}
	return nil
}

func addLine(s string) string {
	return s + "\n" + "-----------" + "\n"
}

func searchVideo(s string) *model.Video {
	video := &model.Video{}
	if b, err := model.DeepFind(s, video); err != nil || !b {
		return nil
	}
	return video
}

func searchVideoList(limit, start int) (videos []*model.Video, err error) {
	videos = []*model.Video{}
	err = model.DB().OrderBy("visit desc").Limit(limit, start).Find(&videos)
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
