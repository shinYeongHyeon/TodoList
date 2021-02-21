package api

import "github.com/labstack/gommon/log"

func must(err error) {
	if err != nil {
		log.Error("Internal error : ", err)
		panic(internalError)
	}
}