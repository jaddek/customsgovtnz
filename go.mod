module github.ocm/jaddek/nzcustomsgov

replace (
	github.com/jaddek/nzcustomsgov/rate => ./rate
)

require (
	github.com/jaddek/nzcustomsgov/rate v0.0.0
)

go 1.22.4
