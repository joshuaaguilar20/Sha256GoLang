
// Write a program that is given a list of file names as arguments then prints
// the sha256 sum for the contents of each file. Print the hashes as a hex string.
package main

import (
	//"strings"
	"crypto/sha256"
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	"os"
)

func main() {
	argsWithoutProg := os.Args[1:]
	var failures int

	if failures > 0 {
		os.Exit(1)
	}
	for i := 0; i < len(argsWithoutProg); i++ {
		err := processData(argsWithoutProg[i])
		if err != nil {
			failures++

		}
	}

}
func processData(fileName string) error {
	info, errDir := os.Stat(fileName)
	if errDir != nil {
		return errors.Wrap(errDir, "stat file")
	}
	if info.IsDir() {
		return nil
	}
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return errors.Wrap(err, "Read file ioutil readfile")
	}
	h := sha256.New()
	h.Write([]byte(data))
	fmt.Printf("%x\t%s\n", h.Sum(nil), fileName)
	return nil
}
