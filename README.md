# Social Media Platform

A full-stack social media platform built with Go backend and React frontend, featuring user authentication, posts, comments, followers, and real-time interactions.

## 🚀 Features

- **User Management**: Registration, authentication, and profile management
- **Posts & Comments**: Create, read, update, and delete posts with commenting system
- **Social Features**: Follow/unfollow users, personalized feed
- **Authentication**: JWT-based authentication with secure token management
- **Rate Limiting**: Built-in rate limiting for API protection
- **Caching**: Redis integration for improved performance
- **Email Integration**: SendGrid integration for notifications
- **API Documentation**: Swagger/OpenAPI documentation
- **Database Migrations**: Structured database schema management
- **Testing**: Comprehensive test suite

## 🛠 Tech Stack

### Backend
- **Language**: Go 1.23.4
- **Framework**: Chi Router
- **Database**: PostgreSQL
- **Cache**: Redis
- **Authentication**: JWT tokens
- **Email**: SendGrid
- **Documentation**: Swagger/Swaggo
- **Logging**: Zap logger

### Frontend
- **Framework**: React 19.0.0
- **Build Tool**: Vite
- **Routing**: React Router DOM
- **Language**: JavaScript/TypeScript

### DevOps & Tools
- **Containerization**: Docker & Docker Compose
- **Hot Reload**: Air (Go) & Vite (React)
- **CI/CD**: GitHub Actions
- **Database Migrations**: golang-migrate
- **Testing**: Go testing framework

## 📁 Project Structure

```
Social/
├── cmd/                          # Application entry points
│   ├── api/                      # REST API server
│   │   ├── main.go              # Main application entry
│   │   ├── api.go               # API configuration and routes
│   │   ├── auth.go              # Authentication handlers
│   │   ├── users.go             # User management endpoints
│   │   ├── posts.go             # Post management endpoints
│   │   ├── feed.go              # Feed generation endpoints
│   │   ├── middleware.go        # HTTP middleware
│   │   ├── errors.go            # Error handling
│   │   ├── json.go              # JSON utilities
│   │   ├── health.go            # Health check endpoint
│   │   └── *_test.go            # API tests
│   └── migrate/                 # Database migration tools
│       ├── migrations/          # SQL migration files
│       └── seed/               # Database seeding
├── internal/                    # Private application code
│   ├── auth/                   # Authentication logic
│   │   ├── auth.go             # Auth interfaces
│   │   ├── jwt.go              # JWT implementation
│   │   └── mocks.go            # Auth mocks for testing
│   ├── db/                     # Database connection and utilities
│   │   ├── db.go               # Database connection setup
│   │   └── seed.go             # Database seeding logic
│   ├── env/                    # Environment configuration
│   │   └── env.go              # Environment variable helpers
│   ├── mailer/                 # Email service
│   │   ├── mailer.go           # Mailer interface
│   │   ├── sendgrid.go         # SendGrid implementation
│   │   └── templates/          # Email templates
│   ├── ratelimiter/            # Rate limiting
│   │   ├── ratelimiter.go      # Rate limiter interface
│   │   └── fixed-window.go     # Fixed window implementation
│   └── store/                  # Data access layer
│       ├── storage.go          # Storage interfaces
│       ├── users.go            # User data operations
│       ├── posts.go            # Post data operations
│       ├── comments.go         # Comment data operations
│       ├── followers.go        # Follower relationships
│       ├── role.go             # User roles
│       ├── pagination.go       # Pagination utilities
│       ├── mocks.go            # Store mocks for testing
│       └── cache/              # Caching layer
├── web/                        # Frontend React application
│   ├── src/                    # React source code
│   │   ├── App.jsx             # Main App component
│   │   ├── main.jsx            # React entry point
│   │   ├── ConfirmationPage.tsx # Email confirmation page
│   │   ├── App.css             # App styles
│   │   ├── index.css           # Global styles
│   │   └── assets/             # Static assets
│   ├── public/                 # Public assets
│   ├── package.json            # Node.js dependencies
│   ├── vite.config.js          # Vite configuration
│   ├── eslint.config.js        # ESLint configuration
│   └── index.html              # HTML template
├── docs/                       # API documentation
│   ├── docs.go                 # Generated Swagger docs
│   ├── swagger.json            # Swagger JSON spec
│   └── swagger.yaml            # Swagger YAML spec
├── scripts/                    # Utility scripts
│   ├── db_init.sql             # Database initialization
│   └── test_concurrency.go     # Concurrency testing
├── bin/                        # Compiled binaries
├── .github/                    # GitHub workflows
│   └── workflows/              # CI/CD pipelines
├── go.mod                      # Go module definition
├── go.sum                      # Go module checksums
├── makefile                    # Build automation
├── docker-compose.yml          # Docker services
├── .air.toml                   # Air hot reload config
├── .envrc                      # Environment variables
├── .gitignore                  # Git ignore rules
└── README.md                   # This file
```

