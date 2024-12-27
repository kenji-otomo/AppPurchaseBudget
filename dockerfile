FROM golang:1.23.4-alpine

# Gitインストール
RUN apk update && \
    apk add --no-cache git

# build-arg で渡された環境変数を使って動作をカスタマイズ
ARG DNS
ARG VUE_URL

# 環境変数をコンテナ内に設定
ENV DNS=$DNS
ENV VUE_URL=$VUE_URL

# 作業ディレクトリを設定
WORKDIR /app

# Go Modulesの使用を許可
ENV GO111MODULE=on

# ローカルのモジュールキャッシュを最適化
COPY go.mod .
COPY go.sum .
RUN go mod download

# ソースコードをコンテナにコピー
COPY . .

# アプリケーションをビルド
RUN go build -o myapp

# 実行可能ファイルをデフォルトのコマンドとして設定
CMD ["./myapp"]