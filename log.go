package core

import "log"

func init() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	log.SetPrefix("[WSC]")
}
