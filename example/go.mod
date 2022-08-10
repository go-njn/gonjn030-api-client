module example

go 1.18

replace github.com/go-njn/gonjn030-api-client => ./..

require github.com/go-njn/gonjn030-api-client v0.0.0-00010101000000-000000000000

require (
	github.com/sirupsen/logrus v1.9.0 // indirect
	golang.org/x/sys v0.0.0-20220715151400-c0bba94af5f8 // indirect
)
