#Builder Container
FROM golang:1.15 AS builder

ENV GOOS=linux
ENV GOARCH=amd64
ENV CGO_ENABLED=0
ENV GO111MODULE on

#作業ディレクトリを/appに指定
WORKDIR /app

#game-api/を/appにcopy
COPY . .

#airの追加
RUN go mod download && \
    go get -u github.com/cosmtrek/air

#airコマンド
CMD ["air"]


