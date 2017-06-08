package bot

import (
	"github.com/VG-Tech-Dojo/vg-1day-2017/original/model"
)

type (
	// Poster はInに渡されたmessageをPOSTするための構造体です
	Poster struct {
		In chan *model.Message
	}
)

// Run はPosterを起動します
func (p *Poster) Run(url string) {
	for m := range p.In {
		out := &model.Message{}
		postJSON(url+"/api/messages", m, out)
	}
}

// NewPoster は新しいPoster構造体のポインタを返します
func NewPoster(bufferSize int) *Poster {
	in := make(chan *model.Message, bufferSize)
	return &Poster{
		In: in,
	}
}
