package main

import (
	"flag"
	"log"
)

func main() {
	var name string
	flag.StringVar(&name, "name", "(沒給名字)", "說明")
	flag.StringVar(&name, "n", "(沒給名字)", "說明")
	flag.Parse()

	log.Printf("name is %s", name)
}
