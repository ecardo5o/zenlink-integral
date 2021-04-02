# Setup

## ENV
```
export GOPATH="your/project/location:$GOPATH"
```

## Development
```
go run main.go
```

## Deployment
```
// compile
GOOS=linux GOARCH=amd64 go build -o ./release main.go

// kill current service if it exist
ssh [username]@[server ip]
ps aux | grep ./api-server/main
kill -9 [port number]

// upload to server from local
scp ./release/main root@[server ip]:~/api-server

// run
ssh [username]@[server ip]
export GIN_MODE=release
nohup ./api-server/main &
```

# Api
Please checkout details on Tower Doc, and test it on Postman in collection "Server Api".

- pairs.getAll