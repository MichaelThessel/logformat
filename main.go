package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	reader := bufio.NewScanner(os.Stdin)

	entry := ""
	for reader.Scan() {
		text := reader.Text()
		if text != "" {
			entry = entry + reader.Text() + "\n"
		}
	}

	if strings.HasSuffix(entry, "\n") {
		entry = entry[0 : len(entry)-1]
	}

	if entry == "" {
		return
	}
	t := time.Now()
	logTime := t.Format("Mon Jan 2 15:04:05 -0700 2006")

	tag := flag.String("t", "default", "Tag to apply")
	flag.Parse()

	fmt.Printf("[%s] [%s] %s\n", logTime, *tag, entry)
}