## 🚦 Prerequisites

- **Go**: 1.23.4 or later
- **Node.js**: 18.0 or later
- **PostgreSQL**: 12.0 or later
- **Redis**: 6.0 or later (optional, for caching)
- **Docker & Docker Compose**: For containerized setup

## ⚡ Quick Start

### 1. Clone the Repository

```bash
git clone https://github.com/karthikbhandary2/Social.git
cd Social
```

### 2. Environment Setup

Create a `.envrc` file in the root directory:

```bash
# Database Configuration
export DB_ADDR="postgres://username:password@localhost/social?sslmode=disable"
export DB_MAX_OPEN_CONNS=30
export DB_MAX_IDLE_CONNS=30
export DB_MAX_IDLE_TIME="15m"

# Server Configuration
export ADDR=":8082"
export EXTERNAL_URL="localhost:8082"
export FRONTEND_URL="http://localhost:5173"
export ENV="development"

# Redis Configuration (optional)
export REDIS_ADDR="localhost:6379"
export REDIS_PW=""
export REDIS_DB=0
export REDIS_ENABLED=true

# Authentication
export AUTH_BASIC_USER="admin"
export AUTH_BASIC_PASS="admin"
export AUTH_TOKEN_SECRET="your-secret-key"

# Email Configuration (SendGrid)
export FROM_EMAIL="your-email@example.com"
export SENDGRID_API_KEY="your-sendgrid-api-key"

# Rate Limiting
export RATELIMITER_REQUESTS_COUNT=20
export RATE_LIMITER_ENABLED=true
```

### 3. Database Setup

#### Option A: Using Docker (Recommended)

```bash
# Start PostgreSQL with Docker Compose
docker-compose up -d db

# Wait for database to be ready, then run migrations
make migrate-up

# Seed the database with sample data
make seed
```

#### Option B: Local PostgreSQL

```bash
# Create database
createdb social

# Run migrations
make migrate-up

# Seed database
make seed
```

### 4. Install Dependencies

#### Backend Dependencies
```bash
go mod tidy
```

#### Frontend Dependencies
```bash
cd web
npm install
cd ..
```

### 5. Start the Application

#### Development Mode (with hot reload)

**Terminal 1 - Backend:**
```bash
# Install Air for hot reload (if not already installed)
go install github.com/cosmtrek/air@latest

# Start backend with hot reload
air
```

**Terminal 2 - Frontend:**
```bash
cd web
npm run dev
```

#### Production Mode

**Backend:**
```bash
go build -o bin/main cmd/api/*.go
./bin/main
```

**Frontend:**
```bash
cd web
npm run build
npm run preview
```

### 6. Access the Application

- **Frontend**: http://localhost:5173
- **Backend API**: http://localhost:8082
- **API Documentation**: http://localhost:8082/swagger/index.html
- **Health Check**: http://localhost:8082/v1/health

## 🔧 Development

### Available Make Commands

```bash
# Database Operations
make migrate-up              # Run all pending migrations
make migrate-down           # Rollback last migration
make migrate-force ARG=1    # Force migration to specific version
make seed                   # Seed database with sample data

# Documentation
make gen-docs              # Generate Swagger documentation

# Testing
make test                  # Run all tests

# Create new migration
make migration create_users_table
```

