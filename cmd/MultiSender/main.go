package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/BlathanAevon/MultiSender/internal/sender"
	"github.com/BlathanAevon/MultiSender/tools"
)

func main() {

	config, err := tools.ParseFlags()

	if err != nil {
		flag.Usage()
		fmt.Println()
		log.Fatal(err)
	}

	if config == nil {
		flag.Usage()
		os.Exit(1)
	}

	if err := sender.Disperse(config); err != nil {
		log.Fatal(err)
	}

}
