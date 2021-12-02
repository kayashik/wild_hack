package models

import (
	"database/sql"
	"fmt"
)

// Create an exported global variable to hold the database connection pool.
var DB *sql.DB

type Comment struct {
	Name string
}

func AllComments() ([]Comment, error) {
	// Note that we are calling Query() on the global variable.
	rows, err := DB.Query("SELECT * FROM comments")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var comments []Comment

	for rows.Next() {
		var comment Comment

		err := rows.Scan(&comment.Name)
		if err != nil {
			return nil, err
		}

		comments = append(comments, comment)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return comments, nil
}

func AddComment(name string) error {
	rows, err := DB.Query(fmt.Sprintf("INSERT INTO comments VALUES '%s'", name))
	if err != nil {
		return err
	}
	defer rows.Close()
	return nil
}
