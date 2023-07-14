package main

import (
	"database/sql"
	"fmt"
)

// It will hold row data returned from query

var db *sql.DB

func main() {

	connectDb()

	albums, err := albumsByArtist("John Coltrane")

	CheckError(err)
	fmt.Printf("Albums found: %v\n", albums)

	// HARD-code ID 2 here to test the query.
	alb, err := albumByID(2)
	CheckError(err)
	fmt.Printf("Album found: %v\n", alb)

	albID, err := addAlbum(Album{
		Title:  "Blue Train",
		Artist: "John Coltrane",
		Price:  56.99,
	})

	CheckError(err)
	fmt.Printf("ID of added album: %v\n", albID)
}

//Look at the code by ignoring err parts. It will be more clear
//but be carefull which errors are you skipping. Some of them
// include parts which you need to read
