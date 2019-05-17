package message

import (
	"context"
	"fmt"
	"github.com/girlvr/yinhe_bot/model"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	shell "github.com/godcong/go-ipfs-restapi"
	"github.com/godcong/go-trait"
	"go.uber.org/zap"
	"golang.org/x/xerrors"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"
)

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
func BootWithGAE(token string) {
	log = trait.ZapSugar()
	bot, err := tgbotapi.NewBotAPI(token)
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
	t := "crVuYHQbUWCerib3"
	_, err = bot.SetWebhook(tgbotapi.NewWebhook("https://bot.dhash.app/" + t))
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

	updates := bot.ListenForWebhook("/" + t)
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

	if !update.Message.IsCommand() { // ignore any non-command Messages
		return
	}
	// Create a new MessageConfig. We don't have text yet,
	// so we should leave it empty.
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
	config := tgbotapi.NewPhotoUpload(update.Message.Chat.ID, "")
	var cts []tgbotapi.Chattable
	log.Infof("%+v", update)
	switch update.Message.Command() {
	case "video", "v", "ban", "b":
		cts = Video(update.Message)
	case "list", "l":
		if update.Message.Chat.IsPrivate() == true {
			cts = List(update.Message)
		} else {
			cts = append(cts, tgbotapi.NewMessage(update.Message.Chat.ID, "仅支持私聊"))
		}
	case "top", "t":
		video := model.Video{}
		b, e := model.Top(&video)
		if e != nil || !b {
			msg.Text = "没有找到对应资源"
			break
		}
		e = parseVideoInfo(&config, &video)
		if e != nil {
			msg.Text = "没有找到对应资源"
			break
		}
	case "help", "h":
		msg.Text = "输入 /v 或 /video +番号 查询视频 如：/v ssni-334 或者 /top　显示推荐视频."
	case "status", "s":
		msg.Text = "I'm ok."
	case "close":
		closeMsg := tgbotapi.NewMessage(message.Chat.ID, "")
		closeMsg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
		cts = append(cts, closeMsg)
	default:
		return
	}
	for _, ct := range cts {
		if _, err := bot.Send(ct); err != nil {
			log.Error(err)
		}
	}
}

func parseVideoInfo(photo *tgbotapi.PhotoConfig, video *model.Video) (err error) {
	photo.Caption = video.Intro
	if len(video.Role) > 0 {
		photo.Caption += video.Role[0]
	}
	photo.Caption = addLine(photo.Caption)
	fb, e := getFile(video.Poster)
	if e != nil {
		return e
	}
	photo.File = *fb
	hasVideo := false
	for _, value := range video.VideoGroupList {
		if value.Sharpness != "" {
			photo.Caption += value.Sharpness + "片源:"
			photo.Caption = addLine(photo.Caption)
		}
		count := int64(1)
		for _, o := range value.Object {
			hasVideo = true
			if value.Sliced {
				photo.Caption += "片段" + strconv.FormatInt(count, 10) + ":" + url(o.Link.Hash) + "/" + value.HLS.M3U8 + "\n"
			} else {
				photo.Caption += "片段" + strconv.FormatInt(count, 10) + ":" + url(o.Link.Hash) + "\n"
			}
			count++
		}
	}
	if photo.Caption == "" || !hasVideo {
		return xerrors.New("无片源信息")
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
