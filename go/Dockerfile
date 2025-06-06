# ベースイメージ
FROM golang:1.23.3

# 作業ディレクトリの作成
WORKDIR /app

# go.mod と go.sum をコピー（依存関係を先にインストールするため）
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# ソースコードをコピー
COPY . .

# アプリケーションをビルド
RUN go build -o main .

# ポートの指定（必要であれば）
EXPOSE 8090

# 起動コマンド
CMD ["./main"]
