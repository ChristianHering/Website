package utils

import (
	"database/sql"
	"encoding/json"
	"time"

	_ "github.com/lib/pq"
	"github.com/pkg/errors"
)

//Posts is a slice of 'Post'
type Posts []Post

//Post represents a single blog post from our 'posts' table
type Post struct {
	ID    int
	Date  time.Time
	Tags  []string
	Image string
	Title string
	Body  string
}

//PostsData is used for blog page handlers
type PostsData struct {
	Posts       Posts
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

	_, err = connection.Exec(`INSERT INTO posts (date, tags, image, title, body) VALUES ($1, $2, $3, $4, $5);`, time.Now().Format("2006-01-02 15:04:05.000000"), b, p.Image, p.Title, p.Body)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

//Read (Post) retrieves a blog post from our 'posts' table
func (p *Post) Read() error {
	r := connection.QueryRow(`SELECT * FROM posts WHERE id = $1 LIMIT 1;`, p.ID)

	var date string
	var tags string

	err := r.Scan(&p.ID, &date, &tags, &p.Image, &p.Title, &p.Body)
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

	_, err = connection.Exec(`UPDATE posts SET tags = $1, image = $2, title = 3, body = $4 WHERE id = $5;`, b, p.Image, p.Title, p.Body, p.ID)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

//Delete (Post) deletes a blog post
func (p *Post) Delete() error {
	_, err := connection.Exec(`DELETE FROM posts WHERE id = $1;`, p.ID)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

//GetNewestPosts (Post) retrieves 'limit' most recent posts
//
//This is used to populate blog.ChristianHering.com's home page
func (ps *Posts) GetNewestPosts(limit string) error {
	r, err := connection.Query(`SELECT * FROM posts ORDER BY id DESC LIMIT $1;`, limit)
	if err != nil {
		return errors.WithStack(err)
	}
	defer r.Close()

	return parseBlogRows(ps, r)
}

//GetPostRange (Post) returns a range of blog posts starting at offset
//
//Returns (len(postsInTheDB) % 7) posts if there aren't at least 7 posts after 'offset'
func (ps *Posts) GetPostRange(offset int) error {
	r, err := connection.Query(`SELECT * FROM posts ORDER BY id DESC OFFSET $1 LIMIT 7;`, offset)
	if err != nil {
		return errors.WithStack(err)
	}
	defer r.Close()

	return parseBlogRows(ps, r)
} //`SELECT * FROM (SELECT * FROM (SELECT * FROM posts ORDER BY id DESC LIMIT $1) t ORDER BY id ASC LIMIT $2) st ORDER BY id DESC;`

//GetSurroundingPosts (Post) returns the
//queried post and the post on either side
func (ps *Posts) GetSurroundingPosts(postID string) error {
	r, err := connection.Query("WITH init AS (SELECT * FROM posts WHERE id = $1)((SELECT posts.* FROM posts CROSS JOIN init WHERE posts.id <= init.id ORDER BY posts.id DESC LIMIT 2) UNION ALL (SELECT posts.* FROM posts CROSS JOIN init WHERE posts.id > init.id ORDER BY posts.id LIMIT 1)) ORDER BY id ASC;" /*SELECT * FROM posts t WHERE id >= (SELECT id FROM posts WHERE id <= $1 ORDER BY id DESC LIMIT 1) ORDER BY id ASC LIMIT 3;"*/, postID)
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

		err := r.Scan(&p.ID, &date, &tags, &p.Image, &p.Title, &p.Body)
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
