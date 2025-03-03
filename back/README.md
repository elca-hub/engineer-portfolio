# DevPort | backend

## マイグレーションについて

今回はAtlasとgormを使用してマイグレーションやスキーマ管理を行なっています。
コマンド関係はcmdフォルダにまとめてあります。
実行する際は**back**ディレクトリで実行してください。

### マイグレーションの実行

```bash
./cmd/migrate-diff
```

### マイグレーションの適用

```bash
./cmd/migrate-apply
```