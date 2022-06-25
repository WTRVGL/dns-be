package test

import (
	checker "github.com/WTRVGL/dns-be"
	"testing"
)

func TestAvailableDomain(t *testing.T) {
	domainName := "ifthisgetsregisteredthenripunittest.be"
	domain, _ := checker.NewDomain(domainName)
	domain, _ = domain.CheckAvailability()
	wantedAvailability := checker.Availability{Status: "available", DateAvailable: ""}
	if domain.Availability != wantedAvailability {
		t.Errorf("CheckAvailability(\"ifthisgetsregisteredthenripunittest.be\"), got %s, want %s", domain.Availability, wantedAvailability)
	}
}

func TestUnavailableDomain(t *testing.T) {
	domainName := "dnsbelgium.be"
	domain, _ := checker.NewDomain(domainName)
	domain, _ = domain.CheckAvailability()
	wantedAvailability := checker.Availability{Status: "inuse", DateAvailable: ""}
	if domain.Availability != wantedAvailability {
		t.Errorf("CheckAvailability(\"dnsbelgium.be\"), got %s, want %s", domain, wantedAvailability)
	}
}
