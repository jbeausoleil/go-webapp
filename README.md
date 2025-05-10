**Directory Structure**
```
/go-webapp/
│
├── cmd/                # Entry point to your app (main.go)
│   └── web/
│       └── main.go
│
├── internal/           # Application logic that's private to this project
│   ├── handlers/       # HTTP route handlers (controller layer)
│   ├── models/         # Database models and access (data layer)
│   ├── middleware/     # Reusable HTTP middleware
│   ├── config/         # App configuration structs and loading
│   ├── templates/      # HTML templates
│   ├── auth/           # Auth logic
│   └── jobs/           # Channels and background jobs
│
├── web/static/         # CSS/JS/images (embedded with Go 1.16+ embed)
├── migrations/         # SQL migration files
├── test/               # Optional: integration tests, mocks, fixtures
│
├── go.mod
└── .env
```