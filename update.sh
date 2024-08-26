#!/bin/bash

git pull
docker compose build frontend backend --no-cache
docker compose restart
