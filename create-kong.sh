#!/bin/bash
# This file should be used only outside the running container (on host).

API=users
SERVICE_ADDR=$API-service:8070

ADMIN_HOSTNAME=localhost
ADMIN_PORT=8050
ADMIN_ADDR=$ADMIN_HOSTNAME:$ADMIN_PORT

# curl -i -X POST http://kong:8001/services/ --header 'Content-Type: application/json' --data '{"name": "auth-service", "host": "auth-service", "port": 8070, "protocol": "http",
# "path": "/auth"}'
# Create Service
curl -i -X POST "http://$ADMIN_ADDR/services/" --header 'Content-Type: application/json' \
     --data '{ "name": "'$API'-service", "host": "'$API'-service", "port": '$ADMIN_PORT', "protocol": "http", "path": "/" }'
# curl -i -X POST http://kong:8001/services/auth-service/routes --header 'Content-Type: application/json' --data '{"paths": ["/auth"], "strip_path": true}'
# Create Route to service
curl -i -X POST "http://$ADMIN_ADDR/services/$API-service/routes" --header 'Content-Type: application/json' \
     --data '{ "methods": ["GET", "POST"] }'
# curl -i -X POST http://kong:8001/upstreams --header 'Content-Type: application/json' --data '{"name": "auth-service"}'
# Add UPSTREAM
curl -i -X POST "http://$ADMIN_ADDR/upstreams/" --header 'Content-Type: application/json'  \
     --data '{ "name": "'$API'-service" }'
# curl -i -X POST http://kong:8001/upstreams/auth-service/targets --header 'Content-Type: application/json' --data '{"target": "auth-service:8070"}'
# Add target to upstream
curl -i -X POST "http://$ADMIN_ADDR/upstreams/${API}-service/targets" --header 'Content-Type: application/json'  \
     --data '{ "target": "'$SERVICE_ADDR'" }'
