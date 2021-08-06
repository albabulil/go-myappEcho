module github.com/albabulil/go-myappEcho

go 1.16

replace github.com/albabulil/go-myappEcho => ./

require (
	github.com/go-playground/validator/v10 v10.8.0
	github.com/joho/godotenv v1.3.0
	github.com/labstack/echo/v4 v4.5.0
	github.com/mattn/go-isatty v0.0.13 // indirect
	go.mongodb.org/mongo-driver v1.7.1
	golang.org/x/crypto v0.0.0-20210711020723-a769d52b0f97
	golang.org/x/net v0.0.0-20210726213435-c6fcb2dbf985 // indirect
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1 // indirect
)
