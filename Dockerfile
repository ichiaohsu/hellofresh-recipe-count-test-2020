FROM golang:1.16.5-buster
WORKDIR /build

RUN echo "appuser:x:1000:1000::/home/appuser:/bin/ash" > appuser

COPY . .

RUN GOOS=linux GOARCH=amd64 go build -o app ./main.go
COPY /appuser /etc/passwd
USER appuser

ENTRYPOINT ["/build/app"]
CMD []