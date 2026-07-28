package helpers

import (
	"context"
	"fmt"

	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/interfaces"
	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/internal/databasestorageapi"
)

// *** для счётчика ***
type Counting struct {
	ch chan struct {
		Message string
		Count   int
	}
}

func NewCounting() *Counting {
	return &Counting{ch: make(chan struct {
		Message string
		Count   int
	})}
}

func (c *Counting) SendMessage(msg string, count int) {
	c.ch <- struct {
		Message string
		Count   int
	}{Message: msg, Count: count}
}

func (c *Counting) GetChan() <-chan struct {
	Message string
	Count   int
} {
	return c.ch
}

// *** для сообщений ***
type Message struct {
	Type, Message string
}

func (m *Message) GetType() string {
	return m.Type
}

func (m *Message) SetType(v string) {
	m.Type = v
}

func (m *Message) GetMessage() string {
	return m.Message
}

func (m *Message) SetMessage(v string) {
	m.Message = v
}

// *** для логирования ***
type Logging struct {
	Ch chan interfaces.Messager
}

func NewLogging() *Logging {
	return &Logging{
		Ch: make(chan interfaces.Messager),
	}
}

func (l *Logging) GetChan() <-chan interfaces.Messager {
	return l.Ch
}

func (l *Logging) Send(msgType, msgData string) {
	l.Ch <- &Message{Type: msgType, Message: msgData}
}

func DataBaseConnectionInitialization(
	ctx context.Context,
	logging interfaces.Logger,
	counting interfaces.Counter,
	opts ...databasestorageapi.DatabaseStorageOptions,
) (*databasestorageapi.DatabaseStorage, error) {
	var (
		apiDBS *databasestorageapi.DatabaseStorage
		err    error
	)

	apiDBS, err = databasestorageapi.New(logging, counting, opts...)
	if err != nil {
		return apiDBS, err
	}

	if err := apiDBS.Start(ctx); err != nil {
		return apiDBS, err
	}

	go func(ctx context.Context, ch <-chan databasestorageapi.SettingsChanOutput) {
		for {
			select {
			case <-ctx.Done():
				return

			case msg := <-ch:
				fmt.Println("Chan output message:")
				fmt.Printf("ID:'%s'\nUUID:'%s'\nCommand:'%s'\n", msg.Id, msg.UUID, msg.Command)
				fmt.Printf("Data:'%s'\n", string(msg.Data))

			}
		}
	}(ctx, apiDBS.GetChannelFromModule())

	return apiDBS, err
}
