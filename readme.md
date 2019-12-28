*create the repo and clone*  

*creates a go.mod in the app root*  
go mod init github.com/autocorrectoff/codixir  

*downloads dependencies from import func*  
go build -o main.go  

*create postgres container*
docker run --name codixir-db -v codixir-db:/var/lib/postgresql/data -e POSTGRES_USER=codixir -e POSTGRES_PASSWORD=codixir -p 5432:5432 postgres:12.1-alpine
