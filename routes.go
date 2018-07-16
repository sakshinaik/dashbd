package main

import (
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func (a *App) appRoutes() []Route {
	return Routes{
		Route{
			"Index",
			"GET",
			"/",
			a.Index,
		},
		Route{
			"UserIndex",
			"GET",
			"/users",
			a.UsersIndex,
		}, Route{
			"UserProfile",
			"GET",
			"/users/{id}",
			a.UserProfile,
		},
		Route{
			"TeamIndex",
			"GET",
			"/teams",
			a.TeamIndex,
		},
		Route{
			"Login",
			"POST",
			"/login",
			a.UserLogin,
		},
		/*Route{
			"TIndex",
			"GET",
			"/todos/{todoId}",
			a.TodoShow,
		},*/
	}
}
