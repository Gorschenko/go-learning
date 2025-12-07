package links

import (
	"math/rand/v2"
	"test/internal/stats"

	"gorm.io/gorm"
)

type Link struct {
	gorm.Model
	Url   string       `json:"url"`
	Hash  string       `json:"hash" gorm:"uniqueIndex"`
	Stats []stats.Stat `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func NewLink(url string) *Link {
	link := &Link{
		Url: url,
	}
	link.GenerateHash()
	return link
}

func (link *Link) GenerateHash() {
	link.Hash = RandStringRuns(6)
}

var letterRunes = []rune("abcdefghijklmnoprstuvwxyz")

func RandStringRuns(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.IntN(len(letterRunes))]
	}
	return string(b)
}
