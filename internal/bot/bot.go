package bot

import (
	"context"
	"errors"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

	"github.com/motomototv/telebot/abstract"
	"github.com/motomototv/telebot/config"
	"github.com/motomototv/telebot/database"
	"github.com/motomototv/telebot/database/ent/schema"
	"github.com/motomototv/telebot/internal/message"
	"github.com/motomototv/telebot/log"
)

var ErrConfigNil = errors.New("config is nil")

type bot struct {
	context context.Context
	cancel  context.CancelFunc
	config  *config.Config
	bot     *tgbotapi.BotAPI
	ct      chan tgbotapi.Chattable
	hook    [config.BotModelMax]func() error
	db      *database.DB
}

func (b bot) Context() context.Context {
	return b.context
}

func (b bot) Bot() *tgbotapi.BotAPI {
	return b.bot
}

func (b bot) Config() *config.Config {
	return b.config
}

func (b bot) DB() *database.DB {
	return b.db
}

func (b bot) startHook() error {
	log.Println("start hook model:", b.config.Bot.Model)
	if err := b.hook[b.config.Bot.Model](); err != nil {
		return err
	}
	return nil
}

func (b bot) Run() error {
	if err := b.startHook(); err != nil {
		return err
	}

	//go func() {
	//e := http.ListenAndServeTLS(fmt.Sprintf(":%s", port), "cert.pem", "key.pem", nil)
	//if e != nil {
	//	log.Println("ERROR:", e)
	//}
	//}()

	//var in tgbotapi.Chattable
	//var ok bool
	//for {
	//	select {
	//	case in, ok = <-b.ct:
	//		if !ok || in == nil {
	//			continue
	//		}
	//		if resp, err := b.bot.Send(in); err != nil {
	//			log.Println("ERROR:", "send message:", err)
	//		} else {
	//			if resp.Document != nil {
	//				//downloadFileID = resp.Document.FileID
	//				//todo(feat):update download
	//			}
	//		}
	//	}
	//}
	return nil
}

// NewBot ...
// @Description: create bot instance
// @param *config.Config
// @return Bot
// @return error
func NewBot(cfg *config.Config) (abstract.Bot, error) {
	if cfg == nil {
		panic("config is nil")
	}
	tgbot, err := tgbotapi.NewBotAPI(cfg.Bot.Token)
	if err != nil {
		return nil, err
	}

	tgbot.Debug = cfg.Debug
	log.Debug = cfg.Debug
	//remove pre start webhook
	//response, err := tgbot.RemoveWebhook()
	//if err != nil {
	//	return nil, err
	//}

	//fmt.Printf("Webhook info:%+v\n", response)
	log.Printfln("Authorized on account %s", tgbot.Self.UserName)
	//_, err = tgbot.SetWebhook(tgbotapi.NewWebhook(cfg.Bot.Host + cfg.Bot.HookAddress))
	//if err != nil {
	//	log.Fatal(err)
	//}
	//info, err := tgbot.GetWebhookInfo()
	//if err != nil {
	//	log.Fatal(err)
	//}
	//if info.LastErrorDate != 0 {
	//	fmt.Printf("Telegram callback failed: %s\n", info.LastErrorMessage)
	//}

	//http.HandleFunc("/ping", func(writer http.ResponseWriter, request *http.Request) {
	//	log.Println("Ping test call")
	//	writer.WriteHeader(http.StatusOK)
	//	_, e := writer.Write([]byte("PONG"))
	//	if e != nil {
	//		log.Println("ERROR:", e)
	//	}
	//})

	ctx, cancel := context.WithCancel(context.TODO())
	ibot := &bot{
		config:  cfg,
		context: ctx,
		cancel:  cancel,
		bot:     tgbot,
		ct:      make(chan tgbotapi.Chattable, 5),
	}

	db, err := database.Open(ctx, cfg.Bot.Database, cfg.Debug)
	if err != nil {
		ibot.Close()
		return nil, err
	}
	count, err := db.Message.Query().Count(ctx)
	if err != nil {
		return nil, err
	}
	if count == 0 {
		if err := db.Migrate(ctx); err != nil {
			return nil, err
		}
	}

	ibot.db = db

	ibot.hook = [config.BotModelMax]func() error{
		config.BotModelWebhook: ibot.hookWeb,
		config.BotModelUpdate:  ibot.hookUpdate,
	}

	return ibot, nil
}

func (b bot) hookWeb() error {
	log.Println("get hooks")
	//_, err = b.bot.SetWebhook(tgbotapi.NewWebhookWithCert("https://www.google.com:8443/"+bot.Token, "cert.pem"))
	//if err != nil {
	//	log.Fatal(err)
	//}
	info, err := b.bot.GetWebhookInfo()
	if err != nil {
		//log.Fatal(err)
	}
	if info.LastErrorDate != 0 {
		log.Printfln("Telegram callback failed: %s", info.LastErrorMessage)
	}

	updates := b.bot.ListenForWebhook("/" + b.config.Bot.HookAddress)
	go b.hookMessage(updates)
	return nil
}

func (b bot) hookUpdate() error {
	log.Println("get updates")
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := b.bot.GetUpdatesChan(u)
	if err != nil {
		return err
	}
	go b.hookMessage(updates)
	return nil
}

