# Quickstart: /user API

この手順はローカル開発環境で /user の契約と最小動作を確認するためのものです。

## 前提
- Docker が起動済み
- apps/go-api/.env.local に Auth0 と DB の設定済み
- 有効な JWT トークンが取得済み

## Environment Setup

### Required Environment Variables
Create `apps/go-api/.env.local` with the following:
```bash
AUTH0_DOMAIN=your-tenant.auth0.com
AUTH0_AUDIENCE=your-api-identifier
DATABASE_URL=postgres://postgres:password@localhost:5432/canalia_dev?sslmode=disable
GO_ENV=local
```

## 手順（詳細）

### 1. データベースセットアップ
```bash
# データベース起動とマイグレーション適用
task db:dev:setup
```

### 2. OpenAPI型生成
```bash
# バックエンドとフロントエンドの型を生成
task openapi-gen
```

### 3. サーバ起動
```bash
# 開発サーバを起動
task dev:server
```

### 4. JWT トークン取得
Auth0から有効なJWTトークンを取得し、以下の環境変数に設定:
```bash
export JWT_TOKEN="your-jwt-token-here"
```

### 5. API エンドポイントテスト

#### GET /user - ユーザー情報取得
```bash
# 正常ケース: 登録済みユーザーの情報を取得
curl -X GET "http://localhost:8080/user" \
  -H "Authorization: Bearer $JWT_TOKEN" \
  -H "Content-Type: application/json"

# Expected Response (200):
{
  "id": 1,
  "email": "user@example.com",
  "contactEmail": "contact@example.com",
  "isVerified": true,
  "phoneNumber": "090-1234-5678",
  "familyName": "田中",
  "givenName": "太郎",
  "isActive": true,
  "createdAt": "2025-09-08T12:00:00Z",
  "updatedAt": "2025-09-08T12:00:00Z"
}

# エラーケース: 未認証
curl -X GET "http://localhost:8080/user" \
  -H "Content-Type: application/json"

# Expected Response (401):
# Unauthorized

# エラーケース: ユーザー未登録
# Expected Response (404):
{
  "code": 404,
  "message": "User not found"
}
```

#### POST /user - ユーザー登録（冪等）
```bash
# 正常ケース: 新規ユーザー登録
curl -X POST "http://localhost:8080/user" \
  -H "Authorization: Bearer $JWT_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "familyName": "田中",
    "givenName": "太郎",
    "contactEmail": "contact@example.com",
    "phoneNumber": "090-1234-5678"
  }'

# Expected Response (200):
{
  "id": 1,
  "email": "user@example.com",
  "contactEmail": "contact@example.com",
  "isVerified": false,
  "phoneNumber": "090-1234-5678",
  "familyName": "田中",
  "givenName": "太郎",
  "isActive": true,
  "createdAt": "2025-09-08T12:00:00Z",
  "updatedAt": "2025-09-08T12:00:00Z"
}

# 正常ケース: 必須フィールドのみ
curl -X POST "http://localhost:8080/user" \
  -H "Authorization: Bearer $JWT_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "familyName": "佐藤",
    "givenName": "花子"
  }'

# 冪等操作: 同じユーザーの情報を更新
curl -X POST "http://localhost:8080/user" \
  -H "Authorization: Bearer $JWT_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "familyName": "田中",
    "givenName": "太郎",
    "contactEmail": "updated@example.com",
    "phoneNumber": "080-9876-5432"
  }'

# バリデーションエラー: 必須フィールド不足
curl -X POST "http://localhost:8080/user" \
  -H "Authorization: Bearer $JWT_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "givenName": "太郎"
  }'

# Expected Response (400):
{
  "code": 400,
  "message": "Validation failed: Key: 'UserRegistrationRequest.FamilyName' Error:Field validation for 'FamilyName' failed on the 'required' tag"
}

# バリデーションエラー: 不正なメール形式
curl -X POST "http://localhost:8080/user" \
  -H "Authorization: Bearer $JWT_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "familyName": "田中",
    "givenName": "太郎",
    "contactEmail": "invalid-email"
  }'

# Expected Response (400):
{
  "code": 400,
  "message": "Validation failed: Key: 'UserRegistrationRequest.ContactEmail' Error:Field validation for 'ContactEmail' failed on the 'email' tag"
}

# バリデーションエラー: 文字数制限超過
curl -X POST "http://localhost:8080/user" \
  -H "Authorization: Bearer $JWT_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "familyName": "'"$(printf 'あ%.0s' {1..101})"'",
    "givenName": "太郎"
  }'

# Expected Response (400):
{
  "code": 400,
  "message": "Validation failed: Key: 'UserRegistrationRequest.FamilyName' Error:Field validation for 'FamilyName' failed on the 'max' tag"
}

# エラーケース: 未認証
curl -X POST "http://localhost:8080/user" \
  -H "Content-Type: application/json" \
  -d '{
    "familyName": "田中",
    "givenName": "太郎"
  }'

# Expected Response (401):
# Unauthorized
```

## 期待結果まとめ

### 正常ケース
- **GET /user**: 200 + User オブジェクト（登録済みの場合）
- **POST /user**: 200 + User オブジェクト（登録または更新の結果）

### エラーケース
- **認証なし**: 401 Unauthorized
- **ユーザー未登録** (GET): 404 + `{"code": 404, "message": "User not found"}`
- **バリデーションエラー** (POST): 400 + `{"code": 400, "message": "Validation failed: ..."}`
- **リクエスト形式エラー** (POST): 400 + `{"code": 400, "message": "Invalid request format"}`

## トラブルシューティング

### よくある問題
1. **JWT トークンの期限切れ**: Auth0から新しいトークンを取得してください
2. **データベース接続エラー**: `task db:dev:setup` でデータベースが起動しているか確認
3. **型生成エラー**: `task openapi-gen` を実行して最新の型を生成
4. **Auth0設定エラー**: `.env.local` の AUTH0_DOMAIN と AUTH0_AUDIENCE が正しいか確認

### デバッグ情報
```bash
# データベース接続確認
psql $DATABASE_URL -c "SELECT version();"

# サーバーログの確認
# サーバー起動時のコンソール出力を確認してエラーがないかチェック

# JWT トークンのデコード（デバッグ用）
# https://jwt.io でトークンの内容を確認（本番環境では使用しないこと）
```
