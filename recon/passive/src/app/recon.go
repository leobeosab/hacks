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
	for i, d := range s.Subdomains {
		domains := make([]models.Domain, 0)
		domains = append(domains, scantools.AmassDNSEnumeration(d.Root)...)

		// TODO: move domains / subdomains into a map for checking if unique
		s.Subdomains[i].Domains = domains
	}

	fmt.Printf("%v\n", s)
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
