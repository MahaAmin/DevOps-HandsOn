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
9. [Docker Compose](#docker-compose)
10. [Sample Application - Voting Application](#sample-application-voting-application)
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

- Delete container:
```
docker rm <CID/CName>
docker rm silly_sammet
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

## Docker Compose

- If you want to run an complex application through multiple containers, for example:
    - app
    - mongodb
    - reddis
    - ansible
- ... It is recommended to use docker compose.

- docker-compose.yml:
```
services:
    web:
        image: "mmumshad/simple-webapp"
    database:
        image: "mongodb"
    messaging:
        image: "redis:alpine"
    orchestration:
        image: "ansible"
```

- Run docker-compose file:
```
docker-compose up
```

<hr>

## Sample Application: Voting Application:

- Architecture:
    - voting-app (python)
    - In-memory db (redis)
    - worker (.Net)
    - db (PostgresSQL)
    - result-app (nodeJS)

- Create application stack manually:
```
docker run -d --name=redis redis
```
```
docker run -d --name=db postgress
```

```
docker run -d --name=vote -p 5000:80 --link redis:redis voting-app
```

```
docker run -d --name=result -p 5001:80 --link db:db result-app
```

```
docker run --name=worker --link db:db --link redis:redis worker
```

- Transform the upper commands into docker-compose file **version 1** :
```
redis:
    image: redis

db:
    image: postgres:9.4

vote:
    image: voting-app
    # if the image is not yet built, and you want to build it from Dockerfile, change it to: build: ./vote-dir

    ports:
        - 5000:80
    links:
        - redis

result:
    image: result-app   # build: ./result-dir
    ports:
        - 5001:80
    links:
        - db

worker:
    image: worker   # build: ./worker-dir
    links:
        - db
        - redis
```


- Docker-compose versions :
    1. Version 1: above example
    2. Version 2:
        - No links sections, replaced by depends_on
        ```
        version: 2
        services:
            redis: 
                ...
            vote: 
                ...
                depends_on:
                    - redis
        ```
    3. Version 3:
        - Supports docker-swarm.
        - No need for links nither depends_on setions. It creates network where containers created from the same docker-compose file can talk to eachother.

        ```
        version: 3
        services:
            redis:
                ...
            db:
                ...
            vote:
                imgae: voting-app
                ports:
                    - 5000:80
        ```

- Create networks and attach containers to them:
    ```
    version: 2
    services:
        redis:
            image: redis
            networks:
                - back-end
        
        db:
            image: postgress:9.4
            networks:
                - back-end
        
        vote:
            image: voting-app
            networks:
                - front-end
                - back-end
        
        result: 
            image: result
            networks:
                - front-end
                - back-end
    
    networks:
        front-end:
        back-end:
    ```

<hr>