package scantools

import (
	"fmt"
	"log"
	"passiverecon/commands"
	"passiverecon/models"
	"regexp"
	"strings"
)

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
