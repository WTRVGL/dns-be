package domain

import (
	"github.com/WTRVGL/dns-be/internal/models"
	checker "github.com/WTRVGL/dns-be/pkg"
	"testing"
)

func TestAvailableDomain(t *testing.T) {
	domainName := "ifthisgetsregisteredthenripunittest.be"
	domain, _ := checker.NewDomain(domainName)
	*domain, _ = domain.CheckAvailability()
	wantedAvailability := models.Availability{Status: "available", DateAvailable: ""}
	if domain.Availability != wantedAvailability {
		t.Errorf("CheckAvailability(\"ifthisgetsregisteredthenripunittest.be\"), got %s, want %s", domain.Availability, wantedAvailability)
	}
}

func TestUnavailableDomain(t *testing.T) {
	domainName := "dnsbelgium.be"
	domain, _ := checker.NewDomain(domainName)
	*domain, _ = domain.CheckAvailability()
	wantedAvailability := models.Availability{Status: "inuse", DateAvailable: ""}
	if domain.Availability != wantedAvailability {
		t.Errorf("CheckAvailability(\"dnsbelgium.be\"), got %s, want %s", domain, wantedAvailability)
	}
}
