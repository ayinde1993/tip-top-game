# TIP-TOP-GAME

# Getting started

Hello, nice having you aboard🤝🕺🏼

## Contributing guidelines

Review the contributing guidelines available [Contributing Guidelines](CONTRIBUTION.md), and kindly follow the processes
in your development workflow.

## New to Golang?

- [Go Intro](https://tour.golang.org)

- [Learning Go](https://medium.com/go-go-go/learning-go-golang-47127a796323)

- [8 Insights](https://medium.com/go-go-go/golang-8-insights-of-the-first-weeks-of-the-real-usage-f01290811b8b)

- [How to Write Go Code](https://golang.org/doc/code.html)

## ORM

- [SQLBoiler] (https://github.com/volatiletech/sqlboiler)

### Installing SQLBoiler
- Go version 1.16 and above

```
go install github.com/volatiletech/sqlboiler/v4@v4.15
go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-psql@v4.16.2
```

- Go version 1.15 and below:
- Install sqlboiler v4 and mysql  driver (mysql, mssql, sqlite3, postgresql also available)
- NOTE: DO NOT run this inside another Go module (like your project) as it will
- pollute your go.mod with a bunch of stuff you don't want and your binary

```
GO111MODULE=on go get -u -t github.com/volatiletech/sqlboiler/v4
GO111MODULE=on go get github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-psql
go get github.com/volatiletech/null/v8
```

### Generate SQLBoiler code

Ensure you have docker running as the command utilizes throw-away docker containers to ensure consistency.

```
$ make sqlgenerate
```

## Setting-up development environment

### Step 1. Install Golang

[Golang](https://golang.org/doc/install)

Make sure your Go version is at least 1.19.

### Step 2. Install Docker

[Install Docker](https://orbstack.dev)

### Step 5. Put sources in the correct folder

```
$ git clone https://github.com/LiftMyBag/SpaceHolder.git

$ cd SpaceHolder

$ go mod download
```

### Checking that everything works

#### 1. Build and run Database in docker

Create a file called `config.env` and add the following data to it:
```
POSTGRES_DB=space_holder_test
POSTGRES_USER=spaceholder-user
POSTGRES_PASSWORD=password123
```
Then run
```
$ make mysql-db
$ docker ps (you should see the postgres image running)
```

Alternatively, you can setup Postgres on your local machine then login to the Postgres DB using the default user `postgres` and create the user with password in the database
console using query below :

`CREATE USER spaceholder-user WITH PASSWORD 'password123';`

#### 2. Install golang migrate for database migration

[Golang Migrate](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate)

#### 3. Run migration file
```
$ brew install golang-migrate (for Mac users, for windows users please check there site for commands to run)
$ make migrateup
```

#### 4. Setup Redis

```
$ make tip-top-game-redis
$ docker ps (you should see the redis image running)
```

#### 5. Start the app

##### For local run

```
$ make run
```

##### For docker run

NB: Change in .env file 'SPACE_HOLDER_DB_HOST=host.docker.internal'

```
$ make build-docker
$ make run-docker
```

### IDEs, Editors

- Atom
- Vim
- Emacs
- GoLand
- Visual Studio Code

### Libraries
- ([For logging](https://github.com/rs/zerolog))
- ([For http router](https://github.com/gin-gonic/gin))


### Note:

Please don't make changes on develop and push directly to GitHub. Also, don't merge your PR to master until a review has
been done on it, and you receive a green light to go ahead.
