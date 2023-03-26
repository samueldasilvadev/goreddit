# goreddit

Simple project based on course: https://course.gowebexamples.com/course.
Project structure:

```
.
├── cmd
│   └── goreddit
│       └── main.go
├── .env.example
├── go.mod
├── goreddit.go
├── go.sum
├── Makefile
├── migrations
│   ├── 1_create_tables.down.sql
│   └── 1_create_tables.up.sql
├── postgres
│   ├── comment_store.go
│   ├── post_store.go
│   ├── store.go
│   └── thread_store.go
├── read.go
├── README.md
└── web
    ├── handler.go
    └── templates
        ├── ThreadsCreate.html
        └── ThreadsList.html
```

Obs.: this project does not implements an architecture pattern, is only test.

## Steps

1. go initial config

```
go mod init samueldasilvadev.com/goreddit
go get github.com/google/uuid
go get github.com/gorilla/mux

go mod download
```

2. Define entities:
- Thread
- Post
- Comment

```
+----------------+              +-------------------+
|                |              |                   |
|                |1           n |                   |
|      Threads   ----------------       Posts       |
|                |              |                   |
|                |              |                   |
+----------------+              +-------------------+
                                        / 1         
                                       /            
                                      /             
                +------------------+ /               
                |                  |/                
                |                  / n                
                |     Comments     |                 
                |                  |                 
                |                  |                 
                +------------------+                 
```

3. Create Structs and Interfaces in goreddit.go

4. Create migrations up and down in ./migrations

5. Install go-migrate: https://github.com/golang-migrate/migrate

6. Create Makefile and run migration (https://course.gowebexamples.com/course/creating-a-reddit-clone-in-go/installing-postgresql-using-docker)

  make postgres
  make adminer
  make migrate

7. Database queries using sqlx (https://course.gowebexamples.com/course/creating-a-reddit-clone-in-go/database-queries-with-sqlx)

8. Serving HTML with chi router.
