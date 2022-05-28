# Johnny's unoriginal (not so) Smart Switch 3.0

After many iterations of using raw html/css/javascript to using React, and jumping from PHP to Node, and finally to Go on the Backend side, this is the definitive version.  
This is a simple app that will scan my home network for smart devices, to power them on and off.  
It will additionally serve as a gateway to send a Wake-On-Lan signal to my PC.

## How to set this up on a low-powered Pi device?

Pre-requisites: Docker

1. Setup `.env` according to the template
2. Build the necessary images according to the comments in `prod_pi.sh`. The command might end up looking like: `ARCH=linux/arm/v7 ./prod_pi.sh`
3. Execute `docker compose -f docker-compose-pi.yml up -d`
