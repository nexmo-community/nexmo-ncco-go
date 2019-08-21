package record

import (
	"ncco/event"
)

type Action struct {
	Format       Format       `json:"format,omitempty"`
	Split        Split        `json:"split,omitempty"`
	Channels     byte         `json:"channels,omitempty"`
	EndOnSilence byte         `json:"endOnSilence,omitempty"`
	EndOnKey     rune         `json:"endOnKey,omitempty"`
	TimeOut      uint16       `json:"timeOut,omitempty"`
	BeepStart    bool         `json:"beepStart,omitempty"`
	EventUrl     string       `json:"eventUrl,omitempty"`
	EventMethod  event.Method `json:"eventMethod,omitempty"`
}

type Format string
type Split string

const (
	CONVERSATION Split = "conversation"
)

const (
	MP3 Format = "mp3"
	WAV Format = "wav"
	OGG Format = "ogg"
)

func (a Action) WithType() interface{} {
	return struct {
		Type string `json:"action"`
		Action
	}{"record", a}
}
