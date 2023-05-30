FROM golang:1.17

WORKDIR /src

COPY go.mod go.sum /src/

RUN go mod download

COPY . /src/

CMD [ "sh", "-c", "make test" ]
