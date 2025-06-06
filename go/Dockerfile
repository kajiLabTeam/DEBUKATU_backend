# ベースイメージ
FROM golang:1.23.3

# 作業ディレクトリの作成
WORKDIR /app

# go.mod と go.sum をコピー（依存関係を先にインストールするため）
# COPY go.mod ./app/
# COPY go.sum ./app/
# RUN go mod download

# ソースコードをコピー
COPY . .

# ポートの指定（必要であれば）
EXPOSE 8080

# 起動コマンド
CMD ["go", "run", "main.go"]
