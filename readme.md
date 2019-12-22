git init -q
git remote add origin https://github.com/autocorrectoff/codixir

*creates a go.mod in the app root*
go mod init github.com/autocorrectoff/codixir

*downloads dependencies from import func*
go build -o main.go
