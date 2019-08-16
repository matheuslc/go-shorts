package shorturl

import (
	"time"
)

// ShortURL defines the entity that represents a short url
type ShortURL struct {
	id, original, short string
	createdAt           time.Time
}
