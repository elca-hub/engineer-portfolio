# Engineer Portfolio Backend

## プロジェクト概要

このプロジェクトは、エンジニアのポートフォリオ管理システムのバックエンド部分です。Go言語を使用し、クリーンアーキテクチャの原則に従って実装されています。

## 技術スタック

- 言語: Go 1.24.0
- フレームワーク: Gin
- データベース:
  - MySQL (SQLBoiler)
  - Redis
- ファイルストレージ: MinIO (S3互換)
- ロギング: Zap
- バリデーション: Go Playground Validator
- スケジューリング: Cron
- メール送信: Gomail

## プロジェクト構造

```txt
.
├── adapter/        # 外部インターフェースのアダプター
├── domain/         # ドメインモデルとビジネスロジック
├── infra/          # インフラストラクチャ層
│   ├── cron/       # スケジューリング関連
│   ├── database/   # データベース接続
│   ├── file_uploader/ # ファイルアップロード
│   ├── log/        # ロギング
│   ├── router/     # HTTPルーティング
│   └── validation/ # バリデーション
├── migrations/     # データベースマイグレーション
├── presenter/      # プレゼンテーション層
├── usecase/        # ユースケース層
├── main.go         # アプリケーションのエントリーポイント
└── Dockerfile      # コンテナ化設定
```

## 主要な機能

1. HTTPサーバー
   - RESTful API
   - CORS対応
   - セッション管理
   - ファイルアップロード

2. バッチ処理
   - 10分間隔での定期実行
   - バックグラウンドタスク

3. データベース操作
   - MySQLによる永続化
   - Redisによるキャッシュ

4. ファイル管理
   - S3互換ストレージ（MinIO）によるファイル保存

## 開発環境

- ホットリロード: Air
- テスト: Testify
- モック: Uber Mock

## セキュリティ機能

- セッション管理
- バリデーション
- セキュアなファイルアップロード

## デプロイメント

- Dockerコンテナ化対応
- 環境変数による設定管理
- マイグレーション管理（Atlas）

## 依存関係

主要な外部パッケージ:

- gin-gonic/gin: Webフレームワーク
- go-redis/redis: Redisクライアント
- sqlboiler: ORM
- zap: ロギング
- gomail: メール送信
- aws-sdk-go-v2: S3互換ストレージ操作
