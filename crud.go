package main

import (
	"database/sql"
	"fmt"
)

func albumsByArtist(name string) ([]Album, error) {
	//album slice to hold data from returned rows.
	var albums []Album

	rows, err := db.Query("SELECT * FROM album WHERE artist = ?", name)
	if err != nil {
		return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
	}

	// we are closing rows becaause if ant resources it holds will be released
	//when the func exits.
	defer rows.Close()

	for rows.Next() {
		var alb Album

		//can takes a list of pointers to Go values, where the column values
		//will be written. Here, you pass pointers to fields in the alb variable,
		//created using the & operator. Scan writes through the pointers
		// to update the struct fields.

		if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
			return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
		}

		albums = append(albums, alb)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
	}
	return albums, nil
}

//this func queries for the album with the specified ID.

func albumByID(id int64) (Album, error) {
	var alb Album

	row := db.QueryRow("SELECT * FROM album WHERE id = ?", id)
	if err := row.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
		if err == sql.ErrNoRows {
			return alb, fmt.Errorf("albumsById %d: %v", id, err)
		}
		return alb, fmt.Errorf("albumsById %d: %v", id, err)
	}
	return alb, nil
}

func addAlbum(alb Album) (int64, error) {
	result, err := db.Exec("INSERT INTO album (title, artist, price) VALUES (?, ?, ?)", alb.Title, alb.Artist, alb.Price)

	// check for an error from the attempt to INSERT
	if err != nil {
		return 0, fmt.Errorf("addAlbum: %v", err)
	}

	id, err := result.LastInsertId()

	// check for an error from the attempt to retrieve the ID
	if err != nil {
		return 0, fmt.Errorf("addAlbum: %v", err)
	}

	return id, nil
}

func deleteAlbum(id int64) (int64, error) {
	_, err := db.Exec("DELETE FROM album WHERE id = ? ", id)

	if err != nil {
		return 0, fmt.Errorf("deleteAlbum: %v", err)
	} else {
		fmt.Printf("Album with id %v is deleted", id)
	}
	return 1, nil
}
