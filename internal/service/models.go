package service

import (
	"errors"
	"fmt"
)

type VideoSource interface {
	Embed() (string, error)
}

type ContentSource struct {
	Id       uint64
	videoId  string
	Platform Platform
}
type Platform int

const (
	YoutubeShorts Platform = iota
	FacebookReels
	Tiktok
)

func GetPlatform(s string, p *Platform) error {
	switch s {
	case "YoutubeShorts":
		*p = YoutubeShorts
		return nil
	case "FacebookReels":
		*p = FacebookReels
		return nil
	case "Tiktok":
		*p = Tiktok
		return nil
	default:
		return errors.New("invalid platform")
	}
}

func (p Platform) String() string {
	return [...]string{"YoutubeShorts", "FacebookReels", "Tiktok"}[p]
}

func (s *ContentSource) Embed() (string, error) {
	fmt.Println(s)
	switch s.Platform {
	case YoutubeShorts:
		return `<iframe width="100%" height="100%" src="https://www.youtube.com/embed/` + s.videoId + `"></iframe>`, nil
	default:
		return "", errors.New("invalid platform")
	}
}
