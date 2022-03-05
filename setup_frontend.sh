#!/bin/bash
docker run --rm -it -v $(pwd):/src -w /src/frontend node:16-bullseye bash -c "npm i -g pnpm && pnpm install"