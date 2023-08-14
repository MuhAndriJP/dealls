dealls

# Usage
```
- git clone https://github.com/MuhAndriJP/dealls
- cd dealls
- cp .env.example .env
```

# How To Run

# Without Docker
```
- go run main.go
```

# With Docker
```
- make compose-up
```
_or_
```
- make build
- make run
```

# Tree
```
.
├── Dockerfile
├── action
│   └── user
│       ├── login_user.go
│       ├── login_user_test.go
│       ├── mocks
│       │   └── IUser.go
│       ├── register_user.go
│       └── register_user_test.go
├── docker-compose.yml
├── entity
│   └── user.go
├── go.mod
├── go.sum
├── handler
│   └── user
│       ├── login.go
│       └── register.go
├── helper
│   ├── helper.go
│   └── response.go
├── main.go
├── makefile
├── middleware
│   ├── jwt_middlewares.go
│   └── log_middlewares.go
├── readme.md
├── repo
│   └── mysql
│       ├── init.go
│       └── user.go
└── routes
    └── routes.go
```