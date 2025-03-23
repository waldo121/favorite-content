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

const YOUTUBE_SHORTS_BASE_URL = "https://www.youtube.com/shorts/"
const FACEBOOK_REELS_BASE_URL = "https://www.facebook.com/reel/"

func AddContent(source url.URL) error {
	switch source.Host {
	case AllowedDomains[0]:
		videoId, found := strings.CutPrefix(source.String(), YOUTUBE_SHORTS_BASE_URL)
		if !found {
			return errors.New("only youtube shorts urls are supported")
		}
		_, err := database.InsertContent(
			videoId,
			YoutubeShorts.String(),
		)
		return err
	case AllowedDomains[1]:
		videoId, found := strings.CutPrefix(source.String(), FACEBOOK_REELS_BASE_URL)
		if !found {
			return errors.New("only facebook reel urls are supported")
		}
		_, err := database.InsertContent(
			videoId,
			FacebookReels.String(),
		)
		return err
	default:
		return errors.New("this domain is not allowed")
	}

}
func GetRandomContent() *ContentSource {
	var contentPlatform Platform
	id, videoId, platformStr, err := database.SelectRandomContent()
	if err != nil {
		return &ContentSource{Id: 0, videoId: "", Platform: 0}
	}
	err = GetPlatform(platformStr, &contentPlatform)
	if err != nil {
		return &ContentSource{Id: 0, videoId: "", Platform: 0}
	}
	return &ContentSource{
		Id:       id,
		videoId:  videoId,
		Platform: contentPlatform,
	}
}
