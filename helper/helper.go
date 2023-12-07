package helper

import "log"

func CheckErrorNil(err error) {
	if err != nil {
		log.Panic(err)
	}
}
