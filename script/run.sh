# !/bin/bash

# stop previous container
docker rm -f youtube-data-app
docker rm -f youtube-mongo-db

# build new image
docker-compose -f docker/docker-compose.yml build

# run the project
docker-compose -f docker/docker-compose.yml up