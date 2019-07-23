package message

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/google/uuid"
)

// DefaultPoint ...
const DefaultPoint = 0.43

// Recognition ...
func Recognition(message *tgbotapi.Message, id string) (able tgbotapi.Chattable, e error) {
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

	result, e := RunRecognition(context.Background(), fp)
	if e != nil {
		return nil, e
	}
	return tgbotapi.NewMessage(message.Chat.ID, "识别出："+strings.Join(result, ",")), nil
}

func RunRecognition(ctx context.Context, path string) (result []string, e error) {
	args := strings.Split(fmt.Sprintf(GetProperty().Recognition, path), " ")

	cmd := exec.CommandContext(ctx, GetProperty().RecognitionCMD, args...)
	cmd.Env = os.Environ()
	stdout, e := cmd.StdoutPipe()
	if e != nil {
		log.Error(e)
		return
	}
	stderr, e := cmd.StderrPipe()
	if e != nil {
		log.Error(e)
		return
	}

	e = cmd.Start()
	if e != nil {
		log.Error(e)
		return
	}
	for {
		reader := bufio.NewReader(io.MultiReader(stderr, stdout))
		lines, _, e := reader.ReadLine()
		if e != nil || io.EOF == e {
			goto END
		}
		log.With("line", string(lines)).Info("lines")
		ss := strings.Split(string(lines), ",")
		if len(ss) > 1 {
			if ss[1] != "no_persons_found" && ss[1] != "unknown_person" {
				result = append(result, ss[1])
			}
		}
	}
END:
	e = cmd.Wait()
	if e != nil {
		log.Error(e)
	}
	log.With("roles", result).Info("result")
	return
}
