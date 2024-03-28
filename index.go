package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Define header
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, `
		<html>
			<head>
				<title>Nick's Site</title>
			</head>
			<body>
				<h1>Hope you're hungry!</h1>
				<p>Come on down, Yinzers.</p>
				<img src="https://upload.wikimedia.org/wikipedia/commons/thumb/1/15/Pittsburgh_Strip_District_Primanti_Bros.jpg/1200px-Pittsburgh_Strip_District_Primanti_Bros.jpg" height ="500" width="500">
			</body>
		</html>
		`)
	})

	// Sports
	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, `
		<html>
			<head>
				<title>About this page</title>
			</head>
			<body>
				<h1>This was an experimenal page</h1>
				<p>Made in Golang</p>
				<img src="https://miro.medium.com/v2/resize:fit:2000/format:webp/1*8bPiDNL1K1ZdK9O_T5IVKw.png" height ="500" width="500">
			</body>
		</html>
		`)
	})
	//Hobbies
	http.HandleFunc("/hobbies", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, `
		<html>
			<head>
				<title>Hobbies</title>
			</head>
			<body>
				<h1>Outside of making websites, I like painting miniatures</h1>
				<p>Check them out.</p>
				<img src="https://i.imgur.com/W0DwhPL.jpeg" height ="500" width="500">
			</body>
		</html>
		`)
	})
	// Start the service
	fmt.Println("Server listening on port 8080...")
	http.ListenAndServe(":8080", nil)
}
