<!-- Code generated by gomarkdoc. DO NOT EDIT -->

# messagequeue

```go
import "messagequeue"
```

## Index

- [func ResponseMessage(m idb.ModelQ)](<#func-responsemessage>)
- [func SendMessage(n int) ([]byte, error)](<#func-sendmessage>)
- [type RabbitCredentials](<#type-rabbitcredentials>)


## func ResponseMessage

```go
func ResponseMessage(m idb.ModelQ)
```

ResponseMessage responds the messages waiting

## func SendMessage

```go
func SendMessage(n int) ([]byte, error)
```

SendMessage sends a message to Rabbitmq server and returns a result via RPC

## type RabbitCredentials

RabbitCredentials stores Rabbitmq credentials for connections

```go
type RabbitCredentials struct {
    Host     string
    Port     string
    User     string
    Password string
    Name     string
}
```



Generated by [gomarkdoc](<https://github.com/princjef/gomarkdoc>)