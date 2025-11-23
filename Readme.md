# Go Playground Server

A small Go web server designed to demonstrate a set of common web development features: Go HTML template rendering, serving static assets, and implementing simple JSON APIs.

It serves a single-page application at the root path (`/`) which utilizes the two JSON APIs.

## Features

* **Go HTML Templates**: Renders the main `index.html` page using Go's `html/template` package.
* **Static Assets**: Serves CSS and JavaScript files from the `/static/` path.
* **Greeting API**: A simple JSON API available at `/greet`.
* **Double API**: A new JSON API available at `/double` for integer calculation.


### Prerequisites

* Go (The project specifies `go 1.25.0` in `go.mod`, but any recent version will work).

### 1. Run the server

You can run the server directly using the `go run` command:

```bash
go run ./cmd/server
```

## LICENSE 
This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.