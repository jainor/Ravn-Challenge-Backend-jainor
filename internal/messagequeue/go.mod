module messagequeue

go 1.13

replace internal/db => ../db/

replace internal/db/entities => ../db/entities

require (
	github.com/lib/pq v1.10.3
	github.com/rabbitmq/amqp091-go v0.0.0-20210823000215-c428a6150891
	internal/db v0.0.0
	internal/db/entities v0.0.0
)
