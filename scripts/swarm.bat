@echo off
docker stack deploy -c ./deploy/swarm.yaml quotes_api
