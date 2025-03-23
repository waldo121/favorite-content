package service

import (
	"errors"
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

func (s *ContentSource) Embed() string {
	switch s.Platform {
	case YoutubeShorts:
		return `<iframe width="100%" height="100%" src="https://www.youtube.com/embed/` + s.videoId + `"></iframe>`
	case FacebookReels:
		return `<div id="fb-root"></div><script async defer src="https://connect.facebook.net/fr_FR/sdk.js#xfbml=1&version=v3.2"></script><div class="fb-video" data-href="https://www.facebook.com/reel/` + s.videoId + `" data-width="500" data-allow-fullscreen="true"></div>`
	default:
		return ""
	}
}
