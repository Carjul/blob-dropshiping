

FROM golang:alpine3.18
# Copia los archivos necesarios para la aplicación
RUN mkdir -p /go/src/app
WORKDIR /go/src/app
COPY . .


RUN go install github.com/cosmtrek/air@latest
RUN go mod tidy


VOLUME [ "/go/src/app" ]

EXPOSE 3000

CMD ["cd cmd","air"]