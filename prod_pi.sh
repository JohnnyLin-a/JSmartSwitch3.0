#!/bin/bash
# A raspberry pi device is very weak, and needs to be built on external machines. Otherwise it takes over 1000seconds...
# Setup buildx and point the builder to the external server
# Then run this script.
# buildx can be set up this way: docker buildx create --name foo --platform linux/arm/v7 ssh://johnnylin-a@bar
# How to use this script: ARCH="linux/arm64" ./prod_pi.sh
docker buildx build --load --platform=$ARCH -t jsmartswitch30_frontend -f Dockerfile.node . && \
docker buildx build --load --platform=$ARCH -t jsmartswitch30_backend -f Dockerfile.golang . && \
docker buildx build --load --platform=$ARCH -t jsmartswitch30_nginx -f Dockerfile.nginx . && \
docker compose -f docker-compose-pi.yml up -d