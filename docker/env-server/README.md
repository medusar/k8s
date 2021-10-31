# env-server
A simple http server written in golang mainly used for learning Kubernetes.
When a request is received, it will response with current 'HOSTNAME' env value and started time.

## Build image and push to docker hub
```
//build it
docker build -t env-server .
//tag it
docker image tag env-server medusar/env-server:v1.0
//login to dockerhub
docker login
//push to dockerhub
docker image push medusar/env-server:v1.0
```

## Reference
- [how to build a golang docker image](https://docs.docker.com/language/golang/build-images/)
- [how to push a docker image to docker hub](https://www.techrepublic.com/article/how-to-build-a-docker-image-and-upload-it-to-docker-hub/)