FROM golang:1.23.5

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /deathbox

EXPOSE 8080

CMD ["/deathbox"]