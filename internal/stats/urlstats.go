package stats

import "github.com/matheuslc/go-shorts/internal/requestsdetails"

// URLStats defines a entity to represent an URLStats
type URLStats struct {
	hits           int
	requestDetails []*requestsdetails.RequestDetails
}
