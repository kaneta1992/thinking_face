package helper

import (
	"encoding/base64"
	"io/ioutil"
	"os"
)

func ImageToBase64(path string) (string, error) {
	imageFile, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer imageFile.Close()

	binary, err := ioutil.ReadAll(imageFile)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(binary), nil
}

func VideoToBase64Segments(path string) ([]string, int, error) {
	videoFile, err := os.Open(path)
	if err != nil {
		return nil, -1, err
	}
	defer videoFile.Close()

	binary, err := ioutil.ReadAll(videoFile)
	if err != nil {
		return nil, -1, err
	}

	ret := []string{}

	// The amount that can be sent at one time with media/upload is 5mb
	// https://developer.twitter.com/en/docs/media/upload-media/api-reference/post-media-upload-append.html
	mediaSegments := DividedByteSlice(binary, 5*1024*1024)
	for _, segment := range mediaSegments {
		ret = append(ret, base64.StdEncoding.EncodeToString(segment))
	}

	return ret, len(binary), nil
}
