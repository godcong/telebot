package message

import (
	"context"
	"github.com/girlvr/yinhe_bot/model"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// ServerURL ...
const ServerURL = "https://ipfs.io/ipfs/"

var bot *tgbotapi.BotAPI
var photoHas = make(map[string][]byte)

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
	hasVideo := false
	// Extract the command from the Message.
	switch update.Message.Command() {
	case "video":
		v := strings.Split(update.Message.Text, " ")
		msg.Text = "video not found"
		if len(v) > 1 {
			video := searchVideo(v[1])
			if video == nil {
				break
			}

			e := parseVideo(&config, video)
			if e != nil {
				break
			}
			hasVideo = true
		}
	case "top":
		video := model.Video{}
		b, e := model.Top(&video)
		msg.Text = "video not found"
		if e == nil && b {
			e := parseVideo(&config, &video)
			if e != nil {
				return
			}
			hasVideo = true
		}
	case "ban":

	case "help":
		msg.Text = "输入 /video #番号# 或者 /top 查询视频."
	case "status":
		msg.Text = "I'm ok."
	default:
		return
	}

	if hasVideo {
		if _, err := bot.Send(config); err != nil {
			logrus.Error(err)
		}
		return
	}
	if _, err := bot.Send(msg); err != nil {
		logrus.Error(err)
	}
}

func parseVideo(cfg *tgbotapi.PhotoConfig, video *model.Video) error {
	cfg.Caption = video.Intro
	cfg.Caption = addLine(cfg.Caption)
	fb, e := getFile(video.Poster)
	if e != nil {
		cfg.Caption += "无片源信息"
		return e
	}
	cfg.File = *fb
	hasVideo := false
	for _, value := range video.VideoGroupList {
		if value.Sharpness != "" {
			cfg.Caption += value.Sharpness + "片源:"
			cfg.Caption = addLine(cfg.Caption)
		}
		count := int64(1)
		for _, o := range value.Object {
			hasVideo = true
			if value.Sliced {
				cfg.Caption += "片段" + strconv.FormatInt(count, 10) + ":" + connectURL(o.Link.Hash) + "/" + value.HLS.M3U8 + "\n"
			} else {
				cfg.Caption += "片段" + strconv.FormatInt(count, 10) + ":" + connectURL(o.Link.Hash) + "\n"
			}

			count++
		}
	}
	if cfg.Caption == "" || !hasVideo {
		cfg.Caption += "无片源信息"
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

func getFile(hash string) (fb *tgbotapi.FileBytes, e error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	url := connectURL(hash)
	logrus.Info("connectURL:", url)
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

func connectURL(hash string) string {
	return ServerURL + hash
}
