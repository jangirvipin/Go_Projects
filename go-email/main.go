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
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("domain , hasMX, hasSPF, spfrecord, hasDMARC,dmarcRecord")
	for scanner.Scan() {
		checkDomain(scanner.Text())
	}
}

func checkDomain(domain string) {
	var hasMX, hasSPF, hasDMARC bool
	var spfRecord, dmarcRecord string

	//mxRecords
	mxRecord, err := net.LookupMX(domain)
	if err != nil {
		log.Println("MX lookup error:", err.Error())
	}

	if len(mxRecord) > 0 {
		hasMX = true
	}

	//spfRecord
	txtRecords, err := net.LookupTXT(domain)
	if err != nil {
		log.Println("TXT lookup error:", err)
	}

	for _, record := range txtRecords {
		if strings.HasPrefix(record, "v=spf1") {
			hasSPF = true
			spfRecord = record
			break
		}
	}

	//hasDMARC
	dmarcRecords, err := net.LookupTXT("_dmarc." + domain)
	if err != nil {
		log.Println("TXT lookup error:", err)
	}

	for _, record := range dmarcRecords {
		if strings.HasPrefix(record, "v=DMARC1") {
			hasDMARC = true
			dmarcRecord = record
			break
		}
	}

	fmt.Printf("%v, %v, %v, %v, %v, %v \n", domain, hasMX, hasSPF, spfRecord, hasDMARC, dmarcRecord)
}
