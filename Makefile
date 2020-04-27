run:
	go run main.go

run_prod:
	export GIN_MODE=release && go run main.go

build_for_linux:
	GOOS=linux GOARCH=amd64 go build -o google-rtb
