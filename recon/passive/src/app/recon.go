package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	models "passiverecon/types"
)

/* NOTICE:
 * For the love of all that is good
 * Do not use this as an example for how things should be done
 * -----------------------------------------
 * This is a hack to serve as a replacement for github.com/leobeosab/sharingan until
 * I have all the features I want for passive scanning done and stable
 */

func main() {
	r := ReadScanFile()
	fmt.Printf("%v\n", r)
}

func ReadScanFile() []string {
	tmpLocation := "ryanwise.me.scan.json"

	jf, err := os.Open(tmpLocation)
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

	fmt.Printf("%v", s)

	return []string{}
}
