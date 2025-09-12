# Research: ユーザー登録APIとユーザー情報確認API（/user）

## Decisions
- Contract source of truth: OpenAPI（`schema/openapi.yaml` を更新、差分は `specs/001-user-api/contracts/user.yaml` 管理）
- Validation: go-playground/validator v10（構造体タグで必須/形式）
- Idempotency: メールアドレスを一意キーに見なす。登録はUPSERT方針（INSERT ON CONFLICT(email) DO UPDATE…）[確認: users.email UNIQUE あり]
- Auth: Auth0 JWT（既存ミドルウェア利用）。GET/POSTとも認証必須。スコープはRole: normal想定
- Response model: PII最小（name, email を基本。追加項目は段階導入）
- Error model: `{ "code": number, "message": string }` に統一

## Rationale
- OpenAPI: 型生成（oapi-codegen/types）とフロントの型整合性維持のため
- validator: Echoと相性が良く、構造体タグで簡潔に定義可能
- sqlc: クエリから型生成し、実装を薄く安全にできる
- UPSERT: 冪等化をシンプルに実現できる

## Alternatives Considered
- 手書きバリデーション → 複雑で漏れやすい
- ORM導入 → 本プロジェクトはsqlc方針
- ユーザー主キーにAuth0 subを使用 → 現状は email を主識別子としているため段階移行

## Open Questions (NEEDS CLARIFICATION)
- 返却項目の確定（電話番号や連絡用メールを返すか）
- 大学メールドメイン制約（必須か、許容ドメインは何か）
- メール未確認（email_verified=false）の取り扱い（登録/参照の是非）

## Security Notes
- JWT検証失敗→401
- 自ユーザー以外の情報は返さない
- レートリミットは将来検討（CloudやAPI Gateway側）
