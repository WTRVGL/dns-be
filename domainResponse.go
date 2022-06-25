package dns_be

type DomainResponse struct {
	DomainInfo struct {
		Alabel             string `json:"alabel"`
		Ulabel             string `json:"ulabel"`
		Status             string `json:"status"`
		Created            string `json:"created"`
		Updated            string `json:"updated"`
		DateAvailable      string `json:"dateAvailable"`
		VerificationStatus string `json:"verificationStatus"`
	}
	NameserverInfo struct {
		nameservers []struct {
			Name      string `json:"name"`
			NameASCII string `json:"nameASCII"`
			Ip        string `json:"ip"`
		}
		Overridden          bool   `json:"Overridden"`
		RegistryNsgroupName string `json:"registryNsgroupName"`
	}
	Registrant            interface{}   `json:"registrant"`
	OnsiteCompanyContacts []interface{} `json:"onsiteCompanyContacts"`
	TechContacts          []interface{} `json:"techContacts"`
	DnsKeyInfo            struct {
		DnsKeys interface{} `json:"dnsKeys"`
	}
	TransferProhibitedStatus struct {
		TransferProhibited             bool `json:"transferProhibited"`
		ServerTransferProhibited       bool `json:"ServerTransferProhibited"`
		MaskedClientTransferProhibited bool `json:"MaskedClientTransferProhibited"`
	}
}
