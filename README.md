# Сервис на основе NATS Streaming


## Запуск сервиса

```shell
go build .\main.go .\http_handler.go .\model.go .\nats-subscriber.go .\CLI.go
.\main.exe --new_only -c test-cluster -id myID foo
```

### Параметры запуска
```
Usage: main.exe [options] <subject>

Options:
	-s,  --server   <url>            NATS Streaming server URL(s)
	-c,  --cluster  <cluster name>   NATS Streaming cluster name
	-id, --clientid <client ID>      NATS Streaming client ID
	-cr, --creds    <credentials>    NATS 2.0 Credentials

Subscription Options:
	--qgroup <name>                  Queue group
	--all                            Deliver all available messages
	--last                           Deliver starting with last published message
	--since  <time_ago>              Deliver messages in last interval (e.g. 1s, 1hr)
	--seq    <seqno>                 Start at seqno
	--new_only                       Only deliver new messages
	--durable <name>                 Durable subscriber name
	--unsub                          Unsubscribe the durable on exit
```
## Тестовая отправка JSON по каналу

```shell
go run .\test\pub.go foo
```
