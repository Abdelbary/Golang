package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	path := flag.String("path", "logs.log", "the log file")
	level := flag.String("level", "ERROR", "the level of parsing  Options [EROOR WARRNING DEBUG INFO]")

	flag.Parse()

	f, err := os.Open(*path)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	r := bufio.NewReader(f)

	for {
		s, err := r.ReadString('\n')

		if err != nil {
			break
		}
		if strings.Contains(s, *level) {
			fmt.Println(s)
		}
	}

}
