package dns_be

var (
	inuse      string
	quarantine string
	available  string
)

type Availability struct {
	Status        string `json:"status"`
	DateAvailable string `json:"dateAvailable"`
}
