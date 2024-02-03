FROM golang:1.21.1

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .

# Build
RUN go build -o /task


CMD ["/task"]