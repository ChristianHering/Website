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
	mux.Handle("/page/{page:[0-9]+}", middlewares.ThenFunc(pageHandler))
	mux.Handle("/post/{id:[0-9]+}", middlewares.ThenFunc(postHandler))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	var posts utils.Posts

	err := posts.GetNewestPosts("7")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		panic(err)
	}

	data := utils.PostsData{
		Posts:       posts,
		CurrentPage: 1,
	}

	err = templates.Templates.ExecuteTemplate(w, "blogIndex.html", data)
	if err != nil {
		panic(err)
	}
}

func pageHandler(w http.ResponseWriter, r *http.Request) {
	var posts utils.Posts

	page, err := strconv.Atoi(mux.Vars(r)["page"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = posts.GetPostRange((page - 1) * 6)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		panic(err)
	}

	data := utils.PostsData{
		Posts:       posts,
		CurrentPage: page,
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

	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		panic(err)
	}

	if len(posts) == 2 {
		if posts[0].ID == id {
			posts = append([]utils.Post{utils.Post{}}, posts...)
		} else {
			posts = append(posts, utils.Post{})
		}
	}

	data := utils.PostData{}

	data.LastPost = posts[0]
	data.Post = posts[1]
	data.NextPost = posts[2]

	err = templates.Templates.ExecuteTemplate(w, "blogPost.html", data)
	if err != nil {
		panic(err)
	}
}
