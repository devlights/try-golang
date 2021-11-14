FROM golang:1.17

ENV CGO_ENABLED=1
ENV GOOS=linux
ENV GOARCH=amd64

WORKDIR /go/src/app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

CMD ["make", "run"]

