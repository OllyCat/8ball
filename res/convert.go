package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func main() {
	fl, err := filepath.Glob("*.png")
	if err != nil {
		panic(err)
	}

	answers := make([]string, 0)

	for _, fn := range fl {
		f, err := os.Open(fn)
		if err != nil {
			log.Printf("File \"%s\" open error:	%v\n", fn, err)
		}
		data, err := ioutil.ReadAll(f)
		if err != nil {
			log.Printf("File \"%s\" read error:	%v\n", fn, err)
		}
		f.Close()

		ds := strconv.Quote(string(data))
		name := strings.TrimSuffix(fn, ".png") + "_png"
		if strings.HasPrefix(fn, "ans") {
			answers = append(answers, ds)
		} else {
			fmt.Printf("%s=%s\n\n", name, ds)
		}
	}

	fmt.Printf("answers := []string{\n")
	for _, s := range answers {
		fmt.Printf("\t%s,\n", s)
	}
	fmt.Printf("}\n")
}
