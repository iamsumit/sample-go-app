FROM golang:1.20-alpine

WORKDIR /app

# Install task locally
RUN go install github.com/go-task/task/v3/cmd/task@latest

ENV GOCACHE /tmp/

CMD ["task", "sample:run-local"]
