package mylogger2

import "log"

func LgInfo(message string) {
	log.Printf("INFO - %v", message)
}
