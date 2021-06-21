package mid

import (
	"log"
	"net/http"
	"time"

	//gin
	"github.com/gorilla/sessions"
)

var (
	// key must be 16, 24 or 32 bytes long (AES-128, AES-192 or AES-256)
	key   = []byte("super-secret-key")
	store = sessions.NewCookieStore(key)
)

// Chain applies middlewares to a http.HandlerFunc
type Middleware func(http.HandlerFunc) http.HandlerFunc

func Chain(f http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	for _, m := range middlewares {
		f = m(f)
	}
	return f
}

func AuthRequired() Middleware {
	// Create a new Middleware
	return func(f http.HandlerFunc) http.HandlerFunc {
		// Define the http.HandlerFunc
		return func(w http.ResponseWriter, r *http.Request) {
			session, _ := store.Get(r, "cookie-name")
			// Check if user is authenticated
			if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
				http.Error(w, "Forbidden", http.StatusForbidden)
				return
			}
			// Call the next middleware/handler in chain
			f(w, r)
		}
	}
}

// Logging logs all requests with its path and the time it took to process
func Logging() Middleware {

	// Create a new Middleware
	return func(f http.HandlerFunc) http.HandlerFunc {

		// Define the http.HandlerFunc
		return func(w http.ResponseWriter, r *http.Request) {

			// Do middleware things
			start := time.Now()
			defer func() {
				log.Println(r.URL.Path, time.Since(start))
				//log.Println(r.URL.Path, time.Since(start))
			}()

			// Call the next middleware/handler in chain
			f(w, r)
		}
	}
}

/*
func authMiddleWare(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, _ := store.Get(r, "cookie-name")
		// Check if user is authenticated
		if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}


		next(w, r)
	}
}
func loggingMiddleWare(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL.Path)
		next(w, r)
	}
}
*/
