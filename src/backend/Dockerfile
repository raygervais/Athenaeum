FROM golang:1.14

# Set necessary environment variables needed for our image
RUN export GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Move to working directory /build
WORKDIR /build

# Copy and download dependency using go mod
COPY ./go.mod ./go.sum ./
RUN go mod download

# Copy the code into the container
COPY . .

EXPOSE 8000