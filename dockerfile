#get a base image
FROM golang:alpine

# Set necessary environmet variables needed for our image
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Move to working directory /build
WORKDIR /build

# Copy the code into the container
COPY . .

# Copy and download dependency using go mod
RUN go mod init chat-app-grpc
RUN go mod tidy
RUN go mod vendor
RUN mkdir -p /go/bin/logs
RUN chmod -R 777  /go/bin/logs

RUN go build -v -o goapp
CMD ["/build/goapp"]
