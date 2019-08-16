package requestsdetails

import "time"

// RequestDetails entity represents the details of a request for a short URL. Those details will be used to collect stats
type RequestDetails struct {
	origin    string
	createdAt time.Time
}
