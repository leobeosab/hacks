package scantools

import (
	"fmt"
	"log"
	"passiverecon/commands"
	"passiverecon/models"
	"regexp"
	"strconv"
	"strings"
)

func DirBust(URL string, wordlistpath string) []models.DirBustResult {
	results := make([]models.DirBustResult, 0)

	// e: full url, q: quiet, z: no progress, t: threads
	rc := fmt.Sprintf("gobuster dir -e -z -t 25 -q --url %s --wordlist %s", URL, wordlistpath)
	data, err := commands.RunCommand(rc)
	if err != nil {
		log.Println("Error encountered during dirb.... skipping")
		log.Println(err)
		return results // return empty array
	}

	reg := regexp.MustCompile(`(\d)*`) // match only digits for status
	for _, u := range strings.Split(data, "\n") {
		content := strings.Split(u, " ")
		if len(content) != 3 || content[0] == "" || content[2] == "" {
			continue // Make sure we don't process anything not formatted correctly
		}

		path := content[0]
		status, err := strconv.Atoi(reg.FindString(content[2]))
		if err != nil {
			log.Printf("Error parsing status for %s\n", u)
		}

		r := models.DirBustResult{
			Path:   path,
			Status: status,
		}

		log.Printf("Found! Path: %s Status: %d\n", path, status)

		results = append(results, r)
	}

	return results
}

func AmassDNSEnumeration(domain string) []models.Domain {

	rc := fmt.Sprintf("amass enum -d %s", domain)
	data, err := commands.RunCommand(rc)
	if err != nil {
		log.Println("Error encountered during AmassEnumeration.... skipping")
		log.Println(err)
		return []models.Domain{}
	}

	subs := strings.Split(data, "\n")
	domains := make([]models.Domain, 0)

	for _, s := range subs {
		if len(s) == 0 {
			continue
		}

		d := models.Domain{
			Name: s,
		}

		domains = append(domains, d)
	}

	return domains
}

func GOBustDNSBusting(domain string, wordlistpath string) []models.Domain {
	rc := fmt.Sprintf("gobuster -z -t 25 -q dns --domain %s --wordlist %s", domain, wordlistpath)
	data, err := commands.RunCommand(rc)
	if err != nil {
		log.Println("Error encountered during Gobuster DNS Enumeration.... skipping")
		log.Println(err)
		return []models.Domain{}

	}
	reg := regexp.MustCompile(`[\S]*`)
	subs := strings.Split(data, "Found: ")
	domains := make([]models.Domain, 0)

	for _, s := range subs {
		s = reg.FindStringSubmatch(s)[0]
		if len(s) == 0 {
			continue
		}

		d := models.Domain{
			Name: s,
		}

		domains = append(domains, d)
	}

	return domains
}
