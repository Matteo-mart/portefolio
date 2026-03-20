package route

import (
	"fmt"

	"github.com/gorilla/mux"
)

func ListRoute(r *mux.Router) {
	fmt.Println("--- Routes ---")
	r.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		t, err := route.GetPathTemplate()
		if err == nil && t != "" {
			fmt.Printf("Path: %s\n", t)
		}
		return nil
	})
	fmt.Println("---------------------------")
}
