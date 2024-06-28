package tools

// openapi client and server generation
//go:generate oapi-codegen --config=openapi/config.yaml openapi/api.yaml
// templ UI template generation
//go:generate templ generate -path ../ui
// Tailwind CSS css generation
//go:generate npx tailwindcss -i ./tailwind/styles.css -o ../ui/static/css/styles.css
// sqlc go code generation
//go:generate sqlc generate -f ./sqlc/sqlc.yaml
