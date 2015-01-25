server: dep
	godep go run subapi/*.go

publisher: dep
	godep go run test/publisher.go

subscriber: dep
	godep go run test/subscriber.go

dep:
	go get github.com/tools/godep

docker:
	docker build -t smsgs_rabbit .

rabbit:
	docker run -d -e RABBITMQ_NODENAME=smsgs_rabbit --name smsgs_rabbit -p 15672:15672 smsgs_rabbit
