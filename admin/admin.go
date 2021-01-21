package admin

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/ChristianHering/Website/utils"
	"github.com/ChristianHering/Website/utils/middleware"
	"github.com/ChristianHering/Website/utils/templates"
	"github.com/gorilla/mux"
	"github.com/justinas/alice"
	"github.com/pkg/errors"
)

//Run Serves an administrative panel for viewing
//stats and managing content on the admin subdomain
func Run(m *mux.Router) {
	mux := m.Host("admin.christianhering.com").Subrouter()

	//Non-Authenticated Handlers

	middlewares := alice.New(middleware.ErrorHandler)

	mux.Handle("/", middlewares.ThenFunc(indexHandler))

	mux.Handle("/login", middlewares.ThenFunc(utils.LoginHandler))
	mux.Handle("/logout", middlewares.ThenFunc(utils.LogoutHandler))
	mux.Handle("/callback", middlewares.ThenFunc(utils.CallbackHandler))

	//Authenticated Handlers

	middlewaresWithAuth := alice.New(middleware.ErrorHandler, middleware.AuthenticationHandler)

	mux.Handle("/dashboard", middlewaresWithAuth.ThenFunc(dashboardHandler))

	mux.Handle("/error", middlewaresWithAuth.ThenFunc(errorHandler))
	mux.Handle("/errorDelete", middlewaresWithAuth.ThenFunc(errorDeletionHandler))
	mux.Handle("/errorDeleteType", middlewaresWithAuth.ThenFunc(errorTypeDeletionHandler))

	mux.Handle("/posts", middlewaresWithAuth.ThenFunc(postsHandler))
	mux.Handle("/postCreate", middlewaresWithAuth.ThenFunc(postCreateHandler))
	mux.Handle("/postUpdate", middlewaresWithAuth.ThenFunc(postUpdateHandler))
	mux.Handle("/postDelete", middlewaresWithAuth.ThenFunc(postDeleteHandler))

	mux.Handle("/stagedPosts", middlewaresWithAuth.ThenFunc(stagedPostsHandler))
	mux.Handle("/stagedPostCreate", middlewaresWithAuth.ThenFunc(stagedPostCreateHandler))
	mux.Handle("/stagedPostUpdate", middlewaresWithAuth.ThenFunc(stagedPostUpdateHandler))
	mux.Handle("/stagedPostDelete", middlewaresWithAuth.ThenFunc(stagedPostDeleteHandler))

	return
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	err := templates.Templates.ExecuteTemplate(w, "adminIndex.html", nil)
	if err != nil {
		panic(errors.WithStack(err))
	}
}

func dashboardHandler(w http.ResponseWriter, r *http.Request) {
	err := templates.Templates.ExecuteTemplate(w, "adminDashboard.html", nil)
	if err != nil {
		panic(errors.WithStack(err))
	}
}

func errorHandler(w http.ResponseWriter, r *http.Request) {
	var e utils.Errors

	err := e.Read("3")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		panic(errors.WithStack(err))
	}

	json.NewEncoder(w).Encode(e)
}

func errorDeletionHandler(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		panic(errors.WithStack(err))
	}

	var e utils.Error

	err = json.Unmarshal(b, &e)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		panic(errors.WithStack(err))
	}

	err = e.Delete()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		panic(errors.WithStack(err))
	}

	w.WriteHeader(http.StatusOK)
}

func errorTypeDeletionHandler(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		panic(errors.WithStack(err))
	}

	var e utils.Error

	err = json.Unmarshal(b, &e)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		panic(errors.WithStack(err))
	}

	err = e.DeleteErrorType()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		panic(errors.WithStack(err))
	}

	w.WriteHeader(http.StatusOK)
}

func postsHandler(w http.ResponseWriter, r *http.Request) {
	var posts utils.Posts

	err := posts.GetNewestPosts("6")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		panic(errors.WithStack(err))
	}

	json.NewEncoder(w).Encode(posts)
}

func postCreateHandler(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var post utils.Post

	err = json.Unmarshal(b, &post)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = post.Create()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		panic(errors.WithStack(err))
	}
}

func postUpdateHandler(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var post utils.Post

	err = json.Unmarshal(b, &post)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = post.Update()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}

func postDeleteHandler(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var post utils.Post

	err = json.Unmarshal(b, &post)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = post.Delete()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}

func stagedPostsHandler(w http.ResponseWriter, r *http.Request) {
	var posts utils.StagingPosts

	err := posts.Read("6")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		panic(errors.WithStack(err))
	}

	json.NewEncoder(w).Encode(posts)
}

func stagedPostCreateHandler(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var stagedPost utils.StagingPost

	err = json.Unmarshal(b, &stagedPost)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = stagedPost.Create()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		panic(errors.WithStack(err))
	}
}

func stagedPostUpdateHandler(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var stagedPost utils.StagingPost

	err = json.Unmarshal(b, &stagedPost)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = stagedPost.Update()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}

func stagedPostDeleteHandler(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var stagedPost utils.StagingPost

	err = json.Unmarshal(b, &stagedPost)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = stagedPost.Delete()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}
