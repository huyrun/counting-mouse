FROM golang:1.17

WORKDIR /src

COPY go.mod go.sum /src/

RUN go mod download

COPY . /src/

RUN CGO_ENABLED=0 go build -a -tags musl -o app cmd/main.go
