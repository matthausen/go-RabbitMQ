An exampe of how to use RabbitMQ messaging between a client and a consumer.


## Run RabbitMQ
Run an instance of RabbitMQ in docker:

`docker run -d --hostname my-rabbit --name some-rabbit -p 15672:15672 -p 5672:5672 rabbitmq:3-management`

Then visit `http://localhost:15672/#` 
username: guest password: guest


## How to run:

Run the consumer:
`go run consumer/consumer.go`

Run the client:
`go run client/client.go`