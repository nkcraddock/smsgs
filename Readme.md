# smessages for your smessaging needs

# start the rabbit container
this will start the rabbit docker container on localhost:5672 and localhost:15672
  $ sudo make rabbit 

# server
	/webhooks - subscriptions pointing topic filters to your url endpoint

# publisher
	test/publisher.go - publishes random events to the main exchange

# subscriber
	test/subscriber.go - reads a specific subscriber queue and prints the messages to stdout

# to test
run all of these:
~~~
  $ make server
  $ make publisher
  $ make subscriber
~~~

then use your favorite api client tool to POST
~~~
POST /webhooks HTTP/1.1
Host: localhost:3001

{ "url": "http://localhost:3001/test", "pub": "qt", "typ": "*", "key": "B" }
~~~
	
