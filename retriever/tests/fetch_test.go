package tests

import (
	"github.com/WTRVGL/dns-be/retriever"
	"testing"
)

func TestAvailableDomain(t *testing.T) {
	domainName := "ifthisgetsregisteredthenripunittest.be"
	availability, _ := retriever.CheckDomainAvailability(domainName)
	want := retriever.DomainAvailability{
		Status:        "available",
		DateAvailable: "",
	}
	if availability != want {
		t.Errorf("vCheckDomainAvailability(\"ifthisgetsregisteredthenripunittest.be\"), got %s, want %s", availability, want)
	}
}

func TestUnavailableDomain(t *testing.T) {
	domainName := "dnsbelgium.be"
	availability, _ := retriever.CheckDomainAvailability(domainName)
	want := retriever.DomainAvailability{
		Status:        "inuse",
		DateAvailable: "",
	}
	if availability != want {
		t.Errorf("vCheckDomainAvailability(\"dnsbelgium.be\"), got %s, want %s", availability, want)
	}
}
