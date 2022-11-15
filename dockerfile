FROM golang:1.16-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY *.go ./

COPY static/*.* ./static/

RUN go build -o /mytodo

EXPOSE 3000

CMD [ "/mytodo" ]