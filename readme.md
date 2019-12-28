*create the repo and clone*  

*creates a go.mod in the app root*  
go mod init github.com/autocorrectoff/codixir  

*downloads dependencies from import func*  
go build -o main.go  

*create postgres container*
docker run --name codixir-db -v codixir-db:/var/lib/postgresql/data -e POSTGRES_USER=codixir -e POSTGRES_PASSWORD=codixir -p 5432:5432 postgres:12.1-alpine

*create table*
create table books(id serial, title varchar, author varchar, year varchar);

*insert sample data*
insert into books(title, author, year) values ('Clean Code: A Handbook of Agile Software Craftsmanship', 'Robert C. Martin', '2008'),
('Spring Microservices in Action', 'John Carnell', '2017'),
('Go Programming Blueprints', 'Mat Ryer', '2016');

*note to future self: use gorm*
https://github.com/jinzhu/gorm

