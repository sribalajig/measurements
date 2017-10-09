FROM golang:1.9

WORKDIR /go/src/github.com/hpi/measurement

COPY . .

RUN go get gopkg.in/mgo.v2
RUN go get -u github.com/gorilla/mux

ENTRYPOINT go run /go/src/github.com/hpi/measurement/pkg/cmd/api/main.go