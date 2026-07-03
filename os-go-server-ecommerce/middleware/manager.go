package middleware

import "net/http"

type Middleware func(http.Handler) http.Handler

type Manager struct {
	globalMiddlewares []Middleware
}

func NewManager() *Manager {
	return &Manager{
		globalMiddlewares: make([]Middleware, 0),
	}
}

func (mngr *Manager) Use(middlewares ...Middleware) {
	mngr.globalMiddlewares = append(mngr.globalMiddlewares, middlewares...)
}

func (mngr *Manager) With(next http.Handler, middlewares ...Middleware) http.Handler {
	n := next

	// logger(http.handlerFunc(GetProducts))
	for _, middleware := range middlewares {
		n = middleware(n)
	}

	// globalMiddleware(logger(http.handleFunc(GetProducts)))
	for _, globalMiddleware := range mngr.globalMiddlewares {
		n = globalMiddleware(n)
	}
	return n
}
