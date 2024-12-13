FROM golang:1.23.4-alpine

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