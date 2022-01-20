#### Movie DB

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
