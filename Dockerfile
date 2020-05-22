FROM golang:1.14

WORKDIR /
COPY . .

RUN go get "github.com/go-sql-driver/mysql"
RUN go get "github.com/gorilla/mux"
