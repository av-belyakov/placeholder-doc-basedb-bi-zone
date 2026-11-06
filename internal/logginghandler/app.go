package logginghandler

import (
	"context"
	"log"

	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/constants"
	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/interfaces"
)

// New конструктор обработки логов
func New(writer interfaces.WriterLoggingData) *LoggingChan {
	return &LoggingChan{
		dataWriter:  writer,
		chanLogging: make(chan interfaces.Messager),
	}
}

// Start обработчик и распределитель логов
func (lc *LoggingChan) Start(ctx context.Context) {
	go func() {
		for {
			select {
			case <-ctx.Done():
				log.Printf("%vthe logger module, router is stopped%v\n", constants.Ansi_Dark_Red, constants.Ansi_Reset)

				return

			case msg := <-lc.chanLogging:
				//передача информации по логам в simpleLogger
				//здесь так же может быть вывод в консоль, с индикацией цветов соответствующих
				//определенному типу сообщений но для этого надо включить вывод на stdout
				//в конфигурационном файле
				_ = lc.dataWriter.Write(msg.GetType(), msg.GetMessage())
			}
		}
	}()
}
