FROM golang:1.23.1

WORKDIR /app

RUN go install github.com/a-h/templ/cmd/templ@latest
RUN go install github.com/air-verse/air@latest

COPY . .

RUN go mod download && go mod verify

# Install Node.js and npm
RUN apt-get update && apt-get install -y curl
RUN curl -sL https://deb.nodesource.com/setup_18.x | bash -
RUN apt-get install -y nodejs

# Verify that Node.js and npm were installed
RUN node -v
RUN npm -v

# Install deps
RUN npm install

CMD ["make", "build"]