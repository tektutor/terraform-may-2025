module main

go 1.24.2

replace addition => ./addition

replace subtraction => ./subtraction

require (
	addition v0.0.0-00010101000000-000000000000
	subtraction v0.0.0-00010101000000-000000000000
)
