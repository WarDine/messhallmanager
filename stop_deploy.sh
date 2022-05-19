#/bin/bash

docker-compose down
docker rmi messhallmanager_recipemanager-service
docker volume ls | awk '{print $2}' | xargs docker volume rm
