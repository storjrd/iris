FROM golang:buster

RUN apt-get update
RUN apt-get upgrade -y

WORKDIR /tmp

RUN wget https://dl.google.com/linux/direct/google-chrome-stable_current_amd64.deb
RUN apt install ./google-chrome-stable_current_amd64.deb -y

RUN whereis google-chrome

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY *.go .

RUN go build -o /ui-tests

CMD [ "/ui-tests" ]
