package ascii

import (
	"errors"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

var (
	fileContents string
	fileName     string
)

func ErrorHandler(w http.ResponseWriter, errMsg string, statusCode int) {
	tmpl, err := template.ParseFiles("templates/error.html")
	if err != nil {
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		log.Printf("Error parsing error template: %v", err)
		return
	}

	data := struct {
		StatusCode int
		ErrMsg     string
	}{
		StatusCode: statusCode,
		ErrMsg:     errMsg,
	}

	w.WriteHeader(statusCode)
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		log.Printf("Error executing error template: %v", err)
	}
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		ErrorHandler(w, "Page Not Found", http.StatusNotFound)
		return
	}

	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		ErrorHandler(w, "Page Not Found", http.StatusNotFound)
		log.Printf("Error parsing template: %v", err)
		return
	}

	err2 := tmpl.Execute(w, nil)
	if err2 != nil {
		ErrorHandler(w, "Internal Server Error", http.StatusInternalServerError)
		log.Printf("Error executing template: %v", err2)
		return
	}
}

func SubmitHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		ErrorHandler(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	if r.URL.Path != "/ascii-art" {
		ErrorHandler(w, "Page Not Found", http.StatusNotFound)
		return
	}

	message := r.FormValue("message")
	bannerfile := r.FormValue("bannerfile")
	if message == "" || bannerfile == "" {
		ErrorHandler(w, "Bad Request: Missing message or banner file", http.StatusBadRequest)
		return
	}

	// Construct the file path
	filePath := "bannerfiles/" + bannerfile + ".txt"

	// Check if the banner file exists and has not been altered
	_, err := FileCheck(filePath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			ErrorHandler(w, "Internal Server Error: Banner file not found", http.StatusInternalServerError)
			log.Printf("Banner file not found: %s", filePath)
		} else if err.Error() == "the banner file has been altered" {
			ErrorHandler(w, "Internal Server Error: An unexpected error occurred. Please try again later.", http.StatusInternalServerError)
			log.Printf("Banner file altered: %s", filePath)
		} else {
			ErrorHandler(w, "Internal Server Error: An unexpected error occurred. Please try again later.", http.StatusInternalServerError)
			log.Printf("Error with banner file: %v", err)
		}
		return
	}

	data := strings.Split(message, "\r\n")
	var asciified strings.Builder
	for _, ch := range data {
		result, err := PrintBanner(ch, bannerfile)
		if err != nil {
			ErrorHandler(w, "Bad Request: Please use valid characters. Only printable characters from the ASCII table are allowed.", http.StatusBadRequest)
			log.Printf("Error printing banner: %v", err)
			return
		}
		asciified.WriteString(result)
	}

	isNotEmpty := false
	timeStamp := time.Now().Format("20060102150405")
	fileName = "file-" + timeStamp + ".txt"
	if len(asciified.String()) != 0 {
		isNotEmpty = true
		fileContents = asciified.String()
	}
	Data := struct {
		Ans      string
		Input    string
		NotEmpty bool
	}{
		Ans:      asciified.String(),
		Input:    message,
		NotEmpty: isNotEmpty,
	}
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		ErrorHandler(w, "Internal Server Error: An unexpected error occurred. Please try again later.", http.StatusInternalServerError)
		log.Printf("Error parsing template: %v", err)
		return
	}

	err = tmpl.Execute(w, Data)
	if err != nil {
		ErrorHandler(w, "Internal Server Error: An unexpected error occurred. Please try again later.", http.StatusInternalServerError)
		log.Printf("Error executing template: %v", err)
	}
}

func ExportHandler(w http.ResponseWriter, r *http.Request) {
	if fileContents == "" || fileName == "" {
		ErrorHandler(w, "No file to export", http.StatusBadRequest)
		return
	}

	// Create a temporary file to store the exported content
	tempFile, err := os.CreateTemp("", "tempfile")
	if err != nil {
		ErrorHandler(w, "Error creating temporary file", http.StatusInternalServerError)
		return
	}
	defer os.Remove(tempFile.Name())
	defer tempFile.Close()

	// Write the contents of the file to the temporary file
	_, err = io.WriteString(tempFile, fileContents)
	if err != nil {
		ErrorHandler(w, "Error writing to temporary file", http.StatusInternalServerError)
		log.Printf("Error writing temporary file: %v", err)
		return
	}

	// Set headers for file download
	w.Header().Set("Content-Disposition", "attachment;filename="+fileName)
	w.Header().Set("Content-Type", "text/plain")
	contentLength := len(fileContents)
	w.Header().Set("Content-Length", strconv.Itoa(contentLength))
	// Serve the temporary file
	http.ServeFile(w, r, tempFile.Name())
}
