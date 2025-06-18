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
- `internal/middleware/`: Auth0 JWT middleware 
- `internal/types/`: Generated OpenAPI types (models only)
- `internal/utils/`: Auth0 user info utilities

**Key Design Decisions:**
- oapi-codegen generates **types only** (not server handlers)
- Manual route registration in `api.RegisterRoutes()` for better control
- Auth0 JWT authentication for all routes
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
# Install dependencies
task pnpm-install
task tidy

# Generate types from OpenAPI schema
task openapi-gen
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
```

## OpenAPI Workflow

1. Edit `schema/main.tsp` to define API changes using TypeSpec
2. Run `task openapi-gen` to:
   - Compile TypeSpec to OpenAPI: `schema/openapi.yaml`
   - Generate frontend types: `apps/next-app/generated/schema.d.ts`
   - Generate backend types: `apps/go-api/internal/types/openapi.gen.go`
3. Update handlers in `apps/go-api/internal/api/` to use new types
4. Add routes manually in `api.RegisterRoutes()`

### TypeSpec Commands

```bash
# Compile TypeSpec to OpenAPI only
pnpm run typespec-compile

# Full generation pipeline (TypeSpec → OpenAPI → client/server types)
task openapi-gen
```

## Environment Setup

Create environment files based on `.env.example`:
- `.env.local` for local development
- `.env.development`, `.env.production` as needed

Required environment variables for Go API:
- `AUTH0_DOMAIN`: Auth0 tenant domain
- `AUTH0_AUDIENCE`: Auth0 API identifier
- `GO_ENV`: Environment name (defaults to "local")