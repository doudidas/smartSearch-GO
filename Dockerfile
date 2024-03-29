#################################
# STEP 1 build executable binary
#################################

FROM golang:latest AS builder


# Git is required for fetching the dependencies.
# RUN apk add --no-cache git

# Set the working directory outside $GOPATH to enable the support for modules.
WORKDIR /src

# Fetch dependencies first; they are less susceptible to change on every build
# and will therefore be cached for speeding up the next build
COPY ./go.mod ./
RUN go mod tidy
# Import the code from the context.
COPY ./ ./

# Build the executable to `/app`. Mark the build as statically linked.
 
RUN CGO_ENABLED=0 go build -o build/smartsearch .
# RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o build/smartsearch .
#################################
# STEP 2 Copy into a small image
#################################

FROM scratch

# Import the compiled executable from the first stage.
COPY --from=builder /src/build /src/bin
COPY default-credantials.json /etc/smartsearch/apiAdmins.json

#Set env variable
ENV GIN_MODE=release
ENV MONGO_HOSTNAME=smartsearch-db
ENV MONGO_PORT=27017
# Expose port
EXPOSE 9000
VOLUME [ "/etc/smartsearch/" ]
# Run the hello binary.
CMD ["/src/bin/smartsearch"]
