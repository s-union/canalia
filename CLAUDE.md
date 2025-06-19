# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

Canalia is a monorepo for a Japanese festival web portal (野田地区理大祭ウェブポータル) with:
- **Frontend**: Next.js app with Storybook component library
- **Backend**: Go API server with Echo framework
- **Shared**: OpenAPI schema for type-safe API communication

## Architecture

### Monorepo Structure
```
apps/
├── next-app/        # Next.js frontend with Storybook
└── go-api/          # Go API server
schema/
└── openapi.yaml     # Shared API specification
```

### Go API Architecture (apps/go-api)
Uses standard Go project layout with manual route control:
- `internal/api/`: HTTP handlers and route registration
- `internal/db/`: Database layer with sqlc-generated code and migrations
- `internal/middleware/`: Auth0 JWT middleware 
- `internal/types/`: Generated OpenAPI types (models only)
- `internal/utils/`: Auth0 user info utilities

**Key Design Decisions:**
- oapi-codegen generates **types only** (not server handlers)
- Manual route registration in `api.RegisterRoutes()` for better control
- Auth0 JWT authentication for all routes
- sqlc + golang-migrate for type-safe database operations
- Go 1.24 tool directives for dependency management
- Environment-specific `.env` files (`.env.local`, `.env.development`, etc.)

### Frontend Architecture (apps/next-app)
- Next.js 15 with App Router and Turbopack
- Auth0 for authentication
- Storybook for component development
- openapi-fetch for type-safe API calls
- Tailwind CSS with shadcn/ui components

## Development Commands

### Setup
```bash
# Complete automated setup (recommended)
task setup

# Manual setup
task pnpm-install
task tidy
task openapi-gen
```

### Database Development
```bash
# Setup development database (Docker + migrations)
task db:dev:setup

# Database operations
task db:dev:migrate:up      # Run migrations
task db:dev:migrate:down    # Rollback migrations
task db:migrate:create      # Create new migration
task db:generate            # Generate Go code from SQL
task db:dev:reset           # Reset database

# Docker operations
task docker:up              # Start PostgreSQL
task docker:down            # Stop database
task docker:clean           # Remove containers and volumes
```

### Development
```bash
# Run both frontend and backend
task dev

# Run only frontend
task dev:client

# Run only backend  
task dev:server
```

### Code Quality
```bash
# Format and lint (uses Biome)
pnpm run check
pnpm run format
pnpm run lint

# Next.js specific linting
pnpm run lint:next
```

### Frontend Specific
```bash
# In apps/next-app:
pnpm dev              # Development server with Turbopack
pnpm build            # Production build
pnpm storybook        # Component development
```

### Backend Specific
```bash
# In apps/go-api:
go run main.go        # Run server
go build              # Build binary
task tidy             # Clean dependencies

# Database tools (Go 1.24 tool directives)
go tool migrate -h    # Migration commands
go tool sqlc -h       # Code generation commands
```

## OpenAPI Workflow

1. Edit `schema/openapi.yaml` to define API changes
2. Run `task openapi-gen` to update generated types:
   - Frontend: `apps/next-app/generated/schema.d.ts`
   - Backend: `apps/go-api/internal/types/openapi.gen.go`
3. Update handlers in `apps/go-api/internal/api/` to use new types
4. Add routes manually in `api.RegisterRoutes()`

## Database Workflow

### Development Database
Uses Docker Compose with PostgreSQL 16 for development:
1. `task db:dev:setup` - Start database and run migrations
2. Add SQL queries to `internal/db/queries/` 
3. `task db:generate` - Generate Go code with sqlc
4. Use generated types in handlers

### Migration Management
1. `task db:migrate:create <name>` - Create new migration files
2. Edit `.up.sql` and `.down.sql` files in `internal/db/migrations/`
3. `task db:dev:migrate:up` - Apply migrations
4. Use `task db:dev:migrate:down` to rollback if needed

**Database Design:**
- Based on `database-design.md` with ER diagram
- Uses PostgreSQL enums for type safety
- Excludes unnecessary timestamps on static/intermediate tables
- Auth0 email-based user identification (no university_id)

## Environment Setup

Create environment files based on `.env.example`:
- `.env.local` for local development
- `.env.development`, `.env.production` as needed

Required environment variables for Go API:
- `AUTH0_DOMAIN`: Auth0 tenant domain
- `AUTH0_AUDIENCE`: Auth0 API identifier
- `DATABASE_URL`: PostgreSQL connection string
- `GO_ENV`: Environment name (defaults to "local")