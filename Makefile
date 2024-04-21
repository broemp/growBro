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
	@go build -o bin/GrowBro .

templ:
	@templ generate -watch -proxy=http://localhost:3000

tailwind:
	@tailwindcss -i view/css/input.css -o public/styles.css --watch

db:
	docker rm -f cannaBro_Postgres && docker run -p 127.0.0.1:5433:5432 -e POSTGRES_USER=cannaBro -e POSTGRES_PASSWORD=cannaBro -e POSTGRES_DB=cannaBro --name cannaBro_Postgres -d postgres
