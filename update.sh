#!/bin/bash

docker compose build --no-cache
docker compose up -d --no-deps
docker system prune -f