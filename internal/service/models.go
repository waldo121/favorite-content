package service

import (
	"net/url"
)

type ContentSource struct {
	EmbedUrl url.URL
	Id       uint64
}
