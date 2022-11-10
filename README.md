# Bachman-CRUD

#### Fast CRUD API
<img src="https://img.shields.io/badge/Build-SUCCESS-green"/>          

## Who is Bachman ?
Charles William Bachman III (December 11, 1924 – July 13, 2017) was an American computer scientist, who spent his entire career as an industrial researcher, developer, and manager rather than in academia. He was particularly known for his work in the early development of database management systems. His techniques of layered architecture include his namesake Bachman diagrams. 
 - [Full Article](https://en.wikipedia.org/wiki/Charles_Bachman)

## Requeriments for development
- Go : [Install](https://go.dev/dl/)
- Any code editor.

## Content
- main.go: This file runs the server with the CRUD api.
- The API: The API is about movies and save the informations in the local memory. 
- This itens can be edited a visualized by this paths:
  1. "/movies" - Method "GET" = get Movies
	2. "/movies/{id} - "Method "GET" = get Movie
	3. "/movies" - Methods "POST" = create Movie
	4. "/movies/{id}" - Methods "POST" = update Movie
	5. "/movies/{id}" - Method "DELETE" = delete Movie
  
 *The parm {id} is the Movie ID*


## Start Bachman
- For the first step go to the main.go's folder using your terminal/cmd.
- After run in the terminal/cmd this command:
```sh
go build
```
- Finally run this command for start the server for the CRUD:
```sh
go run main.go
```
- The server will show if is running.
- for test you can use [Insominia](https://insomnia.rest/download) or [Postman](https://www.postman.com/).
- After you can make the request.
