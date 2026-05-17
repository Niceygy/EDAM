package errors

import "log"

func PanicIfErr(e error) {
	if e != nil {
		log.Panic(e.Error())
	}
}
