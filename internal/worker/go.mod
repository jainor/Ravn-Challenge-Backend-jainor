module worker

go 1.13

replace internal/messagequeue => ../messagequeue/

replace internal/db => ../db/

replace internal/db/entities => ../db/entities/

require (
	internal/db v0.0.0
	internal/messagequeue v0.0.0
)
