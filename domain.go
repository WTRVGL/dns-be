package dns_be

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
	"unicode"
)

type domain struct {
	Name         string       `json:"name"`
	Availability Availability `json:"availability"`
}

func NewDomain(name string) (*domain, error) {
	d := &domain{Name: name}
	err := d.validate()

	if err != nil {
		return nil, err
	}

	return d, nil
}

func (d *domain) validate() error {
	//Splits up domain name. Valid domain should be a [2]string
	sliced := strings.Split(d.Name, ".")
	if len(sliced) > 2 {
		return InvalidDomain
	}
	if sliced[1] != "be" {
		return BadTLD
	}

	//Checks if alphanumerical
	for _, r := range sliced[0] {
		isLetterOrNumber := func(r rune) bool {
			isLetter := unicode.IsLetter(r)
			isNumber := unicode.IsNumber(r)
			//Checks for hyphen '-'
			isAcceptedCharacter := r == 45

			if isLetter || isNumber || isAcceptedCharacter {
				return true
			}
			return false
		}
		if !isLetterOrNumber(r) {
			return InvalidDomain
		}
	}
	return nil
}

func (d *domain) CheckAvailability() (*domain, error) {
	apiUrl := "https://api.dnsbelgium.be/whois/registration/" + d.Name
	resp, _ := http.Get(apiUrl)
	var responseJson []byte
	switch resp.StatusCode {
	case 404:
		d.Availability = Availability{Status: "available", DateAvailable: ""}
		return d, nil
	case 200:
		responseJson, _ = ioutil.ReadAll(resp.Body)
		break
	}

	var domainResponse DomainResponse
	_ = json.Unmarshal(responseJson, &domainResponse)

	//status "quarantine" or "inuse"
	d.Availability.Status = strings.ToLower(strings.Split(domainResponse.DomainInfo.Status, ".")[2])
	d.Availability.DateAvailable = domainResponse.DomainInfo.DateAvailable

	return d, nil
}
