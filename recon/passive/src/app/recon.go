package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"passiverecon/helpers"
	"passiverecon/models"
	"passiverecon/notify"
	"passiverecon/scantools"
)

/* NOTICE:
 * For the love of all that is good
 * Do not use this as an example for how things should be done
 * -----------------------------------------
 * This is a hack to serve as a replacement for github.com/leobeosab/sharingan until
 * I have all the features I want for passive scanning done and stable
 * TLDR; This code fucking sucks
 */

func main() {
	scanfile := flag.String("scanjson", "", "JSON file input")
	flag.Parse()

	s := ReadScanFile(*scanfile)

	dnsRes := DNSScanning(&s)
	for t, a := range dnsRes {
		notify.NotifyDomains(t, "Unique Domains Found: ", &a, notify.Settings().ScanWHName)
	}
	dirb := DirBusting(&s)
	for t, a := range dirb {
		notify.NotifyDirBustResults(t, &a)
	}

	fmt.Printf("%v\n", s)
	if !SaveScanFile(s, *scanfile) {
		notify.SendError("Generic", "Could not save JSON", nil)
	}
}

func DirBusting(s *models.Scan) map[string][]models.DirBustResult {
	results := make(map[string][]models.DirBustResult, 0)
	for tk, tv := range s.Subdomains {
		for dk, dv := range tv.Domains {
			// I HATE THIS
			// YET HERE I AM WRITING IT
			// What am I doing with my life
			dirb := make([]models.DirBustResult, 0)
			for _, url := range helpers.ReturnActiveWebPortURLS(dv.Name) {
				// Get dirbed urls, add to results for ez pz discord notification?
				urls, err := scantools.DirBust(url, s.DirbustWordlistPath)
				if err != nil {
					notify.SendError(dv.Name, "Dirbust", err)
				}
				dirb = append(dirb, urls...)
			}

			dv.DirbResults = dirb
			s.Subdomains[tk].Domains[dk] = dv
			results[dv.Name] = dirb
		}
	}

	return results
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
		amassdomains, err := scantools.AmassDNSEnumeration(t.Root)
		if err != nil {
			notify.SendError(t.Root, "Amass DNS Enumeration, Found: "+string(len(amassdomains)), err)
		}
		notify.NotifyDomains(t.Root, "Amass Results: ", &amassdomains, notify.Settings().LoggingWHName)

		domains = append(domains, amassdomains...)
		gobustdomains, err := scantools.GOBustDNSBusting(t.Root, s.DNSWordlistPath)
		if err != nil {
			notify.SendError(t.Root, "Error doing DNS Busting", err)
		}
		notify.NotifyDomains(t.Root, "Gobuster DNS Results: ", &gobustdomains, notify.Settings().LoggingWHName)

		domains = append(domains, gobustdomains...)

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

func SaveScanFile(scan models.Scan, scanfile string) bool {
	data, err := json.MarshalIndent(scan, "", "\t")
	if err != nil {
		return false
	}
	err = ioutil.WriteFile(scanfile, data, 0644)
	if err != nil {
		return false
	}
	return true
}
