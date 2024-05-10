package main

import (
    "net/http"
    "os/exec"
    "time"
	"log"
)

var warningAlreadySent = false

func checkInternet() {
	client := &http.Client{
        Timeout: 2 * time.Second,
    }

     _, err := client.Get("http://google.com")
	log.Println(err)
    if err == nil {
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