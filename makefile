default:
	@echo "default"

dev:
	watchexec -e go -c -r go run main.go
