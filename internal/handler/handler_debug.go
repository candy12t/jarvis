//go:build debug

package handler

import (
	"github.com/candy12t/jarvis/internal/notifier"
	"github.com/candy12t/jarvis/internal/notifier/dumynotifier"
)

func NotifyClient() notifier.Notifier {
	return dumynotifier.NewClient()
}
