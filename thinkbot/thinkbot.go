package thinkbot

import (
	"github.com/ChimeraCoder/anaconda"
	"github.com/kaneta1992/thinking_face_bot/helper"
)

type ThinkBot struct {
	wrapper *helper.APIWrapper
}

func CreateThinkBot(key, secretKey, token, secretToken string) *ThinkBot {
	return &ThinkBot{helper.CreateAPIWrapper(key, secretKey, token, secretToken)}
}

func (bot *ThinkBot) StartReplyBot() {
	go func() {
		s := bot.wrapper.GetTrackPublicStreamFilter("@thinkbott")
		for t := range s.C {
			switch v := t.(type) {
			case anaconda.Tweet:
				tweet := v
				go func() {
					bot.wrapper.ReplyWithMedia("thikning...", "c:/Users/kanet/Downloads/smug_face_anim.gif", tweet)
				}()
			}
		}
	}()
}

func (bot *ThinkBot) StartTweetBot() {
	go func() {
		bot.wrapper.TweetWithMedia("", "c:/Users/kanet/Downloads/thinking_face.mp4")
	}()
}
