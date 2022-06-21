package test

import (
	"testing"

	"github.com/WTRVGL/dns-be/internal/entities/errors"
	"github.com/WTRVGL/dns-be/pkg"
)

func TestValidDomain(t *testing.T) {
	domainName := "dnsbelgium.be"
	_, err := pkg.NewDomain(domainName)
	if err != nil {
		t.Errorf("validateDomainName(\"dnsbelgium.be\", %d)", err)
	}
}

func TestIncompatibleTLD(t *testing.T) {
	domainName := "wouter.land"
	_, err := pkg.NewDomain(domainName)
	want := errors.BadTLD
	if err != want {
		t.Errorf("NewDomain(\"wouter.land\", %d)", err)
	}
}
func TestNonAcceptedSymbol(t *testing.T) {
	domainName := "wouter_.be"
	_, err := pkg.NewDomain(domainName)
	want := errors.InvalidDomain
	if err != want {
		t.Errorf("NewDomain(\"wouter.land\", %d)", err)
	}
}

func TestInvalidFormat(t *testing.T) {
	domainName := "api.wouter.be"
	_, err := pkg.NewDomain(domainName)
	want := errors.InvalidDomain
	if err != want {
		t.Errorf("NewDomain(\"wouter.land\", %d)", err)
	}
}
