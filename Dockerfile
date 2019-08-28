#################################
# STEP 1 build executable binary
#################################

FROM golang:alpine AS builder


# Git is required for fetching the dependencies.
RUN apk add --no-cache git

# Set the working directory outside $GOPATH to enable the support for modules.
WORKDIR /src

# Fetch dependencies first; they are less susceptible to change on every build
# and will therefore be cached for speeding up the next build
COPY ./go.mod ./go.sum ./
RUN go mod download
# Import the code from the context.
COPY ./ ./

# Build the executable to `/app`. Mark the build as statically linked.
RUN CGO_ENABLED=0 go build -installsuffix 'static' -o /app .

#################################
# STEP 2 Copy into a small image
#################################

FROM scratch

# Import the compiled executable from the first stage.
COPY --from=builder /app /app

# Add Favicon
COPY favicon.ico favicon.ico

#Set env variable
ENV GIN_MODE=release
ENV MONGO_HOSTNAME=smartsearch-db
ENV MONGO_PORT=27017
# Expose port
EXPOSE 9000

# Run the hello binary.
ENTRYPOINT ["/app"]