package retriever

import (
	"errors"
	"strings"
	"unicode"
)

func ValidateDomainName(domainName string) (bool, error) {

	//Splits up domain name. Valid domain should be a [2]string
	sliced := strings.Split(domainName, ".")
	if len(sliced) > 2 {
		return false, errors.New("invalid domain")
	}
	if sliced[1] != "be" {
		return false, errors.New("not a .be domain")
	}

	//Checks if alphanumerical
	for _, r := range sliced[0] {
		if !unicode.IsLetter(r) {
			return false, errors.New("invalid domain")
		}
	}
	return true, nil
}
