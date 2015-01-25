# smessages for your smessaging needs
~~~
To start the rabbit container:
  $ sudo make docker rabbit 
~~~

# server
	/webhooks - subscriptions pointing topic filters to your url endpoint

# publisher
	test/publisher.go - publishes random events to the main exchange

# subscriber
	test/subscriber.go - reads a specific subscriber queue and prints the messages to stdout

	
