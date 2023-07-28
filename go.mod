module github.com/adriano-henrique/rss

go 1.20

replace github.com/adriano-henrique/rss/api v0.0.0 => ./api

require (
	github.com/adriano-henrique/rss/api v0.0.0
	github.com/go-chi/chi/v5 v5.0.10
	github.com/go-chi/cors v1.2.1
	github.com/joho/godotenv v1.5.1
)
