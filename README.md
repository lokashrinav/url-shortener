Here is the formatted `README.md` file for your URL shortener written in Go:

```markdown
# URL Shortener

A simple URL shortener written in Go, where routes are mapped to shortened URLs. If a user accesses any route that matches a path in the map, they are redirected to the corresponding URL.

## Features

- Map routes to shortened URLs
- Automatically redirect users to the intended URL based on predefined mappings
- Lightweight and fast

## How It Works

The application keeps a hardcoded map of shortened paths to original URLs. When a user accesses a path, it checks if the path exists in the map. If a match is found, the server redirects the user to the mapped URL; otherwise, a 404 page is returned.

### Example Redirects

- Accessing `http://localhost:8080/go` will redirect to `https://golang.org`.
- Accessing `http://localhost:8080/openai` will redirect to `https://openai.com`.
- Accessing an unmapped path, like `http://localhost:8080/unknown`, will return a 404 error.

## Getting Started

### Prerequisites

- Go installed on your local machine (Go 1.16 or higher recommended)

### Installation

1. **Clone the repository:**
   ```bash
   git clone https://github.com/your-username/url-shortener.git
   cd url-shortener
   ```

2. **Run the application:**
   ```bash
   go run main.go
   ```

3. **Open your browser and navigate to:**
   ```
   http://localhost:8080/<path>
   ```
   For example: `http://localhost:8080/go`

## Example Code

Here's a simplified version of the code used to handle redirection:

```go
package main

import (
    "net/http"
)

var urlMap = map[string]string{
    "/go":       "https://golang.org",
    "/openai":   "https://openai.com",
    "/github":   "https://github.com",
}

func redirectHandler(w http.ResponseWriter, r *http.Request) {
    target, exists := urlMap[r.URL.Path]
    if exists {
        http.Redirect(w, r, target, http.StatusFound)
    } else {
        http.NotFound(w, r)
    }
}

func main() {
    http.HandleFunc("/", redirectHandler)
    http.ListenAndServe(":8080", nil)
}
```

## Configuration

To customize the mappings, modify the `urlMap` variable in `main.go` to include the paths and corresponding URLs you’d like to redirect.

## License

This project is licensed under the MIT License - see the LICENSE file for details.
```

This README provides clear instructions on how to set up, run, and customize the URL shortener application.