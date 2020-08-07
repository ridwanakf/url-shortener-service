# URL Shortener Service
Service for shortening URL like bit.ly, ugm.id, etc. written in Go.

This project is considered finish (all basic functionality of url shortener had been implemented), However there are some enhancements like Web UI, User login, etc. that is nice to be implemented as well later. Please take a look at [TODO.md](./TODO.md) document for details.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing
purposes.

### Prerequisites

1. Clone repository: `git clone git@github.com:ridwanakf/url-shortener-service.git`
2. Install Postgresql: [Postgresql](https://www.postgresql.org/download/)
3. Run database migrations: refer below
4. Install Redis: [Redis](https://redis.io/download)

### Run Project

Running project:

```$xslt
make run
```

### Migrations

When running in Local, you need to run the db-migrations to setup the app's database for your local machine.

1. Go to directory `url-shortener-service/migrations`
2. Run `go run *.go up`

## Directory Structure

This repository is organized in the following directory structure.

```
url-shortener-service
|-- bin                                    # Contains binary of the built app
|-- cmd                                    # Contains executable codes and serves as entry point of the app
|   |-- main                               # entry point of the app
|-- config                                 # Configuration files needed for deployment
|-- constant                               # Collections of constants file for each module
|-- internal                               # Go files in this folder represent the Big-Pictures and Contracts of the system
|   |-- app                                # Contains dependency injection of the app and other app's related configs
|   |   |-- config                         # Configuration struct for the app
|   |-- delivery                           # Delivery layer of the app
|   |   |-- rest                           # HTTP REST API delivery of the app
|   |   |-- <other_delivery_mechanisms>    # Other delivery mechanisms of the app (eg. GRPC, Console, Web, etc.)
|   |-- entity                             # Enterprise Data structures
|   |   |-- url.go                         # Data structure for URL
|   |   |-- <other_entities>.go            # Other data structures, preferrably 1 struct 1 file 
|   |-- repo                               # Implementations of Repository-pattern to data-sources
|   |   |-- db                             # Implementations of the repositories with Postgres database
|   |   |-- <other_repos>                  # Other Repositories implementations based on interfaces on folder internal.
|   |-- usecase                            # Usecases implementations for Application Business Logic
|   |   |-- url_shortener_usecase.go       # Other use-case implementations based on interfaces on folder internal.
|   |-- repo.go                            # Interfaces / Contracts of all the repositories (Repository Pattern)
|   |-- usecase.go                         # Interfaces / Contracts of all the use-cases (Application Business Logic)
|
|-- migrations                             # Contains Database migration files or the system
|-- web                                    # UI related code (React app)

```
## Contributing

Your contributions are welcome!

Please take a look at [CONTRIBUTING.md](./CONTRIBUTING.md) document for details.

## Tech Stacks

- Golang
- Postgresql
- Redis

