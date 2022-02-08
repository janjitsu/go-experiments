test:
	gotest -v ./... -covermode=atomic

run dev:
	go run app/main.go

migrate dev:
	migrate -path /home/janjitsu/projects/go-experiments/gin-rest-api/db/migrations -database postgres://gouser:123456@localhost:5432/gin-rest-api up


