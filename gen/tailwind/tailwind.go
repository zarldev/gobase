package tailwind

//go:generate echo "Running tailwindcss build..."
//go:generate npx tailwindcss -i styles.css -o ../../ui/static/css/styles.css
//go:generate echo "Done running tailwindcss build command."
