# CodelyTV - Go HTTP API - Hexagonal Architecture

## References (URLs)
- Gin: https://gin-gonic.com/
- Standard Go Project Layout: https://github.com/golang-standards/project-layout
- Cómo estructurar tus proyectos en Go: https://blog.friendsofgo.tech/posts/como_estructurar_tus_aplicaciones_go/
- package-oriented-design: https://www.ardanlabs.com/blog/2017/02/package-oriented-design.html
- Mockery: https://github.com/vektra/mockery
- Testify: https://pkg.go.dev/github.com/stretchr/testify
- Testify mock - AssertExpectations: https://pkg.go.dev/github.com/stretchr/testify@v1.7.0/mock#Mock.AssertExpectations
- go-sqlbuilder: https://github.com/huandu/go-sqlbuilder
- go-sqlmock: https://github.com/DATA-DOG/go-sqlmock
- faker: https://github.com/dmgk/faker
- GoDDD: https://github.com/marcusolsson/goddd
- context package: https://golang.org/pkg/context/
- context article: https://blog.golang.org/context
- context en go: https://blog.friendsofgo.tech/posts/context-en-golang/
- patrones de concurrencia (context): https://blog.friendsofgo.tech/posts/patrones-de-concurrencia-context/

## Commands used for create the project
- Initialize go mod: `go mod init github.com/jorgechavezrnd/go-hexagonal_http_api`
- Get dependencies: `go mod tidy`

## Commands for install some dependencies
- Mockery: `go get github.com/vektra/mockery/v2/.../`

## Commands for generate mocks
- `cd internal`
- Generate course_repository: `mockery --case=snake --outpkg=storagemocks --output=platform/storage/storagemocks --name CourseRepository`
- `cd kit/command`
- Generate command bus: `mockery --case=snake --outpkg=commandmocks --output=commandmocks --name=Bus`

## Command for run all tests
- `go test ./...`

## Command for start HTTP API Server
- `go run cmd/api/main.go`
