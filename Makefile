server: dep
	godep go run subapi/*.go

publisher: dep
	godep go run test/publisher.go

dispatcher: dep
	godep go run test/dispatcher.go

dep:
	go get github.com/tools/godep

rabbit:
	docker run -d -e RABBITMQ_NODENAME=smsgs_rabbit --name smsgs_rabbit -p 127.0.0.1:5672:5672 -p 127.0.0.1:15672:15672 rabbitmq:3-management
