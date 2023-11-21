
# Dogsuli
* Go colly crawler
* Docker container


## Starting proyect with go

```
go mod init dogsuli
go mod tidy
```

* Running the project to test:

```
go run dogsuli.go main.go

```
## TOR Proxy

```
sudo docker run -it -p 8118:8118 -p 9050:9050 -d dperson/torproxy
```
* Or using docker-compose

```
  tor1:
    container_name: tor1
    image: dperson/torproxy
    ports:
      - 8118:8118
      - 9051:9050
```

```
docker-compose up -d
```

* To check

```
$ curl -x socks5h://localhost:9050 ipinfo.io
$ docker exec -it tor1 curl ipinfo.io
```

### Torproxy proyect

https://github.com/dperson/torproxy


## Creating Docker image

```
docker build -t dogsuli .
docker run dogsuli
```