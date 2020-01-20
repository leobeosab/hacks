package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

/*
 * Read list of subdomains on stdin
 * output list of words for s3bucket smashing
 */

func main() {

	info, err := os.Stdin.Stat()
	if err != nil {
		panic(err)
	}

	if info.Mode()&os.ModeNamedPipe == 0 {
		log.Println("Pipe in your subdomains")
		log.Println("Usage: cat subs | substowords rootdomain otherrootdomains")
		return
	}

	reader := bufio.NewScanner(os.Stdin)
	subdomains := make([]string, 0)

	for reader.Scan() {
		subdomains = append(subdomains, reader.Text())
	}

	for _, s := range subdomains {
		spl := strings.SplitN(s, ".", -1)
		spl = spl[:len(spl)-2]
		fmt.Println(strings.Join(spl, "-"))
		fmt.Println(strings.Join(spl, ""))
	}

}
