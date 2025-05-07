package link

import (
	"go/adv-api/internal/stat"
	"math/rand"

	"gorm.io/gorm"
)

const randomaizer int = 6

type Link struct {
	gorm.Model
	Url   string      `json:"url"`
	Hash  string      `json:"hash" gorm:"uniqueIndex"`
	Stats []stat.Stat `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}

func NewLink(url string) *Link {
	link := &Link{
		Url: url,
	}
	link.NewHash()
	return link
}

var letterRunes = []rune("qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM")

func (link *Link) NewHash() {
	link.Hash = RandStringRunes(randomaizer)
}
func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
