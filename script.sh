docker image rm app-journey:v1 --force

docker buildx build \
    -t app-journey:v1 . \
    --no-cache


docker container rm app-journey:v1 --force

docker container run app-journey:v1 \
    --name app-journey \
    -p 8080:8080