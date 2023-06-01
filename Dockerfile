FROM golang:latest

RUN go version
ENV GOPATH=/

COPY ./ ./

# install psql
RUN apt-get update
RUN apt-get -y install postgresql-client


# build go app
RUN go mod download
RUN go build -o smarfit-app ./cmd/main.go

CMD ["./smarfit-app"]