### API Endpoints

#### Authentication
- `POST /v1/authentication/user` - User login
- `POST /v1/users` - User registration
- `PUT /v1/users/activate/:token` - Account activation

#### Users
- `GET /v1/users/:id` - Get user profile
- `PUT /v1/users/:id/follow` - Follow user
- `PUT /v1/users/:id/unfollow` - Unfollow user

#### Posts
- `GET /v1/posts` - List posts with pagination
- `POST /v1/posts` - Create new post
- `GET /v1/posts/:id` - Get specific post
- `PATCH /v1/posts/:id` - Update post
- `DELETE /v1/posts/:id` - Delete post

#### Comments
- `GET /v1/posts/:id/comments` - Get post comments
- `POST /v1/posts/:id/comments` - Add comment

#### Feed
- `GET /v1/users/:id/feed` - Get user's personalized feed

### Testing

```bash
# Run all tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Run specific test
go test ./cmd/api -v

# Run tests with race detection
go test -race ./...
```

### Database Migrations

Create a new migration:
```bash
make migration create_new_table
```

This creates two files in `cmd/migrate/migrations/`:
- `000X_create_new_table.up.sql` - Forward migration
- `000X_create_new_table.down.sql` - Rollback migration

## 🐳 Docker Deployment

### Full Stack with Docker Compose

```bash
# Build and start all services
docker-compose up --build

# Start in background
docker-compose up -d

# View logs
docker-compose logs -f

# Stop services
docker-compose down
```

### Individual Service Containers

```bash
# Build backend image
docker build -t social-api .

# Build frontend image
cd web && docker build -t social-web .
```

## 🧪 Testing

The project includes comprehensive tests:

- **Unit Tests**: Individual component testing
- **Integration Tests**: API endpoint testing
- **Concurrency Tests**: Race condition testing

### Test Structure
```
cmd/api/
├── api_test.go          # API integration tests
├── user_test.go         # User endpoint tests
└── test_utils.go        # Testing utilities
```

## 📊 Monitoring & Metrics

The application exposes metrics at `/debug/vars`:

- **Version**: Application version
- **Database**: Connection pool stats
- **Goroutines**: Active goroutine count
- **Memory**: Runtime memory statistics

## 🔒 Security Features

- **JWT Authentication**: Secure token-based authentication
- **Rate Limiting**: Configurable request rate limiting
- **CORS**: Cross-origin resource sharing configuration
- **Input Validation**: Request payload validation
- **SQL Injection Protection**: Parameterized queries
- **Password Hashing**: Bcrypt password hashing

## 🚀 Deployment

### Environment Variables for Production

```bash
export ENV="production"
export DB_ADDR="postgres://user:pass@prod-db:5432/social?sslmode=require"
export REDIS_ADDR="prod-redis:6379"
export EXTERNAL_URL="https://your-domain.com"
export FRONTEND_URL="https://your-frontend-domain.com"
```

### Health Checks

The application provides health check endpoints:
- `GET /v1/health` - Basic health status
- `GET /debug/vars` - Detailed metrics


## 📝 License

This project is licensed under the Apache 2.0 License - see the [LICENSE](LICENSE) file for details.

## 🆘 Troubleshooting

### Common Issues

1. **Database Connection Failed**
   ```bash
   # Check if PostgreSQL is running
   pg_isready -h localhost -p 5432
   
   # Verify connection string in .envrc
   ```

2. **Redis Connection Failed**
   ```bash
   # Check if Redis is running
   redis-cli ping
   
   # Disable Redis if not needed
   export REDIS_ENABLED=false
   ```

3. **Migration Errors**
   ```bash
   # Check migration status
   migrate -path ./cmd/migrate/migrations/ -database $DB_ADDR version
   
   # Force to specific version if needed
   make migrate-force ARG=1
   ```

4. **Port Already in Use**
   ```bash
   # Change port in .envrc
   export ADDR=":8083"
   ```

## 📞 Support

For support and questions:
- Create an issue on GitHub
- Check existing documentation
- Review the API documentation at `/swagger/index.html`

---

**Happy Coding! 🎉**
