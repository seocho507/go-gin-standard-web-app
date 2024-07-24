# For Local Development

FROM golang:1.22-alpine

RUN apk update && apk add --no-cache git gcc musl-dev

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=1 go build -o main .

CMD ["./main", "--env", "dev"]