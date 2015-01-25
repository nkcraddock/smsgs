# smessages for your smessaging needs

# rabbit docker container
  $ sudo make docker rabbit 

# server
	/webhooks - subscriptions pointing topic filters to your url endpoint

# publisher
	test/publisher.go - publishes random events to the main exchange

# subscriber
	test/subscriber.go - reads a specific subscriber queue and prints the messages to stdout

# to test
run all of these:
~~~
  $ make
  $ go run test/publisher.go
  $ go run test/subscriber.go
~~~

then use your favorite api client tool to POST
~~~
POST /webhooks HTTP/1.1
Host: localhost:3001

{ "url": "http://localhost:3001/test", "pub": "qt", "typ": "*", "key": "B" }
~~~
	
