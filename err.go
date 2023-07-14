package main

import "log"

// if I want also to print messages from this func I need to add msg as parameter here
func CheckError(err error) {
	if err != nil {
		log.Fatalf("%v", err)
	}
}
