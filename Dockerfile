FROM golang:1.14

WORKDIR /
COPY . .
EXPOSE 8080

RUN go get "github.com/go-sql-driver/mysql"
RUN go get "github.com/gorilla/mux"
