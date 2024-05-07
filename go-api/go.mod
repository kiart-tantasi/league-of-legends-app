module go-api

go 1.22.2

require (
	github.com/google/uuid v1.6.0
	github.com/kiart-tantasi/env v0.0.0-00010101000000-000000000000
)

require github.com/joho/godotenv v1.5.1 // indirect

replace github.com/kiart-tantasi/env => ./pkg/env
