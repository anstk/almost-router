package almost

import (
	"log"
	"net/http"
)

type Almost struct {
	routes []route
}

type route struct {
	Path    string
	Method  string
	Handler http.HandlerFunc
}

func Router() *Almost {
	return &Almost{}
}

func (al *Almost) Route(method, path string, handlerFunc http.HandlerFunc) {
	al.routes = append(al.routes, route{
		Method:  method,
		Path:    path,
		Handler: handlerFunc,
	})
}

func (al *Almost) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	// IDEA - Use a real and efficient algorithm for route matching.
	// Like PruningRadixTrie - it's up to 1000x faster than an ordinary Radix Trie.  https://github.com/wolfgarbe/PruningRadixTrie
	// Go Chi uses Radix Trie https://github.com/go-chi/chi/blob/master/tree.go
	// Fasthttp router uses Radix Trie  https://github.com/fasthttp/router
	// Atreugo uses Fasthttp router (Radix Trie) https://github.com/savsgio/atreugo/blob/master/router.go
	// Could these routers be faster? Which algorithm?

	for _, tr := range al.routes {
		match := tr.match(r)
		if !match {
			continue
		}

		tr.Handler.ServeHTTP(w, r)
		return
	}
	http.NotFound(w, r)
}

func (al *Almost) Start(addr string) {
	err := http.ListenAndServe(addr, al)
	if err != nil {
		log.Fatal("[ FATAL] Almost Start", err)
	}
}

func (ar *route) match(r *http.Request) bool {

	// TODO extract parameters
	// TODO use context to pass parameters to Handler

	// simple method compare
	if r.Method != ar.Method {
		return false
	}

	// simple string compare
	if r.URL.Path != ar.Path {
		return false
	}

	return true
}
