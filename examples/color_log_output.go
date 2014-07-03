package main

import (
	"github.com/appio/logutils"
	"log"
)

func main() {
	filter := logutils.NewFilter(nil, true)

	logger := log.New(filter, "", 0)

	logger.Printf("[DEBUG] This is a debug message")
	logger.Printf("[INFO] This is an info message")
	logger.Printf("[WARN] This is a warning")
	logger.Printf("[ERROR] This is an error message")
	logger.Printf("[CRIT] Sh!t just got real")
}
