# workspace (GOPATH) configured at /go
FROM golang:1.20 as builder

#
WORKDIR /app

# Copy the local package files to the container's workspace.
COPY . ./

# installing depends and build
RUN go build -o ./go_to_perfomance main.go

FROM alpine
WORKDIR /app
COPY --from=builder . .

ENTRYPOINT ["/go_to_perfomance"]
