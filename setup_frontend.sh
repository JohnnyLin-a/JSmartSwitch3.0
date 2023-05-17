#!/bin/bash
docker run --rm -it -v $(pwd):/src -w /src/frontend node:18-bullseye bash -c "yarn install"