module endpoint

go 1.13

replace internal/messagequeue => ../messagequeue/

replace internal/db => ../db/

replace internal/db/entities => ../db/entities

require (
	github.com/gin-gonic/gin v1.7.4
	github.com/lib/pq v1.10.3
	internal/db/entities v0.0.0
	internal/messagequeue v0.0.0
)
