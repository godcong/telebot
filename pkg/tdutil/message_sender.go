package tdutil

import (
	"fmt"

	"github.com/motomototv/telebot/pkg/go-tdlib/client"
)

func MessageSender(ms client.MessageSender, v interface{}) error {
	switch s := ms.(type) {
	case *client.MessageSenderUser:
		*(v).(*client.MessageSenderUser) = *s
	case *client.MessageSenderChat:
		*(v).(*client.MessageSenderChat) = *s
	default:
		return fmt.Errorf("unsupported type: %T", ms.MessageSenderType())
	}
	return nil
}
