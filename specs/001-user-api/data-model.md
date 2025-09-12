# Data Model: users

基底テーブル: apps/go-api/internal/db/migrations/001_init.up.sql の `users`

## Fields
- id: SERIAL PK
- email: VARCHAR(255) NOT NULL UNIQUE
- contact_email: VARCHAR(255) Nullable
- is_verified: BOOLEAN NOT NULL DEFAULT FALSE
- phone_number: VARCHAR(20) Nullable
- family_name: VARCHAR(100) NOT NULL
- given_name: VARCHAR(100) NOT NULL
- is_active: BOOLEAN NOT NULL DEFAULT TRUE
- created_at: timestamptz NOT NULL DEFAULT now()
- updated_at: timestamptz NOT NULL DEFAULT now()

## Validation (go-playground/validator タグ想定)
- family_name: required,max=100
- given_name: required,max=100
- contact_email: omitempty,email
- phone_number: omitempty,max=20  （[NEEDS CLARIFICATION: 許容パターン E.164 など]）
- email: Auth0提供値を使用（入力では受け取らない）

## Constraints
- email UNIQUE インデックス: `idx_users_email`
- 冪等化: email をキーに UPSERT（ON CONFLICT (email) DO UPDATE）

## Relations（将来）
- leaders, user_permissions, answered_forms, qa_answers などに参照される

## OpenAPI 対応
- 出力 `User` は name, email を返す（PII最小）
- 登録入力 `UserRegistrationInput` は familyName, givenName（必須）、contactEmail, phoneNumber（任意）
