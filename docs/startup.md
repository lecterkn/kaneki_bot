# 実行手順書

## 前提

- Go言語(`go1.24.1`)がインストールされていること

## 環境変数の設定

### 1. `.env`ファイルを配置

```sh
cp .env.example .env
```

### 2. GeminiのAPIキーを生成

[Google AI Studio](https://aistudio.google.com/apikey) からAPIキーを生成する

### 3. `.env`ファイルにGeminiのAPIキーを設定

```
GEMINI_API_KEY=<生成したAPIキー>
```

### 4. Discordのボットを作成

[Discord Developer Portal](https://discord.com/developers/applications) からボットを作成し、トークンを取得する

### 5. Discordのボットのトークンを設定

```
DISCORD_BOT_TOKEN=<ボットのトークン>
```

## 実行する

```sh
# 通常の起動
go run cmd/main.go

# airを使って起動する場合
air
```
