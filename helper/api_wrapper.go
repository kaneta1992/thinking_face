package helper

import (
	"fmt"
	"net/url"
	"path/filepath"

	"github.com/ChimeraCoder/anaconda"
)

type APIWrapper struct {
	client *anaconda.TwitterApi
}

func CreateAPIWrapper(key, secretKey, token, secretToken string) *APIWrapper {
	anaconda.SetConsumerKey(key)
	anaconda.SetConsumerSecret(secretKey)
	api := anaconda.NewTwitterApi(token, secretToken)
	return &APIWrapper{api}
}

func (w *APIWrapper) uploadImage(path string) (string, error) {
	encoded, err := ImageToBase64(path)
	if err != nil {
		return "", err
	}

	media, err := w.client.UploadMedia(encoded)
	if err != nil {
		return "", err
	}

	return media.MediaIDString, nil
}

func (w *APIWrapper) uploadVideo(path string) (string, error) {
	encodedSegments, bytes, err := VideoToBase64Segments(path)
	if err != nil {
		return "", err
	}

	media, err := w.client.UploadVideoInit(bytes, "video/mp4")
	if err != nil {
		return "", err
	}

	for i, s := range encodedSegments {
		if err = w.client.UploadVideoAppend(media.MediaIDString, i, s); err != nil {
			return "", err
		}
	}

	video, err := w.client.UploadVideoFinalize(media.MediaIDString)
	if err != nil {
		return "", err
	}

	return video.MediaIDString, nil
}

func (w *APIWrapper) uploadMedia(path string) (string, error) {
	ext := filepath.Ext(path)
	switch ext {
	case ".gif", ".jpeg", ".jpg", ".png":
		return w.uploadImage(path)
	case ".mp4":
		return w.uploadVideo(path)
	}
	return "", nil
}

func (w *APIWrapper) TweetWithMedia(message string, path string) (anaconda.Tweet, error) {
	v := url.Values{}
	id, err := w.uploadMedia(path)
	if err != nil {
		return anaconda.Tweet{}, err
	}
	v.Add("media_ids", id)
	return w.client.PostTweet(message, v)
}

func (w *APIWrapper) ReplyWithMedia(message string, path string, tweet anaconda.Tweet) (anaconda.Tweet, error) {
	v := url.Values{}
	id, err := w.uploadMedia(path)
	if err != nil {
		return anaconda.Tweet{}, err
	}
	v.Add("media_ids", id)
	v.Add("in_reply_to_status_id", tweet.IdStr)

	// debug
	fmt.Printf("%-15s: %s\n", tweet.User.ScreenName, tweet.Text)
	fmt.Printf("%s\n", tweet.IdStr)

	return w.client.PostTweet("@"+tweet.User.ScreenName+" "+message, v)
}

func (w *APIWrapper) GetTrackPublicStreamFilter(track string) *anaconda.Stream {
	v := url.Values{}
	v.Set("track", track)
	return w.client.PublicStreamFilter(v)
}
