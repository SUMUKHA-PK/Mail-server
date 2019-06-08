FROM docker.io/golang:1.12.5-alpine3.9

LABEL  maintainer = "Sumukha PK"

WORKDIR github.com/SUMUKHA-PK/Mail-Server

COPY . .

RUN apk add git

RUN go build -o main -v ./cmd/... 

EXPOSE 8080

ENTRYPOINT ["./main"]

CMD ["go-dock1"]