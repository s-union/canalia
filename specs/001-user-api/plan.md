# Implementation Plan: ユーザー登録APIとユーザー情報確認API（/user）

**Branch**: `001-user-api` | **Date**: 2025-09-09 | **Spec**: /home/hayato/projects/canalia/specs/001-user-api/spec.md
**Input**: Feature specification from `/specs/001-user-api/spec.md`

## Execution Flow (/plan command scope)
```
1. Load feature spec from Input path
   → If not found: ERROR "No feature spec at {path}"
2. Fill Technical Context (scan for NEEDS CLARIFICATION)
   → Detect Project Type from context (web=frontend+backend, mobile=app+api)
   → Set Structure Decision based on project type
3. Evaluate Constitution Check section below
   → If violations exist: Document in Complexity Tracking
   → If no justification possible: ERROR "Simplify approach first"
   → Update Progress Tracking: Initial Constitution Check
4. Execute Phase 0 → research.md
   → If NEEDS CLARIFICATION remain: ERROR "Resolve unknowns"
5. Execute Phase 1 → contracts, data-model.md, quickstart.md, agent-specific template file (e.g., `CLAUDE.md` for Claude Code, `.github/copilot-instructions.md` for GitHub Copilot, or `GEMINI.md` for Gemini CLI).
6. Re-evaluate Constitution Check section
   → If new violations: Refactor design, return to Phase 1
   → Update Progress Tracking: Post-Design Constitution Check
7. Plan Phase 2 → Describe task generation approach (DO NOT create tasks.md)
8. STOP - Ready for /tasks command
```

**IMPORTANT**: The /plan command STOPS at step 7. Phases 2-4 are executed by other commands:
- Phase 2: /tasks command creates tasks.md
- Phase 3-4: Implementation execution (manual or via tools)

## Summary
本機能は「/user」でのユーザー登録（POST）とユーザー情報取得（GET）を提供する。Auth0のJWTで認証し、OpenAPIにて契約を管理、go-playground/validatorによる入力検証、sqlc による型安全なDBアクセスを徹底する。登録はメールアドレス（Auth0由来）を一意キーとして冪等（再送で重複レコードを作らない）。GETは自身の完全な情報（id, email, isVerified, contactEmail, phoneNumber, familyName, givenName, isActive, createdAt, updatedAt）を返す。

## Technical Context
**Language/Version**: Go 1.24  
**Primary Dependencies**: Echo, Auth0 JWT Middleware, oapi-codegen（typesのみ）, go-playground/validator v10, sqlc, golang-migrate  
**Storage**: PostgreSQL 16（Docker Compose）  
**Testing**: Go test（contract/integration）、Task での開発補助  
**Target Platform**: Linux（Docker dev 環境）  
**Project Type**: web（frontend: Next.js, backend: Go API）  
**Performance Goals**: p95 < 200ms（/user GET/POST）、DBクエリは単一SELECT/UPSERTで完結  
**Constraints**: 認証必須（Bearer JWT）、最小権限（自分の情報のみ）、PII最小化  
**Scale/Scope**: 初期ユーザー数〜数千規模想定（拡張容易性重視）

## Constitution Check
*GATE: Must pass before Phase 0 research. Re-check after Phase 1 design.*

**Simplicity**:
- Projects: 2（apps/go-api, apps/next-app）
- Using framework directly: Yes（Echoを直接利用、oapi-codegenはtypesのみ）
- Single data model: Yes（DBテーブル users に集約。OpenAPIのUserは同一概念）
- Avoiding patterns: Yes（Repository等は導入しない。sqlcを直接利用）

**Architecture**:
- Featureはアプリ直下で最小追加。共通は既存ミドルウェア/型を再利用
- Libraries: Echo（HTTP）、Auth0 Middleware（JWT検証）、oapi-codegen（型）、sqlc（DB）
- CLI: なし
- Docs: specs/ 以下に生成物を配置

**Testing (NON-NEGOTIABLE)**:
- RED-GREEN-Refactor: 準拠
- Commits: Contract/Integration テストを先に追加
- Order: Contract→Integration→Unit
- Real deps: Dev DB（Docker）を使用
- Integration: 新規POST/更新GETに対して追加
- Forbidden: 実装先行は不可

**Observability**:
- Echoのロガーで構造化ログ（最低限）
- エラーはOpenAPIのエラーフォーマットに整形（code, message）

**Versioning**:
- OpenAPIバージョン 0.1.0を基底。/userスキーマ変更はマイナー更新
- 破壊的変更は contracts/ で検出可能にして段階適用

## Project Structure

### Documentation (this feature)
```
specs/[###-feature]/
├── plan.md              # This file (/plan command output)
├── research.md          # Phase 0 output (/plan command)
├── data-model.md        # Phase 1 output (/plan command)
├── quickstart.md        # Phase 1 output (/plan command)
├── contracts/           # Phase 1 output (/plan command)
└── tasks.md             # Phase 2 output (/tasks command - NOT created by /plan)
```

