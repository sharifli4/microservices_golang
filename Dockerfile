FROM golang:latest
COPY . .
WORKDIR /home/app
CMD ["go","run","main.go"]