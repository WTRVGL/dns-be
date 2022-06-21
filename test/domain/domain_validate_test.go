package domain

import (
	"testing"

	"github.com/WTRVGL/dns-be/internal/models/errors"
	"github.com/WTRVGL/dns-be/pkg"
)

func TestValidDomain(t *testing.T) {
	domainName := "dnsbelgium.be"
	_, err := checker.NewDomain(domainName)
	if err != nil {
		t.Errorf("validateDomainName(\"dnsbelgium.be\", %d)", err)
	}
}

func TestIncompatibleTLD(t *testing.T) {
	domainName := "wouter.land"
	_, err := checker.NewDomain(domainName)
	want := errors.BadTLD
	if err != want {
		t.Errorf("NewDomain(\"wouter.land\", %d)", err)
	}
}
func TestNonAcceptedSymbol(t *testing.T) {
	domainName := "wouter_.be"
	_, err := checker.NewDomain(domainName)
	want := errors.InvalidDomain
	if err != want {
		t.Errorf("NewDomain(\"wouter.land\", %d)", err)
	}
}

func TestInvalidFormat(t *testing.T) {
	domainName := "api.wouter.be"
	_, err := checker.NewDomain(domainName)
	want := errors.InvalidDomain
	if err != want {
		t.Errorf("NewDomain(\"wouter.land\", %d)", err)
	}
}
