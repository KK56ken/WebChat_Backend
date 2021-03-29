# 2020/10/14最新versionを取得
FROM golang:1.16.2-alpine
# アップデートとgitのインストール！！
RUN apk update && apk add git
# appディレクトリの作成
RUN mkdir /go/src/app

# ワーキングディレクトリの設定
WORKDIR /go/src/app

CMD ["go","run","main.go"]