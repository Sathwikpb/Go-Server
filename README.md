# Go-Server
Package main provides a simple HTTP server with two endpoints and static file serving capabilities.

Overview:
- The server listens on port 8080.
- It serves static files from the "./static" directory.
- The "/hello" endpoint responds to HTTP GET requests with a "hello" message.
- The "/form" endpoint handles HTTP POST requests, parses form data, and displays the submitted data.

Endpoints:
1. /hello
   - Method: GET
   - Response: "hello" message
   - Errors: 404 Not Found for non-existent paths, 405 Method Not Allowed for non-GET requests.

2. /form
   - Method: POST
   - Response: "POST req successful" message followed by the submitted form data (name and address).
   - Errors: 404 Not Found for non-existent paths, and 405 Method Not Allowed for non-POST requests.
   
Static File Serving:
- The server serves static files (e.g., HTML, CSS, JS) from the "./static" directory.

Usage:
- Start the server by running the main function.
- Access the "/hello" endpoint via a web browser or HTTP client to receive the "hello" message.
- Access the "/form" endpoint with a POST request containing form data to see the submitted data.

Files:
1. Static file
   - index.html
   - form.html
   - style.css
   - script.js
2. main.go
