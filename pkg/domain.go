package pkg

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
	"unicode"

	"github.com/WTRVGL/dns-be/internal/entities"
	"github.com/WTRVGL/dns-be/internal/entities/errors"
)

type Domain struct {
	Name         string
	Availability entities.Availability
}

func NewDomain(name string) (*Domain, error) {
	d := &Domain{Name: name}
	err := d.validate()

	if err != nil {
		return &Domain{}, err
	}

	return d, nil
}

func (d *Domain) validate() error {
	//Splits up domain name. Valid domain should be a [2]string
	sliced := strings.Split(d.Name, ".")
	if len(sliced) > 2 {
		return errors.InvalidDomain
	}
	if sliced[1] != "be" {
		return errors.BadTLD
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
			return errors.InvalidDomain
		}
	}
	return nil
}

func (d *Domain) CheckAvailability() (Domain, error) {
	apiUrl := "https://api.dnsbelgium.be/whois/registration/" + d.Name
	resp, _ := http.Get(apiUrl)
	var responseJson []byte
	switch resp.StatusCode {
	case 404:
		d.Availability = entities.Availability{Status: "Available", DateAvailable: ""}
		return *d, nil
	case 200:
		responseJson, _ = ioutil.ReadAll(resp.Body)
		break
	}

	var domainResponse entities.DomainResponse
	json.Unmarshal(responseJson, &domainResponse)

	//status "quarantine" or "inuse"
	d.Availability.Status = strings.ToLower(strings.Split(domainResponse.DomainInfo.Status, ".")[2])
	d.Availability.DateAvailable = domainResponse.DomainInfo.DateAvailable

	return *d, nil
}
