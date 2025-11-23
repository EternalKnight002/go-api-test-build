# Go Playground Server

A small Go web server designed to demonstrate a set of common web development features: Go HTML template rendering, serving static assets, and implementing several simple and fun JSON APIs.

It serves a single-page application at the root path (`/`) which utilizes all the JSON APIs.

## Features

* **Go HTML Templates**: Renders the main `index.html` page using Go's `html/template` package.
* **Static Assets**: Serves CSS and JavaScript files from the `/static/` path.
* **Greeting API**: A simple JSON API available at `/greet`.
* **Double API**: A JSON API available at `/double` for integer multiplication.
* **Subtraction API**: A JSON API available at `/subtract` for basic arithmetic subtraction.
* **Fun Pet Name API**: A unique JSON API at `/petname` for generating a pet name using string manipulation.

---

### Prerequisites

* Go (The project specifies `go 1.25.0`).

### 1. Run the server

You can run the server directly using the `go run` command:

```bash
go run ./cmd/server
```

The server will start at `http://localhost:8080`.

### 2. Build the executable

You can compile the server into an executable binary:

```bash
go build -o server ./cmd/server
# Then run:
# ./server
```

## API Endpoints

The server exposes four JSON endpoints:

### 1. Greeting API (`/greet`)

* **URL**: `/greet` or `/greet?name=<name>`
* **Method**: `GET`
* **Description**: Returns a simple greeting message. If no `name` parameter is provided, it defaults to `"Lord"`.

**Example Request:**
```
/greet?name=Alice
```

**Example Response (JSON):**
```json
{
    "message": "Hello, Alice!"
}
```

### 2. Double API (`/double`)

* **URL**: `/double?number=<integer>`
* **Method**: `GET`
* **Description**: Calculates and returns the double of the provided integer `number`. If no valid number is provided, it defaults to `0`.

**Example Request:**
```
/double?number=21
```

**Example Response (JSON):**
```json
{
    "input": 21,
    "result": 42
}
```

### 3. Subtraction API (`/subtract`)

* **URL**: `/subtract?n1=<integer>&n2=<integer>`
* **Method**: `GET`
* **Description**: Calculates the difference: n1 minus n2. Defaults to 0 for missing or invalid inputs.

**Example Request:**
```
/subtract?n1=100&n2=50
```

**Example Response (JSON):**
```json
{
    "n1": 100,
    "n2": 50,
    "result": 50
}
```

### 4. Fun Pet Name Generator API (`/petname`)

* **URL**: `/petname?adjective=<string>&animal=<string>`
* **Method**: `GET`
* **Description**: Generates a unique pet name by reversing the input `adjective`, capitalizing the first letter, and then appending the capitalized `animal`. Defaults to `adjective=mysterious` and `animal=shrimp` if parameters are missing.

**Example Logic/Output** (for `sleepy` and `cat`):
1. Reverse `sleepy` to `ypeels`.
2. Capitalize first letter: `Ypeels`.
3. Combine with capitalized `Cat` → `YpeelsCat`.

**Example Request:**
```
/petname?adjective=sleepy&animal=cat
```

**Example Response (JSON):**
```json
{
    "adjective": "sleepy",
    "animal": "cat",
    "petName": "YpeelsCat"
}
```

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.