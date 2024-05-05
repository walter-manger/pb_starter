# Inspiration
# https://github.com/gobeli/pocketbase-htmx/tree/main

dev:
	npx nodemon --signal SIGTERM -e "templ go" -x "templ generate && go run main.go serve" -i "**/*_templ.go"

watch-css:
	./bin/tailwindcss -i ./src/styles/base.input.css -o ./public/styles/base.css --watch

generate-css:
	./bin/tailwindcss -i ./src/styles/base.input.css -o ./public/styles/base.css --minify

generate:
	templ generate

build: generate-css
	go build

run: generate-css
	go run main.go serve
