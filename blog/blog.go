package blog

import (
	"net/http"
	"strconv"

	"github.com/ChristianHering/Website/utils"
	"github.com/ChristianHering/Website/utils/middleware"
	"github.com/ChristianHering/Website/utils/templates"
	"github.com/gorilla/mux"
	"github.com/justinas/alice"
)

//Run Serves a development blog on the blog subdomain
func Run(m *mux.Router) {
	mux := m.Host("blog.christianhering.com").Subrouter()

	middlewares := alice.New(middleware.ErrorHandler)

	mux.Handle("/", middlewares.ThenFunc(indexHandler))
	mux.Handle("/page/{id:[0-9]+}", middlewares.ThenFunc(pageHandler))
	mux.Handle("/post/{id:[0-9]+}", middlewares.ThenFunc(postHandler))

	return
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	var posts utils.Posts

	err := posts.GetNewestPosts("6")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		panic(err)
	}

	data := utils.PostsData{
		Posts:       posts,
		PageCount:   int((utils.BlogRowCount - 1) / 2),
		CurrentPage: 1,
	}

	err = templates.Templates.ExecuteTemplate(w, "blogIndex.html", data)
	if err != nil {
		panic(err)
	}
}

func pageHandler(w http.ResponseWriter, r *http.Request) {
	var posts utils.Posts

	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		panic(err)
	}

	var limit = 6
	var total = id * 6

	//This is so the last page contains the correct number
	//of blog posts. Eg. If there are 7 blog posts, the 2nt
	//page should contain 1 post, not 6. Just a minor thing
	if (id * 6) > utils.BlogRowCount {
		limit = utils.BlogRowCount % 6
		total = utils.BlogRowCount
	}

	err = posts.GetPostRange(strconv.Itoa(limit), strconv.Itoa(total))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		panic(err)
	}

	data := utils.PostsData{
		Posts:       posts,
		PageCount:   int((utils.BlogRowCount-1)/6) + 1,
		CurrentPage: id,
	}

	err = templates.Templates.ExecuteTemplate(w, "blogIndex.html", data)
	if err != nil {
		panic(err)
	}
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	var posts utils.Posts

	err := posts.GetSurroundingPosts(mux.Vars(r)["id"])
	if err != nil {
		panic(err)
	}

	data := utils.PostData{}

	switch len(posts) {
	case 0:
		err := posts.GetOldestPosts("2")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			panic(err)
		}

		data.LastPost = utils.Post{}
		data.Post = posts[0]
		data.NextPost = posts[1]
	case 2:
		data.LastPost = posts[0]
		data.Post = posts[1]
		data.NextPost = utils.Post{}
	case 3:
		data.LastPost = posts[0]
		data.Post = posts[1]
		data.NextPost = posts[2]
	}

	err = templates.Templates.ExecuteTemplate(w, "blogPost.html", data)
	if err != nil {
		panic(err)
	}
}
