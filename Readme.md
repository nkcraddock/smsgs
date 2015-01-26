# smessages for your smessaging needs

# start the rabbit container
this will start the rabbit docker container on localhost:5672 and localhost:15672

	$ sudo make rabbit 

# server

    subapi/main.go - web api w/ go martini
  	  /webhooks - subscriptions pointing topic filters to your url endpoint
      /events - POST an event to dump it into the bus
      /mock-subscriber - This is the API endpoint the dispatcher pushes to

# dispatcher
	test/dispatcher.go - pulls messages off the queues and delivers them to the appropriate endpoint 

# publisher
	test/publisher.go - publishes random events to the main exchange

# to test
run all of these:
~~~
  $ make server
  $ make dispatcher
  $ make publisher
~~~

then use your favorite api client tool to POST
~~~
POST /webhooks HTTP/1.1
Host: localhost:3001

{ "url": "http://localhost:3001/test", "pub": "qt", "typ": "*", "key": "B" }
~~~
	
