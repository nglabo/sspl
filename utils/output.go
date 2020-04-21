package utils

import "log"

func NewError(msg string, err error) {
	log.Fatalf("\033[1;31m%v\033[0m %v: %v\n", "[error]", msg, err)
}
