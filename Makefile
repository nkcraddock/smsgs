run: dep
	godep go run server.go
dep:
	go get github.com/tools/godep

docker:
	docker build -t smsgs .
	docker run -d -e RABBITMQ_NODENAME=smsgs --name smsgs smsgs

