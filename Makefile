run: build
	@./bin/RedCard

install:
	@go install github.com/a-h/templ/cmd/templ@latest
	@go get ./...
	@go mod vendor
	@go mod tidy
	@go mod download
	@pnpm install -D tailwindcss
	@pnpm install -D daisyui@latest

build:
	tailwindcss -i view/css/input.css -o public/styles.css 
	@templ generate view
	@go build -o bin/RedCard cmd/RedCard.go 

templ:
	@templ generate -watch -proxy=http://localhost:3000

tailwind:
	@tailwindcss -i view/css/input.css -o public/styles.css --watch
