package service

import (
	"errors"
	"net/url"
	"strings"

	"github.com/waldo121/brainrot-index/internal/database"
)

var AllowedDomains = []string{
	"www.youtube.com",
}

func AddContent(source url.URL) error {
	var embedUrl string
	switch source.Host {
	case AllowedDomains[0]:
		videoId, found := strings.CutPrefix(source.String(), "https://www.youtube.com/shorts/")
		if !found {
			return errors.New("only youtube shorts urls are supported")
		}
		embedUrl = "https://www.youtube.com/embed/" + videoId
	default:
		return errors.New("this domain is not allowed")
	}
	_, err := database.InsertContent(embedUrl)
	return err
}
func GetRandomContent() (*ContentSource, error) {
	id, uri, err := database.SelectRandomContent()
	if err != nil {
		return nil, err
	}
	parsedUrl, _ := url.Parse(uri)
	return &ContentSource{
		Id:       id,
		EmbedUrl: *parsedUrl,
	}, nil
}
