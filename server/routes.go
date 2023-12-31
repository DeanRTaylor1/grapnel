package server

import (
	"net/http"

	"github.com/DeanRTaylor1/deans-site/config"
	"github.com/DeanRTaylor1/deans-site/handlers"
	"github.com/go-chi/chi/v5"
)

func (s *Server) RegisterRoutes(router *chi.Mux) {

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		handlers.ServeHome(w, r, s.Logger)
	})

	router.Get("/styles/output.css", func(w http.ResponseWriter, r *http.Request) {
		handlers.ServeCss(w, r, s.Logger)
	})

	router.Get("/faq", func(w http.ResponseWriter, r *http.Request) {
		handlers.ServeFaq(w, r, s.Logger)
	})

	router.Get("/about", func(w http.ResponseWriter, r *http.Request) {
		handlers.ServeAbout(w, r, s.Logger)
	})

	router.Get("/fonts/*", func(w http.ResponseWriter, r *http.Request) {
		handlers.ServeFonts(w, r, s.Logger, s.Config)
	})

	router.Get("/images/icons/*", func(w http.ResponseWriter, r *http.Request) {
		handlers.ServeFavicon(w, r, s.Logger)
	})

	router.Get("/images/*", func(w http.ResponseWriter, r *http.Request) {
		handlers.ServeImages(w, r, s.Logger)
	})

	router.Get("/manifest/images/icons/*", func(w http.ResponseWriter, r *http.Request) {
		handlers.ServeIcons(w, r, s.Logger)
	})

	router.Get("/manifest/*", func(w http.ResponseWriter, r *http.Request) {
		handlers.ServeManifest(w, r, s.Logger)
	})

	router.Get("/scripts/*", func(w http.ResponseWriter, r *http.Request) {
		handlers.ServeScripts(w, r, s.Logger, config.Env)
	})

	router.Get("/blogs", func(w http.ResponseWriter, r *http.Request) {
		handlers.ServeBlog(w, r, s.Logger)
	})

	router.Get("/blogs/data", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetBlogs(w, r, s.Logger)
	})

	router.Get("/blogs/{id}", func(w http.ResponseWriter, r *http.Request) {
		blogID := chi.URLParam(r, "id")

		handlers.GetBlogByID(w, r, s.Logger, blogID)
	})

	router.Get("/contact", func(w http.ResponseWriter, r *http.Request) {
		handlers.ServeContact(w, r, s.Logger)
	})

	// // Serve static files
	// staticDir := getStaticDir()
	// staticFileServer := http.FileServer(http.Dir(staticDir))
	// staticRoute := "/static/"

	// router.Handle(staticRoute+"*", http.StripPrefix(staticRoute, staticFileServer))
}

// func getStaticDir() string {
// 	dir, err := filepath.Abs(filepath.Join(filepath.Dir("."), "static"))
// 	if err != nil {
// 		// Handle error
// 		panic(err)
// 	}
// 	return dir
// }
