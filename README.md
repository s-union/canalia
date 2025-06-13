# Canalia

野田地区理大祭ウェブポータル - Next.js フロントエンドと Go API サーバーを組み合わせたモノレポプロジェクト

## 🚀 クイックスタート

### 必要なツール

以下のツールがインストールされている必要があります：

- **Node.js** v20 LTS以降 ([ダウンロード](https://nodejs.org/))
- **pnpm** v10以降 (`npm install -g pnpm`)
- **Go** v1.24以降 ([ダウンロード](https://golang.org/dl/))
- **Taskfile** ([インストール手順](https://taskfile.dev/installation/))

### セットアップ

```bash
# 1. リポジトリをクローン
git clone git@github.com:s-union/canalia.git
cd canalia

# 2. 自動セットアップを実行（推奨）
task setup
```

これで以下が自動実行されます：
- ✅ 必要なツールの確認
- 📦 依存関係のインストール  
- ⚙️ 環境設定ファイルの作成
- 🔧 型定義の生成
- 🧪 セットアップの検証

## `.env` ファイルの設定

`.env.exmple` を見ながらセットアップ

## 🛠️ 開発

### 開発サーバーの起動

```bash
# フロントエンドとバックエンドを同時起動
task dev

# 個別起動
task dev:client    # Next.js フロントエンド
task dev:server    # Go API サーバー
```

### その他の開発タスク

```bash
# コンポーネント開発
task storybook

# コードフォーマット
task format

# リント検査
task lint

# テスト実行
task test

# ビルド
task build
```

## ⚙️ 環境設定

初回セットアップ後、`apps/go-api/.env.local` を編集してAuth0設定を行ってください：

```bash
AUTH0_DOMAIN=your-auth0-domain.auth0.com
AUTH0_CLIENT_ID=your-client-id
AUTH0_CLIENT_SECRET=your-client-secret
AUTH0_AUDIENCE=your-api-identifier
```

## 📋 利用可能なコマンド

すべてのタスクを確認：
```bash
task --list
```

## 🏗️ プロジェクト構成

```
canalia/
├── apps/
│   ├── next-app/        # Next.js フロントエンド
│   │   ├── src/
│   │   ├── components/  # UIコンポーネント
│   │   └── generated/   # 生成された型定義
│   └── go-api/          # Go API サーバー
│       ├── internal/    # プライベートコード
│       │   ├── api/     # APIハンドラー
│       │   ├── middleware/ # ミドルウェア
│       │   └── types/   # 生成された型定義
│       └── main.go
├── schema/
│   └── openapi.yaml     # API仕様書
└── Taskfile.yml         # ビルドタスク定義
```

## 🔄 開発ワークフロー

### API開発
1. `schema/openapi.yaml` でAPI仕様を定義
2. `task openapi-gen` で型定義を生成
3. `apps/go-api/internal/api/` でハンドラーを実装
4. `api.RegisterRoutes()` でルートを登録

### フロントエンド開発
1. `apps/next-app/src/components/` でコンポーネントを作成
2. `task storybook` でコンポーネントを開発
3. 生成された型定義を使用してAPI呼び出し

## 🧪 テスト

```bash
# すべてのテスト実行
task test

# コンポーネントテスト（Vitest）
pnpm -C apps/next-app test

# ビジュアルテスト（Storybook）
task storybook
```

## 📚 技術スタック

### フロントエンド
- **Next.js 15** - React フレームワーク
- **Tailwind CSS** - スタイリング
- **Storybook** - コンポーネント開発
- **Vitest** - テスト
- **openapi-fetch** - 型安全API呼び出し

### バックエンド
- **Go** - API サーバー
- **Echo** - Web フレームワーク
- **Auth0** - 認証
- **oapi-codegen** - OpenAPI型生成

### 開発ツール
- **Taskfile** - タスクランナー
- **Biome** - リンター・フォーマッター
- **pnpm** - パッケージマネージャー
- **Lefthook** - Git フック

## 🆘 トラブルシューティング

### セットアップで問題が発生した場合

```bash
# 前提条件を再確認
task check-prerequisites

# クリーンアップして再セットアップ
task clean
task setup
```

### 型生成でエラーが発生した場合

```bash
# OpenAPI スキーマを確認
task openapi-gen
```

### 開発サーバーが起動しない場合

```bash
# 環境設定を確認
task check-env

# 依存関係を再インストール
task clean
task install-deps
```
