package config

import "log"

var err error

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
