# Almost Router

<!--
https://shields.io/category/build
-->

[![build](https://github.com/anstk/almost-router/actions/workflows/build.yml/badge.svg)](https://github.com/anstk/almost-router/actions/workflows/build.yml)
[![lint](https://github.com/anstk/almost-router/actions/workflows/lint.yml/badge.svg)](doc/lint.md)
[![lint](https://github.com/anstk/almost-router/actions/workflows/misspell.yml/badge.svg)](https://github.com/anstk/almost-router/actions/workflows/misspell.yml)

For study purposes only.

Just a very simple "router", a WIP (work in progress), to evolve the idea in the future and maybe create a usable router.


Next TODOs
 - Extract parameters
 - Use context to pass parameters to Handler.
 - Use a real and efficient algorithm for route matching.
   - Like PruningRadixTrie - it's up to 1000x faster than an ordinary Radix Trie.  https://github.com/wolfgarbe/PruningRadixTrie
   - Go Chi uses Radix Trie https://github.com/go-chi/chi/blob/master/tree.go
   - Fasthttp router uses Radix Trie  https://github.com/fasthttp/router
   - Atreugo uses Fasthttp router (Radix Trie) https://github.com/savsgio/atreugo/blob/master/router.go
   - Could these routers be faster? Which algorithm?



See example. 
Almost Router can only match an exact word, no patterns.

```go
package main

import (
	"fmt"
	"net/http"

	almost "github.com/anstk/almost-router"
)

func main() {

	r := almost.Router()

	r.Route("GET", "/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Index page"))
	})

	r.Route("GET", "/user", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("User page"))
	})

	r.Route("GET", "/play", play)

	r.Start("localhost:8000")

}

func play(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "played")
}
```
