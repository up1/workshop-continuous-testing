#!/bin/bash
docker-compose down
docker-compose build

echo "Starting db..."
docker-compose up -d db

STATUS="starting database"
while [ "$STATUS" != "healthy" ]
do
    STATUS=$(docker inspect --format {{.State.Health.Status}} db)
    echo "Database state = $STATUS"
    sleep 5
done

echo "Starting api..."
docker-compose up -d api
echo "Finish.."