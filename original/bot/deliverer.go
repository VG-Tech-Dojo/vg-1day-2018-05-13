package bot

import (
	"github.com/VG-Tech-Dojo/vg-1day-2017/original/model"
)

type (
	// Deliverer は1つのチャンネルで複数botを動かすためのヘルパーです
	//
	// msgInで受け取ったmessageをbotsに登録された全botに渡します
	//
	// botsへの登録はBotInで行います
	//
	//   fields
	// 	   BotIn chan *Bot
	// 	   bots  map[*Bot]bool
	// 	   msgIn chan *model.Message
	Deliverer struct {
		BotIn chan *Bot
		bots  []*Bot
		msgIn chan *model.Message
	}
)

// Run はDelivererを起動します
func (b *Deliverer) Run() {
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

// NewDeliverer は新しいDeliverer構造体のポインタを返します
func NewDeliverer(msgIn chan *model.Message) *Deliverer {
	memberIn := make(chan *Bot)
	return &Deliverer{
		BotIn: memberIn,
		bots:  []*Bot{},
		msgIn: msgIn,
	}
}
