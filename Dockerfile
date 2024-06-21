FROM golang:1.22.4

RUN mkdir -p /go/src/app
WORKDIR /go/src/app
COPY . .


RUN go install github.com/air-verse/air@latest
RUN go mod tidy


VOLUME [ "/go/src/app" ]

EXPOSE 8080

CMD ["air"]