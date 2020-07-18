# Simple gin golang app
Simple golang app get date trivia from http://numbersapi.com using gin-gonic framework. this app using default configuration so thats run on `http://0.0.0.0:8080`.
this consist only 1 endpoint `0.0.0.0:8080/trivia`

## build docker container 
Build docker container onli run `build.sh` with parameter `<STATE_ENV>` and `<tag version>`
- `<STATE_ENV>` is env of application like production, staging and development or whatever that developer want to sparate the server environment. this parameter will save to `os env` on docker image. so the backend will read this `os env` to sparate the configuration.

- `<tag version>` is versioning of application
example :
    ``` sh
    sh build.sh development v1.0
    ```
- some detail is written on script `build.sh` and `Dockerfile`

step to change docker image and push to docker registry make some change on `build.sh`
1. change `REGISTRY` parameter as your private registry on 
    ex : `REGISTRY: "asia.gcp.com/private-registry/"`
2. change `APP_NAME` parameter as your docker image name
    ex: `APP_NAME: "simple_gin"`
3 uncommand this line:
    `#docker push ${REGISTRY}${APP_NAME}:${REVISION}`
