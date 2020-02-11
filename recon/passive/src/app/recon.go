package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"passiverecon/models"
	"passiverecon/scantools"
)

/* NOTICE:
 * For the love of all that is good
 * Do not use this as an example for how things should be done
 * -----------------------------------------
 * This is a hack to serve as a replacement for github.com/leobeosab/sharingan until
 * I have all the features I want for passive scanning done and stable
 */

func main() {
	input := flag.String("scanjson", "", "JSON file input")
	flag.Parse()

	s := ReadScanFile(*input)

	DNSScanning(&s)

	fmt.Printf("%v\n", s)
}

func DNSScanning(s *models.Scan) map[string][]models.Domain {
	unique := make(map[string][]models.Domain)

	for i, t := range s.Subdomains {
		if !t.Wildcard {
			continue
		}
		if t.Domains == nil {
			t.Domains = make(map[string]models.Domain)
		}

		domains := make([]models.Domain, 0)
		//domains = append(domains, scantools.AmassDNSEnumeration(t.Root)...)
		domains = append(domains, scantools.GOBustDNSBusting(t.Root, s.DNSWordlistPath)...)

		fmt.Printf("%v\n", t)
		for _, d := range domains {
			if _, ok := t.Domains[d.Name]; ok {
				continue
			}

			t.Domains[d.Name] = d
			unique[t.Root] = append(unique[t.Root], d)
		}

		s.Subdomains[i] = t
	}

	return unique
}

func ReadScanFile(path string) models.Scan {
	jf, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}
	defer jf.Close()

	bv, err := ioutil.ReadAll(jf)
	if err != nil {
		fmt.Println(err)
	}

	var s models.Scan
	err = json.Unmarshal([]byte(bv), &s)
	if err != nil {
		fmt.Println(err)
	}

	return s
}
