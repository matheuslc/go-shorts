package shorturl

import (
	"net/url"
	"time"
)

// ShortURL defines the entity that represents a short url
type ShortURL struct {
	original, short url.URL
	createdAt       time.Time
}
