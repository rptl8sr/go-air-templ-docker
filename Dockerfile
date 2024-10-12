FROM golang:1.23.1

WORKDIR /app

RUN go install github.com/a-h/templ/cmd/templ@latest
RUN go install github.com/air-verse/air@latest

COPY . .

RUN go mod download && go mod verify

CMD ["make", "build"]