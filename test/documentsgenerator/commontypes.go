package documentsgenerator

import "github.com/av-belyakov/placeholder_doc-basedb_bi.zone/interfaces"

type ChMessage struct {
	Type    string
	Message string
}

func NewChMessage() *ChMessage {
	return &ChMessage{}
}

func (chm *ChMessage) GetType() string {
	return chm.Type
}

func (chm *ChMessage) GetMessage() string {
	return chm.Message
}

func (chm *ChMessage) SetType(v string) {
	chm.Type = v
}

func (chm *ChMessage) SetMessage(v string) {
	chm.Message = v
}

type Logging struct {
	ch chan interfaces.Messager
}

func NewLogging() *Logging {
	return &Logging{ch: make(chan interfaces.Messager)}
}

func (l *Logging) SetChan(v chan interfaces.Messager) {
	l.ch = v
}

func (l *Logging) GetChan() <-chan interfaces.Messager {
	return l.ch
}

func (l *Logging) GetChanToData() chan<- interfaces.Messager {
	return l.ch
}

func (l *Logging) Send(msgType, msgData string) {
	msg := NewChMessage()
	msg.SetType(msgType)
	msg.SetMessage(msgData)

	l.ch <- msg
}
