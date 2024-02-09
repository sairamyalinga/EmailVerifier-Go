package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	fmt.Print("Enter domain Name:")
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		verifyDomain(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal("Error:Failed to read from input: %v\n", err)
	}

}

func verifyDomain(domain string) {
	var hasMX, hasSPF, hasDMARC bool
	var SPFrecord, DMARCrecord string
	mxRecords, err := net.LookupMX(domain)

	if err != nil {
		log.Printf("Error: %v\n", err)
	}
	if len(mxRecords) > 0 {
		hasMX = true
	}
	txtRecords, err := net.LookupTXT(domain)
	if err != nil {
		log.Printf("Error: %v\n", err)
	}
	for _, record := range txtRecords {
		if strings.HasPrefix(record, "v=spf1") {
			hasSPF = true
			SPFrecord = record
			break
		}
	}
	dmarcRecords, err := net.LookupTXT("_dmarc." + domain)
	if err != nil {
		log.Printf("Error: %v\n", err)
	}
	for _, record := range dmarcRecords {
		if strings.HasPrefix(record, "v=DMARC1") {
			hasDMARC = true
			DMARCrecord = record
			break
		}
	}

	fmt.Printf("Domain: %v, Have MX: %v, Have SPF: %v, SPF record: %v, Have DMARC: %v, DMARC record: %v",domain,hasMX,hasSPF,SPFrecord,hasDMARC,DMARCrecord)

}
