package controllers

import (
	"fmt"
	"time"

	"github.com/ikatheria/Jarvis-personal-mine/service/logger"
	"github.com/ikatheria/htgo-tts"
)

// TextToSpeech converts text to speech on a string, return status as boolean
func TextToSpeech(message string, waitInSeconds int32) bool {

	time.Sleep(time.Duration(waitInSeconds))
	fmt.Println("Speaking -> ", message)
	speech := htgotts.Speech{Folder: "audio", Language: "en"}
	err := speech.Speak(message)
	if err != nil {
		logger.Error(err)
	}
	return true

}
