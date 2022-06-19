package retriever

type Retriever interface {
	CheckDomainAvailability(domainName string) (DomainAvailability, error)
}

type DomainAvailability struct {
	Status        string `json:"status"`
	DateAvailable string `json:"dateAvailable"`
}
