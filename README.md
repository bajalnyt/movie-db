### Movie DB

![example workflow](https://github.com/bajalnyt/movie-db/actions/workflows/go.yml/badge.svg)


Go Course at https://lets-go-further.alexedwards.net/

#### Project layout

```
├── Makefile # Runs tasks
├── README.md
├── bin     # Compiled binaries
├── cmd
│   └── api # app specific code, the server, authentication
│       └── main.go
├── go.mod # project deps, versions, module path
├── internal    # Various ancillary packages for db/validation/email etc. Called by code in `api`, but never the other way
├── migrations # SQL migrations
└── remote # config files
```
### API

|Method	|URL Pattern	|Handler|	Action|
| ----------- | ----------- | ----------- | ----------- |
|GET	|/v1/healthcheck	|healthcheckHandler|	Show application information
|POST	|/v1/movies	|createMovieHandler	|Create a new movie
|GET	|/v1/movies/:id	|showMovieHandler|	Show the details of a specific movie”
