package test

import (
	"github.com/WTRVGL/dns-be"
	"testing"
)

func TestValidDomain(t *testing.T) {
	domainName := "dnsbelgium.be"
	_, err := dns_be.NewDomain(domainName)
	if err != nil {
		t.Errorf("validateDomainName(\"dnsbelgium.be\", %d)", err)
	}
}

func TestIncompatibleTLD(t *testing.T) {
	domainName := "wouter.land"
	_, err := dns_be.NewDomain(domainName)
	want := dns_be.BadTLD
	if err != want {
		t.Errorf("NewDomain(\"wouter.land\", %d)", err)
	}
}
func TestNonAcceptedSymbol(t *testing.T) {
	domainName := "wouter_.be"
	_, err := dns_be.NewDomain(domainName)
	want := dns_be.InvalidDomain
	if err != want {
		t.Errorf("NewDomain(\"wouter.land\", %d)", err)
	}
}

func TestInvalidFormat(t *testing.T) {
	domainName := "api.wouter.be"
	_, err := dns_be.NewDomain(domainName)
	want := dns_be.InvalidDomain
	if err != want {
		t.Errorf("NewDomain(\"wouter.land\", %d)", err)
	}
}
