package main

import (
	"flag"
	"fmt"
	checker "github.com/WTRVGL/dns-be/pkg"
	"os"
)

func main() {
	domainNameFlag := flag.String("n", "", ".be domain to be checked")
	flag.Parse()

	domain, err := checker.NewDomain(*domainNameFlag)
	if err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(1)
	}
	domain, err = domain.CheckAvailability()

	if err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(1)
	}

	switch domain.Availability.Status {
	case "inuse":
		fmt.Printf("domain %s\n is currently in use", *domainNameFlag)
		os.Exit(0)
		break
	case "quarantine":
		fmt.Printf("domain %s is quarantined and becomes available %s\n", *domainNameFlag, domain.Availability.DateAvailable)
		os.Exit(0)
	case "available":
		fmt.Printf("domain %s is available\n", *domainNameFlag)
		os.Exit(0)
		break
	}

}
