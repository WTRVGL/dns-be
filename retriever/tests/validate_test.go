package tests

import (
	"github.com/WTRVGL/dns-be/retriever"
	"testing"
)

func TestValidDomain(t *testing.T) {
	domain := "wouter.be"
	isValidDomain, err := retriever.ValidateDomainName(domain)
	want := true
	if isValidDomain != want {
		t.Errorf("validateDomainName(\"dnsbelgium.be\", %d), got %t, want %t", err, isValidDomain, false)
	}
}

func TestIncompatibleTLD(t *testing.T) {
	domain := "wouter.land"
	isValidDomain, err := retriever.ValidateDomainName(domain)
	want := "not a .be domain"
	if err.Error() != want {
		t.Errorf("validateDomainName(\"wouter.land\", %d), got %t, want %t", err, isValidDomain, false)
	}
}

func TestNonAlphaNumeric(t *testing.T) {
	domain := "w0ut3r.be"
	isValidDomain, err := retriever.ValidateDomainName(domain)
	want := "invalid domain"
	if err.Error() != want {
		t.Errorf("validateDomainName(\"w0ut3r.be\", %d), got %t, want %t", err, isValidDomain, false)
	}
}

func TestInvalidFormat(t *testing.T) {
	domain := "api.wouter.be"
	isValidDomain, err := retriever.ValidateDomainName(domain)
	want := "invalid domain"
	if err.Error() != want {
		t.Errorf("validateDomainName(\"api.wouter.be\", %d), got %t, want %t", err, isValidDomain, false)
	}
}
