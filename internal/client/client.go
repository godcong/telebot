package client

import (
	"fmt"
	"path/filepath"

	"github.com/motomototv/telebot/config"
	"github.com/motomototv/telebot/pkg/go-tdlib/client"
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
		ApiId:                  config.Client.APIID,
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
		return nil, fmt.Errorf("NewClient error: %s", err)
	}

	optionValue, err := tdlibClient.GetOption(&client.GetOptionRequest{
		Name: "version",
	})
	if err != nil {
		return nil, fmt.Errorf("GetOption error: %s", err)
	}

	fmt.Printf("TDLib version: %s", optionValue.(*client.OptionValueString).Value)

	me, err := tdlibClient.GetMe()
	if err != nil {
		return nil, fmt.Errorf("GetMe error: %s", err)
	}

	fmt.Printf("Me: %s %s [%s]", me.FirstName, me.LastName, me.Username)
	return &Client{
		Client: tdlibClient,
		config: config,
	}, nil
}

func (client *Client) Run() {

}
