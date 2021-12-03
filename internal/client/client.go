package client

import (
	"fmt"
	"path/filepath"
	"reflect"

	"github.com/motomototv/telebot/config"
	"github.com/motomototv/telebot/log"
	"github.com/motomototv/telebot/pkg/go-tdlib/client"
	"github.com/motomototv/telebot/pkg/tdutil"
)

type Client struct {
	*client.Client
	config *config.Config
}

func NewClient(config *config.Config) (*Client, error) {
	authorizer := client.ClientAuthorizer()
	go client.CliInteractor(authorizer)

	authorizer.TdlibParameters <- &client.TdlibParameters{
		UseTestDc:              false,
		DatabaseDirectory:      filepath.Join(".tdlib", "database"),
		FilesDirectory:         filepath.Join(".tdlib", "files"),
		UseFileDatabase:        true,
		UseChatInfoDatabase:    true,
		UseMessageDatabase:     true,
		UseSecretChats:         false,
		ApiID:                  config.Client.APIID,
		ApiHash:                config.Client.APIHash,
		SystemLanguageCode:     "en",
		DeviceModel:            "Server",
		SystemVersion:          "1.0.0",
		ApplicationVersion:     "1.0.0",
		EnableStorageOptimizer: true,
		IgnoreFileNames:        false,
	}

	logVerbosity := client.WithLogVerbosity(&client.SetLogVerbosityLevelRequest{
		NewVerbosityLevel: 0,
	})

	tdlibClient, err := client.NewClient(authorizer, logVerbosity)
	if err != nil {
		return nil, fmt.Errorf("NewClient error: %s\n", err)
	}

	optionValue, err := tdlibClient.GetOption(&client.GetOptionRequest{
		Name: "version",
	})
	if err != nil {
		return nil, fmt.Errorf("GetOption error: %s\n", err)
	}

	fmt.Printf("TDLib version: %s\n", optionValue.(*client.OptionValueString).Value)

	return &Client{
		Client: tdlibClient,
		config: config,
	}, nil
}

func (c *Client) Me() (*client.User, error) {
	me, err := c.Client.GetMe()
	if err != nil {
		return nil, fmt.Errorf("GetMe error: %s", err)
	}

	log.Printfln("Me: %s %s [%s]", me.FirstName, me.LastName, me.Username)
	return me, nil
}

func (c *Client) GetUserRequest(sender client.MessageSender) (*client.User, error) {
	var id client.MessageSenderUser
	if sender.MessageSenderType() != client.TypeMessageSenderUser {
		return nil, fmt.Errorf("sender type is not user")
	}

	err := tdutil.MessageSender(sender, &id)
	if err != nil {
		return nil, err
	}
	return c.Client.GetUser(&client.GetUserRequest{
		UserID: id.UserID,
	})
}

func (c *Client) ChatMembers(chatid int64) (*client.ChatMembers, error) {
	members, err := c.Client.SearchChatMembers(&client.SearchChatMembersRequest{
		ChatID: chatid,
		Query:  "",
		Limit:  100,
		Filter: &client.ChatMembersFilterMembers{},
	})
	if err != nil {
		return nil, fmt.Errorf("SearchChatMembers error: %s", err)
	}
	log.Printfln("SearchChatMembers: %#v", members)
	return members, nil
}

func (c *Client) Run() {
	listener := c.Client.GetListener()
	defer listener.Close()
	for update := range listener.Updates {
		if update.GetClass() == client.ClassUpdate {
			switch update.GetType() {
			case client.TypeUpdateChatLastMessage:
				v, ok := update.(*client.UpdateChatLastMessage)
				if !ok {
					continue
				}
				log.Printfln("ChatID:%#v", v.ChatID)
				if v.LastMessage == nil {
					continue
				}
				json, err := v.LastMessage.MarshalJSON()
				if err != nil {
					log.Printfln("MarshalJSON error: %s", err)
				}
				log.Printfln("Message:%s", string(json))
				if v.LastMessage.Content == nil {
					continue
				}

				log.Printfln("MessageContentType:%#v", v.LastMessage.Content.MessageContentType())
				log.Printfln("Content:%#v", v.LastMessage.Content)
				processMessage(v.LastMessage)
			}

		}
	}
}

func processMessage(msg *client.Message) {
	switch msg.Content.MessageContentType() {
	case client.TypeMessageText:
		v, ok := msg.Content.(*client.MessageText)
		if !ok {
			log.Printfln("MessageText error:", reflect.TypeOf(msg.Content))
			return
		}
		if v.Text != nil {
			fmt.Println("MessageText:", v.Text.Text)
		}
		//log.Printfln("MessageText:%s", v.Text.Text)
	}
}
