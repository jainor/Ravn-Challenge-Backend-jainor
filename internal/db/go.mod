module db

go 1.13

replace internal/db/entities => ./entities

require (
	github.com/lib/pq v1.10.3
	internal/db/entities v0.0.0
)
