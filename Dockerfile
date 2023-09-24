FROM golang:1.21.1

WORKDIR /app

COPY . .

RUN mkdir data && go mod download && go mod verify

RUN go build -o gourl

CMD [ "./gourl" ]