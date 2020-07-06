# Studying Docker

## Table of Content:
1. [Docker Commands](#docker-commands)
2. [More on Docker Run](#more-on-docker-run)
3. [Example: Run NGINX](#example-run-nginx)
4. [Container Environment Variables](#container-environment-variables)
5. [Docker Images](#docker-images)
6. [CMD VS ENTRYPOINT](#cmd-vs-entrypoint)
7. [Docker Networking](#docker-networking)
8. [Docker Storage](#docker-storage)
<hr>


## Docker Commands:
- Pull the image from dockerhub -if image does not already exist- and run a container instance from the image:

```
docker run <docker/whalesay>    
```

- List all **running container**:
```
docker ps
```

- List all containers:
```
docker ps -a
```

- Stop running container:
```
docker stop <CID/CName>
docker stop silly_sammet
```

- Stop all running container:
```
docker stop $(docker ps -q)
```

- Delete container:
```
docker rm <CID/CName>
docker rm silly_sammet
```

- Delete all container:
```
docker rm $(docker ps -aq)
```

- List existing images:
```
docker images
```

- Delete image:
**You must delete all containers depending on that image before deleting the image**
```
docker rmi <IID/IName>
```

- Delete all images:
```
docker rmi $(docker images -q)
```


- To know the base image of a container:
```
docker run python:3.6 cat /etc/*release*
```

- Download an image without running a container:
```
docker pull <Image>
docker pull nginx
```

- Execute a command on a running container:
```
docker exec <CID/CName> <command>
docker exec silly_sammet cat /etc/hosts
```

- Run an instance container from an image in a **detached** mode:
```
docker run -d <docker/whalesay>
```

- Reattach deattaced container:
```
docker attach <CID>
```

- Inspect container (Get full information about a container):
```
docker inspect <CID/CName>
```

- Show container logs:
```
docker logs <CID/CName>
```

- Delete all docker resources (images/container/networks/ ...etc)
```
docker system prune
```

<hr>

## More on Docker Run:
- Run specific version of an image (use tag):
**default tag "latest" redis:latest**
```
docker run redis:4.0
```

- RUN - STDIN:
    - i --> interactive. Accept input data from user.
    - t --> terminal. Show container terminal.

```
docker run -it kodekloud/simple-prompt-docker
```

- RUN - Port Mapping:
```
docker run -p hostPort:containerPort image
docker run -p 80:80 nginx
```

- RUN - Volume Mapping:
```
docker run -v hostDirPath:containerDirPath image
docker run -v /opt/mysql:/var/lib/mysql mysql
```

- RUN - Environment Variables:
```
docker run -r APP_COLOR=blue image
```

<hr>

## Example: Run NGINX

```
docker run -d nginx
```

- To get the internal ip of the container:
```
docker inspect hardcore_grothendieck
```

- From the browser:
```
http://172.17.0.2:80
```

<hr>

## Container Environment Variables:

- RUN - Environment Variables:
```
docker run -r APP_COLOR=blue image
```

- Inspect container's environment variables:
```
docker inspect <CID/CName> 
```
--> "Config" --> "Env"

- You can find all environment varaibles of an image in the description part on dockerhub. 


<hr>

## Docker Images

**How to create your own image?**
- OS - Ubuntu
- Update apt repo
- Install dependencies using apt
- Install python dependencies using pip
- Copy source code to /opt folder
- Run web server using "flask" command

**Dockerfile:**
```
FROM Ubuntu

RUN apt-get update
RUN apt-get install python

RUN pip install flask
RUN pip install flask-mysql

COPY . /opt/source-code

ENTRYPOINT FLASK_APP=/opt/source-code/app.py flask run
```

**Build Image:**
```
docker build . -t userName/my-app
```

**Push to docker hub registry:**
```
docker push userName/my-app
```

<hr>

## CMD VS ENTRYPOINT

```
docker run ubuntu sleep 1000
```

```
FROM Ubuntu 

CMD ["sleep", "5"]
```

```
FROM Ubuntu
ENTRYPOINT ["sleep"]
```
--> docker run ubuntu 1000

```
FROM Ubuntu
ENTRYPOINT ["sleep"]
CMD ["5"] // default value can be overwritten
```

- docker run ubuntu // sleep 5
- docker run ubuntu 1000 //sleep 1000
- docker run --entrypoint "sleep 900" ubuntu

<hr>

## Docker Networking

**Default networks:**
- Bridge (Default):
    - docker0: 172.17.0.1
    - Provides network isolation between containers.
- None:
    - The container has no network.
    - Totally isolated.
- Host:
    - Uses same host network.
    - No need for port mapping.
    - Only one container can use the same port number.

**User-defined networks:**
- Create new network inside docker:
```
docker network create \
    --driver bridge \
    --subnet 182.18.0.0/16 \
    custom-isolated-network-name
```

- List existing networks:
```
docker netwrk ls
```

- Inspect a network:
```
docker network inspect <bridge>
```

- Inspect container network:
```
docker inspect <CID/CName>
```
--> "Network" --> "Bridge"

- Embedded DNS:
```
mysql.connect( containerName );
```

<hr>

## Docker Storage
- Docker filesystem:
    - /var/lib/docker
        - /var/lib/docker/containers
        - /var/lib/docker/volumes
        
- Layered architecture.
<hr>

