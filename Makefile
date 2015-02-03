build: dep
	mkdir -p bin
	go build -o bin/webapi ./webapi
	go build -o bin/publisher ./publisher
	go build -o bin/dispatcher ./dispatcher

server: dep
	go run webapi/*.go

publisher: dep
	go run publisher/*.go

dispatcher: dep
	go run dispatcher/*.go

dep:
	go get github.com/tools/godep
	go get github.com/streadway/amqp
	go get github.com/go-martini/martini
	go get github.com/nu7hatch/gouuid

docker: dep build
	

rabbit:
	docker pull dockerfile/rabbitmq
	docker run -d -e RABBITMQ_NODENAME=smsgs_rabbit --name rabbit -p 127.0.0.1:5672:5672 -p 127.0.0.1:15672:15672 rabbitmq:3-management

mongo:
	docker pull dockerfile/mongodb
	docker run -d -p 27017:27017 --name mongo dockerfile/mongodb
