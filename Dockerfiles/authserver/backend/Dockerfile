FROM golang:1.20.4
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY *.go ./
COPY service/*.go ./service
RUN CGO_ENABLED=0 GOOS=linux go build -o /docker-gs-ping
EXPOSE 5052
CMD ["/docker-gs-ping"]

