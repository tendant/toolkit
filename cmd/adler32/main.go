package main

import (
	"fmt"
	"hash/adler32"
)

func main() {
	domainHash("example.com")
}

func domainHash(domain string) {
	checkSum := adler32.Checksum([]byte(domain))
	fmt.Printf("Checksum for domain(%s): %x\n", domain, checkSum)
	fmt.Printf("DomainLabelKey (%s)\n acme.cert-manager.io/http-domain: \"%d\"\n", domain, checkSum)
}
