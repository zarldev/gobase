package main

// openapi client and server generation
//go:generate echo "OpenAPI Code Generation"
//go:generate oapi-codegen --config=./tools/openapi/config.yaml ./tools/openapi/api.yaml

// templ UI template generation
//go:generate echo "Templ HTML Generation"
//go:generate templ generate -path ui

// Tailwind CSS css generation
//go:generate echo "Tailwind CSS Generation"
//go:generate npx tailwindcss -c tools/tailwind/tailwind.config.js -i tools/tailwind/styles.css -o ui/static/css/styles.css

// sqlc go code generation
//go:generate echo "SQLC Code Generation"
//go:generate sqlc generate -f ./tools/sqlc/sqlc.yaml

//go:generate echo "Completed"
