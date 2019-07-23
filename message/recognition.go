package message

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/google/uuid"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

// DefaultPoint ...
const DefaultPoint = 0.43

// Recognition ...
func Recognition(id string) (able tgbotapi.Chattable, e error) {
	s, e := bot.GetFileDirectURL(id)
	if e != nil {
		log.Error(e)
	}
	log.Infof("%s:(%s)", id, s)
	resp, e := http.Get(s)
	if e != nil {
		return nil, e
	}

	ext := filepath.Ext(s)

	fp := filepath.Join(time.Now().Format("20060102"), uuid.New().String())
	fp, e = filepath.Abs(fp)
	if e != nil {
		return nil, e
	}
	_ = os.MkdirAll(fp, os.ModePerm)
	newfile := filepath.Join(fp, "unknown"+ext)
	log.With("path", newfile).Info("new file")
	file, e := os.OpenFile(newfile, os.O_CREATE|os.O_SYNC|os.O_RDWR, os.ModePerm)
	if e != nil {
		return nil, e
	}
	written, e := io.Copy(file, resp.Body)
	if e != nil {
		return nil, e
	}
	log.With("size", written).Info("picture written")

	if property.Point <= 0 {
		property.Point = DefaultPoint
	}

	//c := fmt.Sprintf(GetProperty().Recognition,, fp)

	exec.Command("")

	return
}
