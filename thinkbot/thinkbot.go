package thinkbot

import (
	"math/rand"
	"strings"

	"github.com/ChimeraCoder/anaconda"
	"github.com/kaneta1992/thinking_face_bot/helper"
)

type ThinkBot struct {
	wrapper *helper.APIWrapper
}

func CreateThinkBot(key, secretKey, token, secretToken string) *ThinkBot {
	return &ThinkBot{helper.CreateAPIWrapper(key, secretKey, token, secretToken)}
}

func getRandomMediaPath() string {
	return helper.RandomSelect(helper.DirWalk("./media"))
}

func generateMessage() string {
	return strings.Repeat("🤔", rand.Intn(10)+1)
}

func (bot *ThinkBot) StartReplyBot() {
	go func() {
		s := bot.wrapper.GetTrackPublicStreamFilter("@thinkbott")
		for t := range s.C {
			switch v := t.(type) {
			case anaconda.Tweet:
				tweet := v
				go func() {
					_, err := bot.wrapper.ReplyWithMedia(generateMessage(), getRandomMediaPath(), tweet)
					helper.CheckIfError(err)
				}()
			}
		}
	}()
}

func (bot *ThinkBot) StartTweetBot() {
	go func() {
		_, err := bot.wrapper.TweetWithMedia(generateMessage(), getRandomMediaPath())
		helper.CheckIfError(err)
	}()
}
