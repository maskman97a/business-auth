FROM golang:latest

ENV WORKDIR /usr/src/app
WORKDIR "$WORKDIR"

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . "$WORKDIR"

RUN CGO_ENABLED=0 GOOS=linux cd ./cmd && go build -v -o /business-auth


EXPOSE 8010

CMD ["/business-auth"]
