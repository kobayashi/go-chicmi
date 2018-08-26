package chicmi

import (
	"net/http"

	"github.com/dghubble/sling"
)

const baseURL = "https://www.chicmi.com/api/"

// Service is endpoint to access chicmi api
type Service struct {
	sling *sling.Sling
}

// APIError is custom api error
type APIError struct {
	Errors     Errors     `json:"errors"`
	ErrorCodes ErrorCodes `json:"error_codes"`
}

// Errors is api errors
type Errors struct {
	Invalid string `json:"invalid"`
}

// ErrorCodes is api error codes
type ErrorCodes struct {
	Invalid string `json:"invalid"`
}

// SearchParams is set as url parameters to get events from chicmi
type SearchParams struct {
	City         string `url:"city,omitempty"`
	Types        string `url:"types,omitempty"`
	Sectors      string `url:"sectors,omitempty"`
	Designers    string `url:"designers,omitempty"`
	Stores       string `url:"stores,omitempty"`
	Users        string `url:"users,omitempty"`
	FeaturedOnly string `url:"featured_only,omitempty"`
	MaxResults   int    `url:"max_results,omitempty"`
	DateFrom     string `url:"date_from,omitempty"`
	DateTo       string `url:"date_to,omitempty"`
	Source       string `url:"source,omitempty"`
}

// EventSearchParams set used as url parameters to get event detail from chicmi
type EventSearchParams struct {
	EventID int `url:"event_id,omitempty"`
}

// NewClient returns a new Client
func NewClient(httpClient *http.Client) *Service {
	return &Service{
		sling: sling.New().Client(httpClient).Base(baseURL),
	}
}

// Error returns api error
func (e APIError) Error() string {
	return e.Errors.Invalid
}
