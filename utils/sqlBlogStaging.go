package utils

import (
	"database/sql"
	"encoding/json"

	_ "github.com/go-sql-driver/mysql" //SQL Driver for MySQL/MariaDB
	"github.com/pkg/errors"
)

//StagingPosts is a slice of 'Post'
type StagingPosts []StagingPost

//StagingPost represents a single blog post from our 'poststaging' table
type StagingPost struct {
	ID    int
	Tags  []string
	Image string
	Title string
	Body  string
}

//Create (StagingPost) adds a new blog post to our 'poststaging' table
func (sp *StagingPost) Create() error {
	b, err := json.Marshal(sp.Tags)
	if err != nil {
		return errors.WithStack(err)
	}

	_, err = connection.Exec(`INSERT INTO poststaging (date, tags, image, title, body) VALUES (?, ?, ?, ?, ?);`, b, sp.Image, sp.Title, sp.Body)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

//Read (StagingPosts) retrieves 'limit' most recent posts
//
//This is used to populate blog.ChristianHering.com's home page
func (sps *StagingPosts) Read(limit string) error {
	r, err := connection.Query(`SELECT * FROM poststaging ORDER BY id DESC LIMIT ?;`, limit)
	if err != nil {
		return errors.WithStack(err)
	}
	defer r.Close()

	return parseBlogStagingRows(sps, r)
}

//Update (StagingPost) modifies a blog post entry
func (sp *StagingPost) Update() error {
	b, err := json.Marshal(sp.Tags)
	if err != nil {
		return errors.WithStack(err)
	}

	_, err = connection.Exec(`UPDATE poststaging SET tags = ?, image = ?, title = ?, body = ? WHERE id = ?;`, b, sp.Image, sp.Title, sp.Body, sp.ID)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

//Delete (StagingPost) deletes a blog post
func (sp *StagingPost) Delete() error {
	_, err := connection.Exec(`DELETE FROM poststaging WHERE id = ?;`, sp.ID)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func parseBlogStagingRows(sps *StagingPosts, r *sql.Rows) error {
	for r.Next() {
		var sp StagingPost
		var tags string

		err := r.Scan(&sp.ID, &tags, &sp.Image, &sp.Title, &sp.Body)
		if err != nil {
			return errors.WithStack(err)
		}

		err = json.Unmarshal([]byte(tags), &sp.Tags)
		if err != nil {
			return errors.WithStack(err)
		}

		*sps = append(*sps, sp)
	}

	err := r.Err()
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}
