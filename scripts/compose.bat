@echo off
docker compose -f ./deploy/compose.yaml -p quotes_api up -d
