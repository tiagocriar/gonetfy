package main

import (
	"log"
	"net/http"
	"os/exec"
	"time"
)

var warningAlreadySent = false

func checkInternet() {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	resp, err := client.Get("http://clients3.google.com/generate_204")
	if err != nil {
		log.Println(err)
		if !warningAlreadySent {
			sendWarning()
			warningAlreadySent = true
		}
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNoContent {
		if warningAlreadySent {
			sendInfo()
			warningAlreadySent = false
		}
	} else {
		if !warningAlreadySent {
			sendWarning()
			warningAlreadySent = true
		}
	}
}

func sendWarning() {
	exec.Command("notify-send", "Internet caiu, man! Relaxa e espera um pouco.").Run()
}

func sendInfo() {
	warningAlreadySent = false
	exec.Command("notify-send", "Internet voltou, hora de trabalhar!").Run()
}

func main() {
	for {
		checkInternet()
		time.Sleep(3 * time.Second)
	}
}
