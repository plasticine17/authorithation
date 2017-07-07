package Network

import "github.com/gorilla/mux"

func newRouter() *mux.Router {

	var router = mux.NewRouter().StrictSlash(true)

	for _, route := range routes {
		router.
			Methods(route.Methot).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)

	}

	return router
}
