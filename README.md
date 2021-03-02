# Dockerfiles

# Requirements 

- Docker 
- Golang

# Build images

```sh
docker build -f dockerfiles/Dockerfile_1 -t dockerfile_1 .
docker build -f dockerfiles/Dockerfile_2 -t dockerfile_2 .
docker build -f dockerfiles/Dockerfile_3 -t dockerfile_3 .
```

# Run Container

```sh
docker run --rm -p 3000:3000 dockerfile_1
docker run --rm -p 3000:3000 dockerfile_2
docker run --rm -p 3000:3000 dockerfile_3
```
