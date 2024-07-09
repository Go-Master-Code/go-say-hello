package mylogger2

import (
	"fmt"
	"log"
)

func LgInfo(message string) {
	log.Printf("INFO - %v", message)
}

func Cetak() string {
	return "Hello"
}

func Perubahan() {
	fmt.Println("Function perubahan")
}
