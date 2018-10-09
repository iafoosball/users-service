# STEP 1 build executable binary

FROM golang:1.10 as builder

# Download and install the latest release of dep
ADD https://github.com/golang/dep/releases/download/v0.4.1/dep-linux-amd64 /usr/bin/dep
RUN chmod +x /usr/bin/dep

#Download the service
RUN mkdir -p /go/src/github.com/iafoosball
RUN go get -u github.com/iafoosball/users-service
#WORKDIR /go/src/github.com/iafoosball
#RUN git clone https://github.com/iafoosball/users-service.git
WORKDIR /go/src/github.com/iafoosball/users-service
RUN ls

#Download and install swagger in go and run codegen
RUN go get -u github.com/go-swagger/go-swagger/cmd/swagger
RUN /go/bin/swagger generate server -f /go/src/github.com/iafoosball/users-service/users.yml -A users
RUN go get -u golang.org/x/net/netutil
RUN dep ensure -vendor-only

#Install the service
WORKDIR /go/src/github.com/iafoosball/users-service/cmd/users-server/
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags "-s -w" -a -installsuffix cgo -o users-service .


# STEP 2 build a small image
# start from scratch
#FROM scratch

FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy our static executable
COPY --from=builder /go/src/github.com/iafoosball/users-service/cmd/users-server/users-service .
CMD ["./users-service","--port","4444","--host","0.0.0.0"]