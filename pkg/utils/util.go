package utils

import "log"

func ErrorCheck(description string, e error) {
	if e != nil {
		log.Fatalf(description, e)
	}
}
