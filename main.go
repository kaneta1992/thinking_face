package main

import (
	"log"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/kaneta1992/thinking_face_bot/thinkbot"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	key := os.Getenv("CONSUMER_KEY")
	secretKey := os.Getenv("CONSUMER_SECRET")
	token := os.Getenv("ACCESS_TOKEN")
	secretToken := os.Getenv("ACCESS_TOKEN_SECRET")
	bot := thinkbot.CreateThinkBot(key, secretKey, token, secretToken)
	bot.StartReplyBot()
	bot.StartTweetBot()
	log.Println("start...")
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	log.Println(<-ch)
}
