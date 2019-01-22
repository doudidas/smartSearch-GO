FROM golang:latest as build
RUN mkdir /app 
ADD . /app/ 
WORKDIR /app 
RUN go get -v github.com/gin-gonic/gin && go get -v gopkg.in/mgo.v2 && CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM scratch
COPY --from=build app/main ./app
EXPOSE 9000
CMD ["./app", "mongo"]