### Source Code (repository root)
```
# Option 1: Single project (DEFAULT)
src/
├── models/
├── services/
├── cli/
└── lib/

tests/
├── contract/
├── integration/
└── unit/

# Option 2: Web application (when "frontend" + "backend" detected)
backend/
├── src/
│   ├── models/
│   ├── services/
│   └── api/
└── tests/

frontend/
├── src/
│   ├── components/
│   ├── pages/
│   └── services/
└── tests/

# Option 3: Mobile + API (when "iOS/Android" detected)
api/
└── [same as backend above]

ios/ or android/
└── [platform-specific structure]
```

**Structure Decision**: Web application（backend: apps/go-api, frontend: apps/next-app）

## Phase 0: Outline & Research
1. **Extract unknowns from Technical Context** above:
   - For each NEEDS CLARIFICATION → research task
   - For each dependency → best practices task
   - For each integration → patterns task

2. **Generate and dispatch research agents**:
   ```
   For each unknown in Technical Context:
     Task: "Research {unknown} for {feature context}"
   For each technology choice:
     Task: "Find best practices for {tech} in {domain}"
   ```

3. **Consolidate findings** in `research.md` using format:
   - Decision: [what was chosen]
   - Rationale: [why chosen]
   - Alternatives considered: [what else evaluated]

**Output**: research.md（OpenAPI契約の拡張、検証ルール、冪等化方針、エラーモデル、セキュリティを決定）

## Phase 1: Design & Contracts
*Prerequisites: research.md complete*

1. **Extract entities from feature spec** → `data-model.md`:
   - Entity name, fields, relationships
   - Validation rules from requirements
   - State transitions if applicable

2. **Generate API contracts** from functional requirements:
   - POST /user（登録）・GET /user（取得）
   - OpenAPIに User（完全なDBフィールド：id, email, isVerified, contactEmail, phoneNumber, familyName, givenName, isActive, createdAt, updatedAt）, UserRegistrationInput（入力用：familyName, givenName, contactEmail, phoneNumber）を定義
   - `/contracts/user.yaml` に差分契約を出力（後で `schema/openapi.yaml` に統合）

3. **Generate contract tests** from contracts:
   - 各エンドポイント1本（POST, GET）
   - スキーマ・ステータスコード（200/400/401など）を検証（初回は失敗）

4. **Extract test scenarios** from user stories:
   - Each story → integration test scenario
   - Quickstart test = story validation steps

5. **Update agent file incrementally** (O(1) operation):
   - Run `/scripts/update-agent-context.sh [claude|gemini|copilot]` for your AI assistant
   - If exists: Add only NEW tech from current plan
   - Preserve manual additions between markers
   - Update recent changes (keep last 3)
   - Keep under 150 lines for token efficiency
   - Output to repository root

**Output**: data-model.md, /contracts/*, failing tests（別途作成）, quickstart.md, agent-specific file

## Phase 2: Task Planning Approach
*This section describes what the /tasks command will do - DO NOT execute during /plan*

**Task Generation Strategy**:
- Load `/templates/tasks-template.md` as base
- Generate tasks from Phase 1 design docs (contracts, data model, quickstart)
- Each contract → contract test task [P]
- Each entity → model creation task [P] 
- Each user story → integration test task
- Implementation tasks to make tests pass

**Ordering Strategy**:
- TDD order: Tests before implementation 
- Dependency order: Models before services before UI
- Mark [P] for parallel execution (independent files)

**Estimated Output**: 15-25 tasks（OpenAPI更新→型生成→sqlcクエリ→ハンドラ→検証→テスト修正）

**IMPORTANT**: This phase is executed by the /tasks command, NOT by /plan

## Phase 3+: Future Implementation
*These phases are beyond the scope of the /plan command*

**Phase 3**: Task execution (/tasks command creates tasks.md)  
**Phase 4**: Implementation (execute tasks.md following constitutional principles)  
**Phase 5**: Validation (run tests, execute quickstart.md, performance validation)

## Complexity Tracking
*Fill ONLY if Constitution Check has violations that must be justified*

| Violation | Why Needed | Simpler Alternative Rejected Because |
|-----------|------------|-------------------------------------|
| [e.g., 4th project] | [current need] | [why 3 projects insufficient] |
| [e.g., Repository pattern] | [specific problem] | [why direct DB access insufficient] |


## Progress Tracking
*This checklist is updated during execution flow*

**Phase Status**:
- [x] Phase 0: Research complete (/plan command)
- [x] Phase 1: Design complete (/plan command)
- [ ] Phase 2: Task planning complete (/plan command - describe approach only)
- [ ] Phase 3: Tasks generated (/tasks command)
- [ ] Phase 4: Implementation complete
- [ ] Phase 5: Validation passed

**Gate Status**:
- [x] Initial Constitution Check: PASS
- [x] Post-Design Constitution Check: PASS
- [ ] All NEEDS CLARIFICATION resolved（ドメイン制約の詳細は残件）
- [ ] Complexity deviations documented

---
*Based on Constitution v2.1.1 - See `/memory/constitution.md`*