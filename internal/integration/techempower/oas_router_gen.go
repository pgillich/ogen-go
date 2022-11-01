// Code generated by ogen, DO NOT EDIT.

package techempower

import (
	"net/http"
	"strings"
)

// ServeHTTP serves http request as defined by OpenAPI v3 specification,
// calling handler that matches the path or returning not found error.
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	elem := r.URL.Path
	if prefix := s.cfg.Prefix; len(prefix) > 0 {
		if strings.HasPrefix(elem, prefix) {
			// Cut prefix from the path.
			elem = strings.TrimPrefix(elem, prefix)
		} else {
			// Prefix doesn't match.
			s.notFound(w, r)
			return
		}
	}
	if len(elem) == 0 {
		s.notFound(w, r)
		return
	}

	// Static code generated router with unwrapped path search.
	switch {
	default:
		if len(elem) == 0 {
			break
		}
		switch elem[0] {
		case '/': // Prefix: "/"
			if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
				elem = elem[l:]
			} else {
				break
			}

			if len(elem) == 0 {
				break
			}
			switch elem[0] {
			case 'c': // Prefix: "cached-worlds"
				if l := len("cached-worlds"); len(elem) >= l && elem[0:l] == "cached-worlds" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					// Leaf node.
					switch r.Method {
					case "GET":
						s.handleCachingRequest([0]string{}, w, r)
					default:
						s.notAllowed(w, r, "GET")
					}

					return
				}
			case 'd': // Prefix: "db"
				if l := len("db"); len(elem) >= l && elem[0:l] == "db" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					// Leaf node.
					switch r.Method {
					case "GET":
						s.handleDBRequest([0]string{}, w, r)
					default:
						s.notAllowed(w, r, "GET")
					}

					return
				}
			case 'j': // Prefix: "json"
				if l := len("json"); len(elem) >= l && elem[0:l] == "json" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					// Leaf node.
					switch r.Method {
					case "GET":
						s.handleJSONRequest([0]string{}, w, r)
					default:
						s.notAllowed(w, r, "GET")
					}

					return
				}
			case 'q': // Prefix: "queries"
				if l := len("queries"); len(elem) >= l && elem[0:l] == "queries" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					// Leaf node.
					switch r.Method {
					case "GET":
						s.handleQueriesRequest([0]string{}, w, r)
					default:
						s.notAllowed(w, r, "GET")
					}

					return
				}
			case 'u': // Prefix: "updates"
				if l := len("updates"); len(elem) >= l && elem[0:l] == "updates" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					// Leaf node.
					switch r.Method {
					case "GET":
						s.handleUpdatesRequest([0]string{}, w, r)
					default:
						s.notAllowed(w, r, "GET")
					}

					return
				}
			}
		}
	}
	s.notFound(w, r)
}

// Route is route object.
type Route struct {
	name        string
	operationID string
	count       int
	args        [0]string
}

// Name returns ogen operation name.
//
// It is guaranteed to be unique and not empty.
func (r Route) Name() string {
	return r.name
}

// OperationID returns OpenAPI operationId.
func (r Route) OperationID() string {
	return r.operationID
}

// Args returns parsed arguments.
func (r Route) Args() []string {
	return r.args[:r.count]
}

// FindRoute finds Route for given method and path.
func (s *Server) FindRoute(method, path string) (r Route, _ bool) {
	var (
		args = [0]string{}
		elem = path
	)
	r.args = args
	if elem == "" {
		return r, false
	}

	// Static code generated router with unwrapped path search.
	switch {
	default:
		if len(elem) == 0 {
			break
		}
		switch elem[0] {
		case '/': // Prefix: "/"
			if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
				elem = elem[l:]
			} else {
				break
			}

			if len(elem) == 0 {
				break
			}
			switch elem[0] {
			case 'c': // Prefix: "cached-worlds"
				if l := len("cached-worlds"); len(elem) >= l && elem[0:l] == "cached-worlds" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					switch method {
					case "GET":
						// Leaf: Caching
						r.name = "Caching"
						r.operationID = "Caching"
						r.args = args
						r.count = 0
						return r, true
					default:
						return
					}
				}
			case 'd': // Prefix: "db"
				if l := len("db"); len(elem) >= l && elem[0:l] == "db" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					switch method {
					case "GET":
						// Leaf: DB
						r.name = "DB"
						r.operationID = "DB"
						r.args = args
						r.count = 0
						return r, true
					default:
						return
					}
				}
			case 'j': // Prefix: "json"
				if l := len("json"); len(elem) >= l && elem[0:l] == "json" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					switch method {
					case "GET":
						// Leaf: JSON
						r.name = "JSON"
						r.operationID = "json"
						r.args = args
						r.count = 0
						return r, true
					default:
						return
					}
				}
			case 'q': // Prefix: "queries"
				if l := len("queries"); len(elem) >= l && elem[0:l] == "queries" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					switch method {
					case "GET":
						// Leaf: Queries
						r.name = "Queries"
						r.operationID = "Queries"
						r.args = args
						r.count = 0
						return r, true
					default:
						return
					}
				}
			case 'u': // Prefix: "updates"
				if l := len("updates"); len(elem) >= l && elem[0:l] == "updates" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					switch method {
					case "GET":
						// Leaf: Updates
						r.name = "Updates"
						r.operationID = "Updates"
						r.args = args
						r.count = 0
						return r, true
					default:
						return
					}
				}
			}
		}
	}
	return r, false
}