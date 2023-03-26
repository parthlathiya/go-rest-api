FROM golang:1.14

WORKDIR /
COPY . .
EXPOSE 8092

RUN go get "github.com/go-sql-driver/mysql"
RUN go get "github.com/gorilla/mux"

CMD ["/server"]
