### Movie DB

![example workflow](https://github.com/bajalnyt/movie-db/actions/workflows/go.yml/badge.svg)


––Following along the Go Course at https://lets-go-further.alexedwards.net/.
This project creates a JSON API for retrieving and managing information about movies, similar to OMDB. It uses PostgreSQL for the database.

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
### API endpoints

| Method	 | URL Pattern	              | Handler                                 | 	Action                      |
|---------|---------------------------|-----------------------------------------|------------------------------|
| GET	    | /v1/healthcheck	          | healthcheckHandler                      | Show application information |
| POST	   | /v1/movies	               | createMovieHandler	                     | Create a new movie           |
| GET     | /v1/movies                || Show the details of all movies          |
| POST    | /v1/movies                || Create a new movie                      |
| GET     | /v1/movies/:id            || Show the details of a specific movie    |
| PATCH   | /v1/movies/:id            || Update the details of a specific movie  |
| Delete  | /v1/movies/:id            || Delete a specific movie                 |
| POST    | /v1/users                 || Register a new user                     |
| PUT     | /v1/users/activated       || Activate a specific user                |
| PUT     | /v1/users/password        || Update the password for a specific user |
| POST    | /v1/tokens/authentication || Generate a new authentication token     |
| POST    | /v1/tokens/password-reset || Generate a new password-reset token     |
| GET     | /debug/vars               || Display application metrics             |
