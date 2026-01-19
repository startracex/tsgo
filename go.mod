module github.com/startracex/dev-server

go 1.25

replace github.com/microsoft/typescript-go => ./pkg/tsgo

require github.com/startracex/go-pub v0.0.0-20251012180452-cce541aafcd7

require golang.org/x/mod v0.32.0 // indirect

ignore ./submodules
