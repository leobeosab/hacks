package commands

import (
	"bufio"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
	"strings"
)

func RunCommand(c string) (string, error) {
	output := ""

	log.Printf("Executing %s", c)

	ca := FormatCommandString(c)
	cmd := exec.Command(ca[0], ca[1:]...)
	stderr, err := cmd.StderrPipe()
	if err != nil {
		return output, errors.New("Error creating Std error pipe")
	}
	stdout, err := cmd.StdoutPipe()
	stdoutS := bufio.NewScanner(stdout)
	if err != nil {
		return output, errors.New("Error creating Std out pipe")
	}
	err = cmd.Start()
	if err != nil {
		errData, _ := ioutil.ReadAll(stderr)
		errString := fmt.Sprintf("Error starting %s\nError: %s\nData:\n %s", c, err.Error(), errData)
		return output, errors.New(errString)
	}

	for stdoutS.Scan() {
		data := stdoutS.Text() + "\n"
		output = output + data
		log.Printf("Command: [%s] \noutput: %s\n", c, data)
	}

	err = cmd.Wait()
	if err != nil {
		errData, _ := ioutil.ReadAll(stderr)
		errString := fmt.Sprintf("Error running %s\nError: %s\nData:\n %s", c, err.Error(), errData)
		return output, errors.New(errString)
	}

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
