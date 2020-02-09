package commands

import (
	"log"
	"os/exec"
	"strings"
)

func RunCommand(c string) (string, error) {
	ca := FormatCommandString(c)
	out, err := exec.Command(ca[0], ca[1:]...).Output()
	if err != nil {
		return "", err
	}

	log.Printf("Executing %s", c)
	output := string(out[:])
	return output, nil
}

func FormatCommandString(c string) []string {
	r := make([]string, 0)
	tmp := ""

	inQ := false

	// dirty but works
	spl := strings.Split(c, " ")
	for _, s := range spl {
		containsQ := strings.Contains(s, "\"")
		if !inQ && containsQ { // If not in quote block but string contains quote set flag
			inQ = true
			tmp += s
		} else if inQ && containsQ { // If in quotes and contains a quote end quote block and append content to array
			inQ = false
			tmp += s
			r = append(r, tmp)
			tmp = ""
		} else if inQ && !containsQ { // If in quotes and no quote in string append to tmp
			tmp += " " + s
		} else { // Add to result if no quotes conditions are met
			r = append(r, s)
		}
	}

	return r
}
