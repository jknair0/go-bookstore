FROM golang:1.14.3-alpine
WORKDIR /src
COPY go.* .
RUN go mod download
COPY . .
RUN go build -o ./out/bookstore .
RUN chmod +x ./out/bookstore
CMD "./out/bookstore"