package internal

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

// App application context
type App struct {
	hostName  string
	reportURL string
}

func getEnvOrDie(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("\"%v\" is required", key)
		os.Exit(1)
	}
	return value
}

// Launch launches the daemon
// recursive
func (app *App) collectInfo(channel chan *ServerInformation) {
	stat := GetStat()
	if stat != nil {
		channel <- stat
		app.collectInfo(channel)
	} else {
		close(channel)
	}
}

// ServerPayload payload
type ServerPayload struct {
	Host string             `json:"host"`
	Stat *ServerInformation `json:"stat"`
}

// sending information over http
func (app *App) sendInformation(info *ServerInformation) {

	payload := ServerPayload{
		Host: app.hostName,
		Stat: info,
	}

	js, err := json.Marshal(payload)
	if err != nil {
		return
	}

	response, err := http.Post(app.reportURL, "application/json", bytes.NewBuffer(js))
	if err != nil {
		fmt.Printf("FAILED: to post : %v => %v\n", app.reportURL, string(js))
		fmt.Printf("%v\n", err)
	} else {
		fmt.Printf("SUCCESS: payload to : %v => %v\n", app.reportURL, string(js))
		fmt.Printf("%v\n", response)
	}

	defer func() {
		if r := recover(); r != nil {
			fmt.Print(err)
		}
	}()

}

// Launch launches the daemon
func (app *App) Launch() {
	// configure http to not complain about https
	cfg := &tls.Config{
		InsecureSkipVerify: true,
	}
	http.DefaultClient.Transport = &http.Transport{
		TLSClientConfig: cfg,
	}

	app.parseEnv()

	channel := make(chan *ServerInformation)
	go app.collectInfo(channel)
	for {
		info, more := <-channel
		if more {
			app.sendInformation(info)
		} else {
			break
		}
	}
}

// Init the application
func (app *App) parseEnv() {
	hostName := getEnvOrDie("HOST_NAME")
	reportURL := getEnvOrDie("REPORT_URL")
	app.hostName = hostName
	app.reportURL = reportURL
}

// CreateApp creates application
func CreateApp() *App {

	app := &App{}

	return app
}
