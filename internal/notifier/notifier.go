package notifier

import "github.com/candy12t/jarvis/internal/model"

type Notifier interface {
	NotifyContents(string, model.Contents) error
	NotifyNoContents(string) error
}

type Notice struct {
	Notifier
	date string
}

func NewClient(notifier Notifier, date string) *Notice {
	return &Notice{
		Notifier: notifier,
		date:     date,
	}
}

func (n *Notice) Notify(contents model.Contents) error {
	if contents == nil {
		return nil
	}
	return n.NotifyContents(contents)
}

func (n *Notice) NotifyContents(contents model.Contents) error {
	return n.Notifier.NotifyContents(n.date, contents)
}

func (n *Notice) NotifyNoContents() error {
	return n.Notifier.NotifyNoContents(n.date)
}
