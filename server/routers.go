/*
 * Takí Project
 *
 * IF683 Takí project
 *
 * API version: 1.0.0
 * Contact: mvgmb@cin.ufpe.br
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package server

import (
	"database/sql"
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var db *sql.DB

func NewRouter(_db *sql.DB) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	db = _db

	return router
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},

	Route{
		"StoreStoreIdListListIdDelete",
		strings.ToUpper("Delete"),
		"/store/{storeId}/list/{listId}",
		StoreStoreIDListListIDDelete,
	},

	Route{
		"StoreStoreIdListListIdGet",
		strings.ToUpper("Get"),
		"/store/{storeId}/list/{listId}",
		StoreStoreIdListListIdGet,
	},

	Route{
		"StoreStoreIdListListIdPut",
		strings.ToUpper("Put"),
		"/store/{storeId}/list/{listId}",
		StoreStoreIdListListIdPut,
	},

	Route{
		"StoreStoreIdListListIdRouteGet",
		strings.ToUpper("Get"),
		"/store/{storeId}/list/{listId}/route",
		StoreStoreIdListListIdRouteGet,
	},

	Route{
		"StoreStoreIdListNewPost",
		strings.ToUpper("Post"),
		"/store/{storeId}/list/new",
		StoreStoreIdListNewPost,
	},

	Route{
		"StoreStoreIdListsGet",
		strings.ToUpper("Get"),
		"/store/{storeId}/lists",
		StoreStoreIdListsGet,
	},

	Route{
		"StoreStoreIdMapGet",
		strings.ToUpper("Get"),
		"/store/{storeId}/map",
		StoreStoreIdMapGet,
	},

	Route{
		"StoreStoreIdProductsGet",
		strings.ToUpper("Get"),
		"/store/{storeId}/products",
		StoreStoreIdProductsGet,
	},

	Route{
		"StoresGet",
		strings.ToUpper("Get"),
		"/stores",
		StoresGet,
	},

	Route{
		"UserGet",
		strings.ToUpper("Get"),
		"/user",
		UserGet,
	},

	Route{
		"UserNewPost",
		strings.ToUpper("Post"),
		"/user/new",
		UserNewPost,
	},
}
