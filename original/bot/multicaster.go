package bot

import (
	"github.com/VG-Tech-Dojo/vg-1day-2017/original/model"
)

type (
	// Multicaster は1つのチャンネルで複数botを動かすためのヘルパーです
	//
	// msgInで受け取ったmessageをbotsに登録された全botに渡します
	//
	// botsへの登録はBotInで行います
	//
	//   fields
	// 	   BotIn chan *Bot
	// 	   bots  map[*Bot]bool
	// 	   msgIn chan *model.Message
	Multicaster struct {
		BotIn chan *Bot
		bots  []*Bot
		msgIn chan *model.Message
	}
)

// Run はMulticasterを起動します
func (b *Multicaster) Run() {
	for {
		select {
		case bot := <-b.BotIn:
			b.bots = append(b.bots, bot)
		case msg := <-b.msgIn:
			for _, bot := range b.bots {
				bot.in <- msg
			}
		}
	}
}

// NewMulticaster は新しいMulticaster構造体のポインタを返します
func NewMulticaster(msgIn chan *model.Message) *Multicaster {
	memberIn := make(chan *Bot)
	return &Multicaster{
		BotIn: memberIn,
		bots:  []*Bot{},
		msgIn: msgIn,
	}
}
