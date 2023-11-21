# Docker
* If we want to use docker, it is necessary to understand some concepts:

## Image
We build an sketch based in the definitions and the operating system requirements our binary needs.

## Container
We call container when we run an image, we can run as much containerss as our system lets us from one image.

## Network
Our container behaves itself like a virtual machine with its own local network. So we should understand that it has its own localhost, not the same as the host machine.

If we want to use the port from other container launched in the same network from a docker-compose.yml, we cannot use localhost:PORT, we need to call it from its container name e.g. tor:PORT.

We have two different choices when publishing a port:

* Publishing  it to outside PORT_LOCALHOST:PORT_DOCKER_NETWORK:

```
  tor1:
    container_name: tor1
    image: dperson/torproxy
    ports:
      - 8118:8118
      - 9051:9050
```

* If I want to use it just inside the docker network. Host machine will not be able to use it.

```
  tor1:
    container_name: tor1
    image: dperson/torproxy
    ports:
      - 8118
      - 9050
```

## Steps to follow

### Option 1:

1. Build an image or pull it
2. Run a container from an image
3. Stop the container
4. Remove the container
5. If finished, remove the image


### Option 1:

1. Build an image or pull it
2. Launch docker-compose.yml


## Commands
* Create an image:

```
docker build -t dogsuli .
docker build --tag dogsuli:1.0 .
```

* Delete an image:

```
docker rmi dogsuli:latest
docker rmi dogsuli:1.0
```

* Launch a container

```
docker run -p 8080:8080 dogsuli:latest
docker run -p 8080 dogsuli:1.0
docker run --name my-container -p 8080:8080 -d your-image-name

```

* Stop a container and remove it

```
docker stop your-container-name
docker rm your-container-name
```
* See the log, the prints and logs from your app:

```
docker logs container-name
docker logs -f container-name
```

## Docker-compose
* It is a yaml file which helps you to automatize the launching of docker containers.
* It launches containers and stops them removing them and its local network too.

### Commands

```
docker-compose up
docker-compose up -d #Detached
docker-compose down
```

* If `docker-compose` command does not work, use `docker compose` instead.