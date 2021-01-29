package filter

import (
	"math/rand"
	"time"

	"github.com/sirupsen/logrus"
)

const (
	HEADS = "heads"
	TAILS = "tails"
)

var coin []string

func init() {
	rand.Seed(time.Now().UnixNano())
	coin = []string{HEADS, TAILS}
}

func Flip() string {

	side := coin[rand.Intn(len(coin))]
	logrus.Info("Flipped the coin and it is ", side)
	return side
}
