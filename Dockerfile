############################
# STEP 1 build executable binary
############################
FROM golang:alpine AS builder

# Setup work directory
WORKDIR /go/src/app
# Add my code on container
ADD . /go/src/app
# Use script to avoid multple docker RUN 
RUN /go/src/app/docker.sh

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