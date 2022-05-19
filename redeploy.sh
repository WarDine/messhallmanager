./stop_deploy.sh
docker-compose up -d
docker logs messhallmanager_messhallmanager-service
docker exec -it messhallmanager ip a
