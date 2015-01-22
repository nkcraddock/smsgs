run: dep
	godep go run server.go
dep:
	go get github.com/tools/godep
