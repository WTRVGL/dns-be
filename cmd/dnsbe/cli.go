package main

import (
	"errors"
	"flag"
	"fmt"
	checker "github.com/WTRVGL/dns-be"
	"os"
)

var (
	NoArguments = errors.New("no arguments used, please execute with a -n flag\n see -h for more information")
	DomainInUse = errors.New("domain is currently in use")
)

func main() {
	err := Run()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	os.Exit(0)
}

func Run() error {
	domainNameFlag := flag.String("n", "", ".be domain to be checked")
	flag.Parse()

	if *domainNameFlag == "" {
		return NoArguments
	}

	domain, err := checker.NewDomain(*domainNameFlag)
	if err != nil {
		return err
	}

	domain, err = domain.CheckAvailability()

	if err != nil {
		return err
	}

	switch domain.Availability.Status {
	case "inuse":
		return DomainInUse
	case "quarantine":
		fmt.Printf("domain %s is quarantined and becomes available %s\n", *domainNameFlag, domain.Availability.DateAvailable)
		break
	case "available":
		fmt.Printf("domain %s is available\n", *domainNameFlag)
		break
	}

	return nil
}
