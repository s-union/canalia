# Quickstart: /user API

この手順はローカル開発環境で /user の契約と最小動作を確認するためのものです。

## 前提
- Docker が起動済み
- apps/go-api/.env.local に Auth0 と DB の設定済み

## 手順（概要）
1. DB起動とマイグレーション適用（Taskfile）
2. OpenAPI型生成（oapi-codegen）
3. サーバ起動
4. 認証トークンを付けて GET /user を確認
5. 認証トークンを付けて POST /user を実行（familyName, givenName）→ 冪等

## 期待結果
- GET: 200 + { name, email }
- POST: 200 + { name, email }（登録または更新の結果）
- 未認証: 401
- バリデーション不備: 400 + { code, message }
