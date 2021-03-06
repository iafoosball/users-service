#!/bin/bash
# This file should be used only outside the running container (on host).

API=users
API_PORT=8020
SERVICE_ADDR=$API-service:$API_PORT

ADMIN_HOSTNAME=localhost
ADMIN_PORT=8051
ADMIN_ADDR=$ADMIN_HOSTNAME:$ADMIN_PORT

# Create Service
curl -i -X POST "http://$ADMIN_ADDR/services/" --header 'Content-Type: application/json' \
     --data '{ "name": "'$API'-service", "host": "'$API'-service", "port": '$API_PORT', "protocol": "http", "path": "/" }'

# Create Route to service
curl -i -X POST "http://$ADMIN_ADDR/services/$API-service/routes" --header 'Content-Type: application/json' \
     --data '{ "paths": ["/users"], "strip_path": false  }'

# Add UPSTREAM
curl -i -X POST "http://$ADMIN_ADDR/upstreams/" --header 'Content-Type: application/json'  \
     --data '{ "name": "'$API'-service" }'

# Add target to upstream
curl -i -X POST "http://$ADMIN_ADDR/upstreams/${API}-service/targets" --header 'Content-Type: application/json'  \
     --data '{ "target": "'$SERVICE_ADDR'" }'
