package server

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
	"strings"
)

func RunHttpServer(addr string, AddRouteFn func(chi.Router)) {
	addr = parseAddr(addr)
	router := chi.NewRouter()

	setupMiddleware(router)

	AddRouteFn(router)

	printStartInfo(addr, router)

	err := http.ListenAndServe(addr, router)
	if err != nil {
		panic(fmt.Sprintf("\u001B[31mHttpserver failed to start:%s \u001B[m", err))
	}
}

func setupMiddleware(router chi.Router) {
	router.Use(middleware.Recoverer)
}

func parseAddr(addr string) string {
	if strings.HasPrefix(addr, ":") {
		addr = "127.0.0.1" + addr
	}
	return addr
}

func printStartInfo(addr string, router chi.Router) {
	_ = chi.Walk(router, func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		fmt.Printf("\033[32m%-10s\033[m - \033[1;4m%-10s\033[m    (has \033[1m%d\033[m middleware)\n", fmt.Sprintf("[%s]", method), route, len(middlewares))
		return nil
	})
	fmt.Printf("\n\033[1mHttpserver Starts Running: \033[m\033[1;4;32m%s\033[m\n", addr)
}
