package main

import (
	"log"
	"os"
	"runtime"

	"github.com/chaintex/currencies/http"
)

func enableLogToFile() (*os.File, error) {
	const logFileName = "log/error.log"
	f, err := os.OpenFile(logFileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return nil, err
	}

	//clear error log file
	if err = f.Truncate(0); err != nil {
		log.Fatal(err)
	}

	log.SetOutput(f)
	return f, nil
}

func main() {
	numCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(numCPU)
	//set log for server
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	chainTexENV := os.Getenv("CHAINTEX_ENV")

	if os.Getenv("LOG_TO_STDOUT") != "true" {
		f, err := enableLogToFile()
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
	}

	//run server
	server := http.NewHTTPServer(":3002")
	server.Run(chainTexENV)
}
