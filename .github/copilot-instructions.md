# Canalia Development Instructions

**ALWAYS follow these instructions first. Only search for additional information or context if these instructions are incomplete or contain errors.**

Canalia is a monorepo for a Japanese festival web portal with Next.js frontend, Go API backend, and shared OpenAPI schema for type-safe communication.

## Required Prerequisites

Install these tools in order. **NEVER SKIP any prerequisite.**

1. **Node.js 20+ LTS**: Download from https://nodejs.org/dist/v20.19.4/node-v20.19.4-linux-x64.tar.xz
2. **pnpm 10+**: `npm install -g pnpm`  
3. **Go 1.24+**: Download from https://golang.org/dl/go1.24.5.linux-amd64.tar.gz
4. **Task**: Download latest from https://github.com/go-task/task/releases/latest/download/task_linux_amd64.tar.gz
5. **Docker**: Install from https://docs.docker.com/get-docker/

## Critical Setup Instructions

**IMPORTANT**: Setup takes 7-10 minutes. NEVER CANCEL during setup or builds.

### 1. Fresh Clone Setup
```bash
# Clone repository
git clone https://github.com/s-union/canalia.git
cd canalia

# Verify prerequisites (takes 10 seconds)
task check-prerequisites

# Complete setup (takes 7-10 minutes, NEVER CANCEL)
task setup
```

**TIMEOUT WARNING**: Use 15+ minute timeouts for `task setup`. If it fails due to network issues with Google Fonts, see troubleshooting section.

### 2. Environment Configuration

**ALWAYS create environment files after setup:**

```bash
# Backend environment
cp apps/go-api/.env.example apps/go-api/.env.local

# Edit the .env.local file with your Auth0 settings:
# AUTH0_DOMAIN=your-domain.auth0.com
# AUTH0_AUDIENCE=your-api-identifier
```

## Build and Development Commands

**Build Times and Timeouts:**
- Frontend build: 5-10 seconds
- Backend build: 30 seconds  
- Full build: 1 minute total
- **Set timeouts to 5+ minutes for all builds to avoid cancellation**

### Essential Commands
```bash
# Build everything (1 minute, NEVER CANCEL)
task build

# Run code quality checks (4 seconds)
task lint

# Format code
task format
```

### Development Servers

**CRITICAL**: Start database BEFORE running servers that need it.

```bash
# Database setup (5 seconds + 2 seconds migration)
task db:dev:setup

# Start all services together (recommended)
task dev

# OR start individually:
task dev:client    # Next.js at localhost:3000 (starts in 1.6 seconds)
task dev:server    # Go API at localhost:8080 (starts instantly)
task storybook     # Component development at localhost:6006 (starts in 1 second)
```

## Database Development

**Docker Required** - Database uses PostgreSQL 16 in Docker.

```bash
# Initial database setup (NEVER CANCEL, takes 10 seconds total)
task db:dev:setup

# Database operations
task db:dev:migrate:up      # Apply migrations (2 seconds)
task db:dev:migrate:down    # Rollback migrations
task db:migrate:create <name>  # Create new migration
task db:generate           # Generate Go code from SQL (1 second)
task db:dev:clean          # Clear test data
task db:dev:reset          # Reset entire database

# Docker management
task docker:up:db          # Start PostgreSQL only
task docker:down           # Stop all services
task docker:clean          # Remove containers and volumes
```

## OpenAPI Workflow

**ALWAYS regenerate types after schema changes:**

```bash
# 1. Edit schema/openapi.yaml
# 2. Generate types (takes 1 second)
task openapi-gen

# This updates:
# - apps/next-app/generated/schema.d.ts (Frontend types)
# - apps/go-api/internal/types/openapi.gen.go (Backend types)

# 3. Update handlers in apps/go-api/internal/api/
# 4. Register routes manually in api.RegisterRoutes()
```

## Testing and Quality

**Current Test Status**: Vitest configured for Storybook testing, but no unit tests exist yet.

```bash
# Run Storybook tests (when tests exist)
cd apps/next-app && pnpm vitest

# Visual testing via Storybook
task storybook

# Code quality (4 seconds total)
task lint    # Biome + ESLint checks
task format  # Auto-format code
```

## Validation Requirements

**ALWAYS test these scenarios after making changes:**

### 1. Build Validation
```bash
# Test full build works (NEVER CANCEL, 1 minute timeout minimum)
task build

# Verify both apps built successfully:
# - Frontend: apps/next-app/.next/ directory exists
# - Backend: apps/go-api/canalia binary exists
```

### 2. Development Environment Validation
```bash
# 1. Start database
task db:dev:setup

# 2. Start API (with proper .env.local)
task dev:server
# Should show: "⇨ http server started on [::]:8080"

# 3. Start frontend
task dev:client  
# Should show: "Ready in 1649ms" and run on localhost:3000

# 4. Test Storybook
task storybook
# Should show: "Local: http://localhost:6006/"
```

### 3. Database Validation
```bash
# Verify database operations work
task db:dev:setup
task db:dev:migrate:up
task db:generate
task db:dev:clean
```

## Troubleshooting

### Google Fonts Network Error
If `task setup` fails with "Failed to fetch Noto Sans JP from Google Fonts":

```bash
# Temporarily comment out Google Fonts import in apps/next-app/src/app/layout.tsx:
# Line 2: // import { Noto_Sans_JP } from 'next/font/google';
# Line 6-11: Comment out const notoSansJP = ...
# Line 25: Change to className="font-sans"

# Then run setup again
task setup

# Restore the fonts after setup completes
```

### Docker Compose Issues
Modern Docker uses `docker compose` not `docker-compose`. If Taskfile commands fail:

```bash
# Use modern Docker Compose syntax:
docker compose -f docker-compose.dev.yml up -d postgres
docker compose -f docker-compose.dev.yml down
```

### Missing Environment Variables
If Go API fails to start:

```bash
# Ensure .env.local exists in apps/go-api/
cp apps/go-api/.env.example apps/go-api/.env.local

# Edit with valid values (test values work for development):
# AUTH0_DOMAIN=test-domain.auth0.com
# AUTH0_AUDIENCE=test-api-identifier
```

## Project Structure

```
apps/
├── next-app/           # Next.js 15 frontend
│   ├── src/           # Source code
│   ├── .storybook/    # Storybook config
│   └── generated/     # OpenAPI generated types
└── go-api/            # Go API server  
    ├── internal/api/  # HTTP handlers
    ├── internal/db/   # Database layer (sqlc + migrations)
    └── main.go        # Server entry point
schema/
└── openapi.yaml       # Shared API specification
```

## Common File Locations

- **API handlers**: `apps/go-api/internal/api/`
- **Database queries**: `apps/go-api/internal/db/queries/`
- **Migrations**: `apps/go-api/internal/db/migrations/`
- **Frontend components**: `apps/next-app/src/components/`
- **Route registration**: `apps/go-api/internal/api/` (manual in RegisterRoutes)

## Performance Notes

- **Frontend build**: Uses Turbopack for fast development (1.6s startup)
- **Backend build**: Go compiles quickly (30 seconds)
- **Database**: PostgreSQL in Docker with hot reload
- **Code quality**: Biome for fast linting/formatting (4 seconds)

## CI/CD Notes

- **Chromatic**: Visual testing for Storybook components
- **Auth0**: Required for all API routes (configure in .env.local)
- **Networking**: External font loading may fail in restricted environments

**Remember**: ALWAYS run `task lint` before committing. NEVER CANCEL long-running builds or setup commands.