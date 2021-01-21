package utils

import (
	"database/sql"
	"encoding/json"
	"time"

	_ "github.com/go-sql-driver/mysql" //SQL Driver for MySQL/MariaDB
	"github.com/pkg/errors"
)

//BlogRowCount holds the number of rows in our 'posts' table
var BlogRowCount int

//Posts is a slice of 'Post'
type Posts []Post

//Post represents a single blog post from our 'posts' table
type Post struct {
	ID           int
	Date         time.Time
	Tags         []string
	Image        string
	Title        string
	BodyShowcase string
	Body         string
}

//PostsData is used for blog page handlers
type PostsData struct {
	Posts       Posts
	PageCount   int
	CurrentPage int
}

//PostData is used for blog post handlers
type PostData struct {
	LastPost Post
	Post     Post
	NextPost Post
}

//Create (Post) adds a new blog post to our 'posts' table
func (p *Post) Create() error {
	b, err := json.Marshal(p.Tags)
	if err != nil {
		return errors.WithStack(err)
	}

	_, err = connection.Exec(`INSERT INTO posts (date, tags, image, title, body) VALUES (?, ?, ?, ?, ?);`, time.Now().Format("2006-01-02 15:04:05.000000"), b, p.Image, p.Title, p.Body)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

//Read (Post) retrieves a blog post from our 'posts' table
func (p *Post) Read() error {
	r := connection.QueryRow(`SELECT * FROM posts WHERE id = ? LIMIT 1;`, p.ID)

	var date string
	var tags string

	err := r.Scan(&p.ID, &date, &tags, &p.Image, &p.Title, &p.BodyShowcase, &p.Body)
	if err != nil {
		return errors.WithStack(err)
	}

	p.Date, err = time.Parse("2006-01-02 15:04:05.000000", date)
	if err != nil {
		return errors.WithStack(err)
	}

	err = json.Unmarshal([]byte(tags), &p.Tags)
	if err != nil {
		return errors.WithStack(err)
	}

	err = r.Err()
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

//Update (Post) modifies a blog post entry
func (p *Post) Update() error {
	b, err := json.Marshal(p.Tags)
	if err != nil {
		return errors.WithStack(err)
	}

	_, err = connection.Exec(`UPDATE posts SET tags = ?, image = ?, title = ?, body = ? WHERE id = ?;`, b, p.Image, p.Title, p.Body, p.ID)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

//Delete (Post) deletes a blog post
func (p *Post) Delete() error {
	_, err := connection.Exec(`DELETE FROM posts WHERE id = ?;`, p.ID)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

//GetNewestPosts (Post) retrieves 'limit' most recent posts
//
//This is used to populate blog.ChristianHering.com's home page
func (ps *Posts) GetNewestPosts(limit string) error {
	r, err := connection.Query(`SELECT * FROM posts ORDER BY id DESC LIMIT ?;`, limit)
	if err != nil {
		return errors.WithStack(err)
	}
	defer r.Close()

	return parseBlogRows(ps, r)
}

//GetOldestPosts (Post) retrieves 'limit' most recent posts
//
//This is used to populate blog.ChristianHering.com's home page
func (ps *Posts) GetOldestPosts(limit string) error {
	r, err := connection.Query(`SELECT * FROM posts ORDER BY id ASC LIMIT ?;`, limit)
	if err != nil {
		return errors.WithStack(err)
	}
	defer r.Close()

	return parseBlogRows(ps, r)
}

//GetPostRange (Post) returns a range of blog posts
//
//total = limit + (# of posts that have been viewed)
//
//This is used to dynamically load X more blog posts
//after Y blog posts have already been retrieved without
//querying for X+Y blog posts then throwing out Y of them.
func (ps *Posts) GetPostRange(limit string, total string) error {
	r, err := connection.Query(`SELECT * FROM (SELECT * FROM posts ORDER BY id DESC LIMIT ?) t ORDER BY id ASC LIMIT ?;`, total, limit)
	if err != nil {
		return errors.WithStack(err)
	}
	defer r.Close()

	return parseBlogRows(ps, r)
}

//GetSurroundingPosts (Post) returns the
//queried post and the post on either side
func (ps *Posts) GetSurroundingPosts(postID string) error {
	r, err := connection.Query(`SELECT * FROM posts t WHERE id >= (SELECT id FROM `+"`posts`"+` lo WHERE id <= ? ORDER BY id DESC LIMIT 1,1) ORDER BY id ASC LIMIT 0,3;`, postID)
	if err != nil {
		return errors.WithStack(err)
	}
	defer r.Close()

	return parseBlogRows(ps, r)
}

func parseBlogRows(ps *Posts, r *sql.Rows) error {
	for r.Next() {
		var p Post
		var date string
		var tags string

		err := r.Scan(&p.ID, &date, &tags, &p.Image, &p.Title, &p.BodyShowcase, &p.Body)
		if err != nil {
			return errors.WithStack(err)
		}

		p.Date, err = time.Parse("2006-01-02 15:04:05.000000", date)
		if err != nil {
			return errors.WithStack(err)
		}

		err = json.Unmarshal([]byte(tags), &p.Tags)
		if err != nil {
			return errors.WithStack(err)
		}

		*ps = append(*ps, p)
	}

	err := r.Err()
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}
