# Vue-api

This is the repository for Vue api of Udemy-Working-with-Vue-3-and-Go-Golang-Vue.

### Docker setup
$ docker-compose build

$ docker-compose up -d

### version
- Build in Go version 1.17
- Install [make](https://www.gnu.org/software/make/manual/) globally via [Homebrew](https://brew.sh/) in Mac

### modules
- Uses [chi router](https://github.com/go-chi/chi)
- Uses [chi cors](https://github.com/go-chi/cors)
- Uses [jackc/pgx](https://github.com/jackc/pgx)
-> $ go get github.com/jackc/pgconn
-> $ go get github.com/jackc/pgx/v4
-> $ go get github.com/jackc/pgx/v4/stdlib

- Uses [mozillazg / go-slugify](https://github.com/mozillazg/go-slugify)
- Uses [DATA-DOG / go-sqlmock](https://github.com/DATA-DOG/go-sqlmock)
-> $ go get github.com/DATA-DOG/go-sqlmock
-> $ go test -v .
-> $ go test -coverprofile=coverage.out && go tool cover -html=coverage.out

### references
- [PostgreSQL Error Codes](https://www.postgresql.org/docs/14/errcodes-appendix.html)