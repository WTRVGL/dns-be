package retriever

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

var API_BASE = "https://api.dnsbelgium.be/whois/registration/"

func CheckDomainAvailability(domainName string) (DomainAvailability, error) {
	domainIsValid, err := ValidateDomainName(domainName)

	if !domainIsValid {
		return DomainAvailability{}, err
	}

	apiUrl := API_BASE + domainName
	resp, err := http.Get(apiUrl)

	var responseJson []byte
	switch resp.StatusCode {
	case 404:
		return DomainAvailability{
			Status:        "available",
			DateAvailable: "",
		}, nil
	case 200:
		responseJson, _ = ioutil.ReadAll(resp.Body)
		break
	}

	var domainResponse DomainResponse
	json.Unmarshal(responseJson, &domainResponse)

	//status "quarantine" or "inuse"
	return DomainAvailability{
		Status:        strings.ToLower(strings.Split(domainResponse.DomainInfo.Status, ".")[2]),
		DateAvailable: domainResponse.DomainInfo.DateAvailable,
	}, nil
}
