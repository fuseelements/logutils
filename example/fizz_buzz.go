package main

import (
	"github.com/appio/logutils"
	stdlog "log"
)

func main() {
	filter := logutils.NewFilter(nil, true)

	log := stdlog.New(filter, "", 0)

	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			log.Printf("[CRIT] FIZZBUZZ")
		} else if i%5 == 0 {
			log.Printf("[ERROR] BUZZ")
		} else if i%3 == 0 {
			log.Printf("[WARN] FIZZ")
		} else {
			log.Printf("[DEBUG] %d", i)
		}
	}
}
