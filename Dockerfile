FROM golang:latest

WORKDIR /ss/bomber-man

COPY go.mod ./

RUN go mod download

COPY . .

RUN go build -o main .

CMD [ "./main" ]