package service

import (
	"errors"
	"net/url"
	"strings"

	"github.com/waldo121/brainrot-index/internal/database"
)

var AllowedDomains = []string{
	"www.youtube.com",
	"www.facebook.com",
}

func AddContent(source url.URL) error {
	switch source.Host {
	case AllowedDomains[0]:
		videoId, found := strings.CutPrefix(source.String(), "https://www.youtube.com/shorts/")
		if !found {
			return errors.New("only youtube shorts urls are supported")
		}
		_, err := database.InsertContent(
			videoId,
			YoutubeShorts.String(),
		)
		return err
	default:
		return errors.New("this domain is not allowed")
	}

}
func GetRandomContent() (*ContentSource, error) {
	var contentPlatform Platform
	id, videoId, platformStr, err := database.SelectRandomContent()
	if err != nil {
		return nil, err
	}
	err = GetPlatform(platformStr, &contentPlatform)
	if err != nil {
		return nil, err
	}
	return &ContentSource{
		Id:       id,
		videoId:  videoId,
		Platform: contentPlatform,
	}, nil
}
