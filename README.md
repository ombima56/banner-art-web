# banner-art-web

banner-art-web is a web application that allows users to convert text to ASCII art using various banner styles. Users can view the generated ASCII art on the web interface and download it as a text file.

## Table of Contents

- [Installation](#installation)
- [Usage](#usage)
- [Endpoints](#endpoints)
- [Contributing](#contributing)
- [License](#license)

## Installation

1. **Clone the repository**:
```sh
git clone https://github.com/ombima56/banner-art-web.git
cd banner-art-web
```

## Install dependencies:
Make sure you have Go installed.

```go
go mod tidy
```
Run the application:

```go
go run main.go
```
Optionally, you can specify a port number:  `between 1024 - 65535`
## Example
```go
go run main.go 60000
```
## Usage

- Open your web browser and navigate to http://localhost:8080 or http://localhost:`[port number]`.
- Enter your text and select a banner style (`standard.txt`, `shadow.txt`, or `thinkertoy.txt`).
- Click on "Generate" to view the ASCII art.
- Click on "Download" to save the generated ASCII art as a text file.

## Endpoints
```sh
1. GET /
Renders the main page where users can input text and select a banner style.
2. POST /ascii-art
Generates ASCII art from the input text using the selected banner style.
3. GET /export
Exports the generated ASCII art as a text file.
4. GET /health
A health check endpoint that returns 200 OK.
```
## Contribution

- Fork the repository.
- Create your feature branch (git checkout -b feature/awesome-feature).
- Commit your changes (git commit -m 'feat: add awesome feature').
- Push to the branch (git push origin feature/awesome-feature).
- Open a pull request.

## License

This project is licensed under the MIT License. See the [LICENSE](./LICENSE) file for details.
