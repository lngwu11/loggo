package loggo_test

import (
	"log"
	"net/http"

	"github.com/lngwu11/loggo"
)

func ExampleLogger_Writer_httpServer() {
	logger := loggo.New()
	w := logger.Writer()
	defer w.Close()

	srv := http.Server{
		// create a stdlib log.Logger that writes to
		// loggo.Logger.
		ErrorLog: log.New(w, "", 0),
	}

	if err := srv.ListenAndServe(); err != nil {
		logger.Fatal(err)
	}
}

func ExampleLogger_Writer_stdlib() {
	logger := loggo.New()
	logger.Formatter = &loggo.JSONFormatter{}

	// Use loggo for standard log output
	// Note that `log` here references stdlib's log
	// Not loggo imported under the name `log`.
	log.SetOutput(logger.Writer())
}
