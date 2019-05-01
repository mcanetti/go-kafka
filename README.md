# set producer topic
edit go-kafka/main.go. Change topic variable to the desired topic name.

# start producer
```bash
go run main.go
```

# generate messages:
```bash
http http://localhost:8080
```

# set consumer topic
edit go-kafka/consumer/main.go. Change topic variable to the desired topic name.

# start consumer
```bash
go run consumer/main.go
```
the consumer should just collect all messages for the topic and shut down in 10 seconds. Increase life time editing go-kafka/consumer/main.go