func (b bot) hookMessage(updates tgbotapi.UpdatesChannel) {
	//defer close(updates)
	log.Println("Start message hook")
	for update := range updates {
		log.Printfln("Received new message:%+v\n", update)
		if update.Message == nil {
			continue
		}

		err := b.switchMessage(update)
		if err != nil {
			log.Println("ERROR:", "message:", err)
		}

	}

	//log.Println("new members:", update.Message.NewChatMembers)
	//
	//if update.Message.NewChatMembers != nil {
	//	var usr []string
	//	for _, u := range *update.Message.NewChatMembers {
	//		usr = append(usr, getName(&u))
	//	}
	//	m := fmt.Sprintf(property.Welcome, strings.Join(usr, ","), property.GroupName)
	//	var msg tgbotapi.Message
	//	var err error
	//	if msg, err = bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, m)); err != nil {
	//		log.Error("send message error:", err)
	//	}
	//	go func(message tgbotapi.Message) {
	//		time.Sleep(30 * time.Second)
	//		response, e := bot.DeleteMessage(tgbotapi.NewDeleteMessage(message.Chat.ID, message.MessageID))
	//		if e != nil {
	//			log.Error(e)
	//		}
	//		log.Infof("delete resp:%+v", response)
	//	}(msg)
	//}
	//
	//if !update.Message.IsCommand() {
	//	if update.Message.Chat.IsPrivate() {
	//		ps := update.Message.Photo
	//		fid := ""
	//		if ps != nil {
	//			idx := len(*ps) - 1
	//			if idx >= 0 {
	//				fid = (*ps)[idx].FileID
	//
	//			}
	//		}
	//
	//		if fid != "" {
	//			//ct <- tgbotapi.NewMessage(update.Message.Chat.ID, "正在识别,稍后将推送结果!")
	//			//go func() {
	//			//	a, e := Recognition(update.Message, fid)
	//			//	if e != nil {
	//			//		log.Error(e)
	//			//		ct <- tgbotapi.NewMessage(update.Message.Chat.ID, "对不起这个妹子长得太有个性,我没认出来!")
	//			//	} else {
	//			//		ct <- a
	//			//	}
	//			//}()
	//		} else {
	//			log.Infof("private:%+v", update.Message)
	//			ct <- tgbotapi.NewMessage(update.Message.Chat.ID, "您好，有什么可以帮您？")
	//			ct <- tgbotapi.NewMessage(update.Message.Chat.ID, help)
	//		}
	//	}
	//} else {
	//	// Create a new MessageConfig. We don't have text yet,
	//	// so we should leave it empty.
	//	log.Infof("%+v", update)
	//	switch update.Message.Command() {
	//	case "video", "v", "ban", "b":
	//		v := strings.Split(update.Message.Text, WhiteSpace)
	//		if len(v) > 1 {
	//			ct <- tgbotapi.NewMessage(update.Message.Chat.ID, "正在搜索："+v[1])
	//			for _, vt := range Video(update.Message, v[1]) {
	//				log.Info("output video info")
	//				ct <- vt
	//			}
	//		}
	//	case "list", "l":
	//		if update.Message.Chat.IsPrivate() == true {
	//			for _, vt := range List(update.Message) {
	//				ct <- vt
	//			}
	//		} else {
	//			bot := fmt.Sprintf("仅支持私聊(%s)", property.BotName)
	//			ct <- tgbotapi.NewMessage(update.Message.Chat.ID, bot)
	//		}
	//	case "top", "t":
	//		session := db.NewSession()
	//		defer session.Close()
	//		videos, e := model.Top(session, 0)
	//		if e != nil {
	//			ct <- tgbotapi.NewMessage(update.Message.Chat.ID, "没有找到对应资源")
	//			break
	//		}
	//		photo := tgbotapi.NewPhotoUpload(update.Message.Chat.ID, "")
	//		e = parseVideoInfo(&photo, *videos)
	//		if e != nil {
	//			ct <- tgbotapi.NewMessage(update.Message.Chat.ID, "没有找到对应资源")
	//			break
	//		}
	//		ct <- photo
	//	case "status", "s":
	//		ct <- tgbotapi.NewMessage(update.Message.Chat.ID, "I'm ok")
	//	case "close":
	//		closeMsg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
	//		closeMsg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
	//		ct <- closeMsg
	//	case "down", "d":
	//		ct <- Download(update.Message)
	//	case "help", "h":
	//		ct <- tgbotapi.NewMessage(update.Message.Chat.ID, help)
	//	case "fuck":
	//		fuck := fmt.Sprintf("你想fuck谁？可以私聊我哦（%s)", property.BotName)
	//		ct <- tgbotapi.NewMessage(update.Message.Chat.ID, fuck)
	//	default:
	//		return
	//	}
	//}
}

func (b bot) switchMessage(update tgbotapi.Update) error {
	var err error
	err = message.Message(b, schema.MessageTypeMessage, update)
	if err != nil {
		return err
	}

	if update.Message.NewChatMembers != nil {
		err = message.Message(b, schema.MessageTypeChatMember, update)
		if err != nil {
			return err
		}
	}

	return nil
}

func (b *bot) Close() {
	if b.cancel != nil {
		b.cancel()
		b.cancel = nil
	}
}

var _ abstract.Bot = (*bot)(nil)
