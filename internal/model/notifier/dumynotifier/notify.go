package dumynotifier

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/candy12t/jarvis/internal/model"
	"github.com/candy12t/jarvis/internal/notifier"
)

var _ notifier.Notifier = &Client{}

type Client struct {
	out io.Writer
}

func NewClient() *Client {
	return &Client{
		out: os.Stdout,
	}
}

func (c *Client) NotifyContents(date string, contents model.Contents) error {
	for _, content := range contents {
		text := fmt.Sprintf("%s\n%s\n%s", fmt.Sprintf("%s", date), content.Title, content.ContentURL)
		log.Printf("[notify]: %s\n", strings.Replace(text, "\n", " ", 3))
	}
	return nil
}

func (c *Client) NotifyNoContents(date string) error {
	text := fmt.Sprintf("%s 新着なし", date)
	log.Printf("[notify]: %s\n", text)
	return nil
}
