package main

import (
	"log"
	"log/slog"
	"os"
)

func main() {

	log.Println("hey!")

	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("hey again!")

	log.SetPrefix("PREFIX: ")
	log.Println("hey again!")

	log.SetFlags(log.LstdFlags | log.Lshortfile | log.Lmsgprefix)
	log.Println("hey again!")

	jsonHandler := slog.NewJSONHandler(os.Stderr, nil)
	myLog := slog.New(jsonHandler)
	myLog.Info("hey again")
}
