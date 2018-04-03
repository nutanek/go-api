FROM golang
 
ADD . /go/src/github.com/nutanek/go-api
WORKDIR /go/src/github.com/nutanek/go-api

RUN go get github.com/gin-gonic/gin
RUN go install github.com/nutanek/go-api

ENTRYPOINT /go/bin/go-api
 
EXPOSE 8080
