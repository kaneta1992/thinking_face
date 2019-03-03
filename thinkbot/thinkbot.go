package thinkbot

import (
	"log"
	"math/rand"
	"os/exec"
	"strings"

	"github.com/ChimeraCoder/anaconda"
	"github.com/carlescere/scheduler"
	"github.com/kaneta1992/thinking_face_bot/helper"
)

type ThinkBot struct {
	wrapper *helper.APIWrapper
}

func CreateThinkBot(key, secretKey, token, secretToken string) *ThinkBot {
	return &ThinkBot{helper.CreateAPIWrapper(key, secretKey, token, secretToken)}
}

func updateMediaSubmodule() (bool, error) {
	cmd := exec.Command("git", "pull", "origin", "master")
	cmd.Dir = "./media"
	out, err := cmd.Output()
	if err != nil {
		return false, err
	}
	log.Println("Updated media submodule")
	return string(out) != "Already up to date.\n", nil
}

func commitAndPushMediaSubmodule() error {
	err := exec.Command("git", "add", "./media").Run()
	if err != nil {
		return err
	}
	err = exec.Command("git", "commit", "-m", "Update Media Submodule").Run()
	if err != nil {
		return err
	}
	err = exec.Command("git", "push", "origin", "HEAD").Run()
	if err != nil {
		return err
	}
	log.Println("Commit and pushed media submodule")
	return nil
}

func getRandomMediaPath() string {
	return helper.RandomSelect(helper.DirWalk("./media"))
}

func generateMessage() string {
	return strings.Repeat("🤔", rand.Intn(10)+1)
}

func (bot *ThinkBot) StartReplyBot() {
	go func() {
		for {
			log.Println("start replay bot...")
			log.Println(bot)
			s := bot.wrapper.GetTrackPublicStreamFilter("@thinkbott")
			for t := range s.C {
				switch v := t.(type) {
				case anaconda.Tweet:
					tweet := v
					go func() {
						_, err := bot.wrapper.ReplyWithMedia(generateMessage(), getRandomMediaPath(), tweet)
						helper.CheckIfErrorLog(err)
						log.Println("Done Reply")
					}()
				}
			}
			log.Println("disconected")
		}
	}()
}

func (bot *ThinkBot) StartTweetBot() {
	job := func() {
		log.Printf("Tweet...")
		log.Println(bot)
		updated, err := updateMediaSubmodule()
		helper.CheckIfErrorLog(err)
		if updated {
			err = commitAndPushMediaSubmodule()
			helper.CheckIfErrorLog(err)
			log.Println("Done Tweet")
		}
		_, err = bot.wrapper.TweetWithMedia(generateMessage(), getRandomMediaPath())
		helper.CheckIfErrorLog(err)
	}

	scheduler.Every().Day().At("00:20").Run(job)
	scheduler.Every().Day().At("07:05").Run(job)
	scheduler.Every().Day().At("12:05").Run(job)
	scheduler.Every().Day().At("15:05").Run(job)
	scheduler.Every().Day().At("18:05").Run(job)
	scheduler.Every().Day().At("22:05").Run(job)
}
