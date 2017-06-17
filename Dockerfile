FROM golang:1.8

#RUN echo $GOPATH => /go

#ARG redisHostName=default-redis-server
#ARG mysqlHostName=default-mysql-server

RUN apt-get -y update && apt-get install -y git


RUN go get -u github.com/goadesign/goa/... && \
#go get -u github.com/pilu/fresh
go get -u github.com/hiromaily/fresh

RUN go get -u github.com/hiromaily/go-goa/...
#RUN go get -d -v ./ext/cmd/
#RUN go get -d -v

#submodule for swaggger-ui
WORKDIR /go/src/github.com/hiromaily/go-goa/resources/swagger-ui
RUN git submodule init && git submodule update


RUN mkdir -p /go/src/github.com/hiromaily/go-goa/tmp/log
#RUN mkdir -p /go/src/github.com/hiromaily/go-goa/ext && \
# mkdir -p /go/src/github.com/hiromaily/go-goa/goa && \
# mkdir -p /go/src/github.com/hiromaily/go-goa/public && \
# mkdir -p /go/src/github.com/hiromaily/go-goa/resources

WORKDIR /go/src/github.com/hiromaily/go-goa
# go get can retrieve all data from git repo
#COPY ./ext ./
#COPY ./goa ./
#COPY ./public ./
#COPY ./resources ./
#COPY ./Makefile ./
#COPY ./runner.conf ./

#
RUN ln -s /go/src/github.com/hiromaily/go-goa/goa/swagger /go/src/github.com/hiromaily/go-goa/public/swagger

RUN CGO_ENABLED=0 GOOS=linux go build -o /go/bin/go-goa ./ext/cmd/main.go

#EXPOSE 80
CMD ["/go/bin/go-goa"]