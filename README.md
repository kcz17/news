# News
A Sock Shop service that provides news/update information.

### To build this service

In order to build the project locally you need to make sure that dependencies are installed. Once that is in place you
can build by running:

```
go mod download
go build -o news
```

The result is a binary named `news`, in the current directory.

#### Docker
`docker-compose build`

### To run the service on port 8080

#### Go native

If you followed to Go build instructions, you should have a "news" binary in $GOPATH/src/github.com/kcz17/news/cmd/newssvc/.
To run it use:
```
./news
```

#### Docker
`docker-compose up`

### Check whether the service is alive
`curl http://localhost:8080/health`

### Use the service endpoints
`curl http://localhost:8080/news`

### Releasing
- `docker build -t kcz17/news:[VERSION] -f docker/news/Dockerfile .`
- `docker build -t kcz17/news-db:[VERSION] docker/news-db`
