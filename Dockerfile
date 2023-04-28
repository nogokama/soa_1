FROM golang:latest
WORKDIR /go/src/app
COPY . /go/src/app
# RUN  apt update &&  apt install -y netcat
RUN go build -o app ./server
ENTRYPOINT ["./app"]
