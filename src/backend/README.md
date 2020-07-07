# Athenaeum Backend

![Backend Build Status](https://github.com/raygervais/Athenaeum/workflows/Go/badge.svg)

## Onboarding and Setup

- Stack: Golang 1.14
- Dependencies: See `go.mod`
  - Install with: `go mod download`
- Use (Runtime): `go run .`
- Use (Binary): `go build -o main cmd/api/main.go && ./main`
- Use (Container): `docker build . -t athena-backend && docker run -p 3000:3000 athena-backend`

    
### Setup

```bash
go mod download
go run .
```

