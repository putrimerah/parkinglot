FROM golang:1.21

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go install github.com/mattn/go-sqlite3

CMD ["go", "run", "./cmd"]