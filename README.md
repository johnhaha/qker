# qker

qker is a dumb go-quic based pkg

## Install

```bash
go get github.com/johnhaha/qker@v0.0.2
```

## server

```go
ctx, cancel := context.WithCancel(context.Background())
defer cancel()
s := qker.NewServer("127.0.0.1:4567")
s.SetHandler(func(c qker.Ctx) error {
    msg := c.String()
    log.Println("data comes", msg, "remote addr is", c.RemoteAddr().String())
    return nil
})
log.Fatal(s.StartServer(ctx))
```

## client

```go
ctx, cancel := context.WithCancel(context.Background())
defer cancel()
c := qker.NewClient("127.0.0.1:4567")
err := c.Dial()
if err != nil {
    panic(err)
}
c.Send("foo")
<-ctx.Done()
```
