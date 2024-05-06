go test -race -timeout 120s -count=1 -coverprofile=c.out $(go list ./... | grep -vE 'auth-service/(mocks|internal/gen|internal/models|internal/config|internal/dependencies)')
