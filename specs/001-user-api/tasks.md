# Tasks: ユーザー登録APIとユーザー情報確認API（/user）

**Input**: Design documents from `/specs/001-user-api/`
**Prerequisites**: plan.md (required), research.md, data-model.md, contracts/

## Execution Flow (main)
```
1. Load plan.md from feature directory
   → Extract: tech stack, libraries, structure
2. Load optional design documents:
   → data-model.md: Extract entities → model tasks
   → contracts/: Each file → contract test task
   → research.md: Extract decisions → setup tasks
3. Generate tasks by category:
   → Setup, Tests, Core, Integration, Polish
4. Apply task rules:
   → Different files = [P]
   → Tests before implementation (TDD)
5. Number tasks sequentially (T001, T002...)
6. Generate dependency notes and parallel examples
7. Return: SUCCESS (tasks ready for execution)
```

## Phase 3.1: Setup
- [ ] T001 Ensure Docker DB up and migrations applied (Taskfile)  
      Notes: `task db:dev:setup`
- [ ] T002 Merge `specs/001-user-api/contracts/user.yaml` into `schema/openapi.yaml`  
      Files: `schema/openapi.yaml`
- [ ] T003 Regenerate OpenAPI types (backend + frontend)  
      Notes: `task openapi-gen`（backend: `apps/go-api/internal/types/openapi.gen.go`, frontend: `apps/next-app/generated/schema.d.ts`）

## Phase 3.2: Tests First (TDD)
- [ ] T004 [P] Contract test GET /user  
      Goal: 200 schema(User), 401 未認証  
      Files: backend test placeholder（後で `apps/go-api` に統合テストを追加）
- [ ] T005 [P] Contract test POST /user  
      Goal: 200 schema(User), 400 バリデーション, 401 未認証
- [ ] T006 [P] Integration test: Register then Get self  
      Steps: POST /user → GET /user → 整合性

## Phase 3.3: Core Implementation
- [ ] T007 Create sqlc query: Upsert user by email  
      Files: `apps/go-api/internal/db/queries/user.sql`  
      SQL: `INSERT ... ON CONFLICT(email) DO UPDATE ... RETURNING *`
- [ ] T008 Create sqlc query: Get user by email  
      Files: `apps/go-api/internal/db/queries/user.sql`  
      SQL: `SELECT * FROM users WHERE email = $1`
- [ ] T009 Add request DTO with validator tags  
      Files: `apps/go-api/internal/api/user.go`（または新規 DTO ファイル）
- [ ] T010 Implement POST /user handler（validation → sqlc upsert → response）  
      Files: `apps/go-api/internal/api/user.go`
- [ ] T011 Implement GET /user handler（認証ユーザーのemailでSELECT）  
      Files: `apps/go-api/internal/api/user.go`
- [ ] T012 Wire route for POST /user  
      Files: `apps/go-api/internal/api/handler.go`

## Phase 3.4: Integration
- [ ] T013 Ensure Auth0 JWT middleware guards both endpoints  
      Files: `apps/go-api/internal/middleware/jwt.go`, route registration
- [ ] T014 Map sqlc model → OpenAPI User（PII最小）  
      Files: `apps/go-api/internal/api/user.go`
- [ ] T015 Error shaping: `{code,message}` for 400/401  
      Files: `apps/go-api/internal/api/user.go`

## Phase 3.5: Polish
- [ ] T016 [P] Unit tests for validator rules（names length, email format, etc.）
- [ ] T017 [P] Update `quickstart.md` with curl examples and environment notes
- [ ] T018 [P] Lint/format and ensure generated code up-to-date（`task check`）

## Dependencies
- Setup (T001-T003) → Tests (T004-T006) → Core (T007-T012) → Integration (T013-T015) → Polish (T016-T018)
- T007, T008 block handler implementations（T010, T011）
- T012 depends on route wiring

## Parallel Example
```
# Run contract tests in parallel
T004, T005, T006 [P]
```

## Notes
- OpenAPI is the source of truth; keep `schema/openapi.yaml` authoritative
- sqlc queries must compile via `task db:generate` after addition
- Responses limit PII to `name` and `email`
