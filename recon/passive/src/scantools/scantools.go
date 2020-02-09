package scantools

import (
	"fmt"
	"passiverecon/commands"
	"passiverecon/models"
	"strings"
)

func AmassDNSEnumeration(domain string) []models.Domain {

	rc := fmt.Sprintf("amass enum -d %s", domain)
	data, err := commands.RunCommand(rc)
	if err != nil {
		fmt.Println(err)
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
