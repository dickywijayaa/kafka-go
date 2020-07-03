# kafka-go
Learn creating consumer and producer using kafka in golang

Steps :
1. Make sure you already have java in your local machine!
2. After that download kafka from https://kafka.apache.org/
3. Run through tutorial from https://kafka.apache.org/quickstart until steps 2 are done
4. If you already done until step 2, then you are ready to run the consumer and producer
5. Run consumer first by execute : `go run consumer/main.go`
6. After that run producer by execute : `go run consumer/main.go`
7. Check the consumer terminal there will be messages / events like these :
```
Message on notification-topic[0]@1: Hello
Message on notification-topic[0]@2: World
Message on notification-topic[0]@3: My
Message on notification-topic[0]@4: Name
Message on notification-topic[0]@5: Dicky
```

That's it! As simple as that you successfully run a simple producer and consumer of kafka in golang.
I would do more code later in this repository.