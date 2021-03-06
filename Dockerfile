FROM golang:1.16

ENV CGO_ENABLED=1
ENV GOOS=linux
ENV GOARCH=amd64

WORKDIR /go/src/app

COPY go.mod go.sum ./
RUN go mod download
RUN go install github.com/cheekybits/genny

COPY . .

CMD ["make", "run"]

