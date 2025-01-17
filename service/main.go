package main

import (
	"net/http"
	"os"

	"github.com/ikatheria/Jarvis-personal-mine/service/controllers"
	"github.com/ikatheria/Jarvis-personal-mine/service/logger"
)

var (
	port = os.Getenv("PORT")
)

func main() {
	if len(port) == 0 {
		port = "3000"
	}

	{
		http.HandleFunc("/", controllers.HomeController)
		http.HandleFunc("/message", controllers.MessagesController)
		http.HandleFunc("/reminder", controllers.ReminderController)
		http.HandleFunc("/email", controllers.EmailController)
	}

	logger.Critic(http.ListenAndServe(":"+port, nil))
}
