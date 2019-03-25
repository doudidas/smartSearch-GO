############################
# STEP 1 build executable binary
############################
FROM golang:alpine AS builder
# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git
RUN mkdir /app 
ADD . /app/ 
WORKDIR /app 
# Fetch dependencies.
# Using go get.
RUN go get -d -v github.com/gin-gonic/gin github.com/lib/pq github.com/mongodb/mongo-go-driver/bson github.com/mongodb/mongo-go-driver/mongo
# Build the binary.
RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /go/bin/main
############################
# STEP 2 build a small image
############################
FROM scratch
# Copy our static executable.
COPY --from=builder /go/bin/main /go/bin/main
#Set env variable
ENV GIN_MODE=release
# Expose port
EXPOSE 9000
# Run the hello binary.
ENTRYPOINT ["/go/bin/main"]