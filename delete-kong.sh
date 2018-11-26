#!/bin/bash
# This file should be used only outside the running container (on host).

API=users
SERVICE_ADDR=$API-service:8070

ADMIN_HOSTNAME=localhost
ADMIN_PORT=8050
ADMIN_ADDR=$ADMIN_HOSTNAME:$ADMIN_PORT

# Get route id on Service
ROUTE_ID=`curl -s "http://$ADMIN_ADDR/services/$API-service/routes" | jq ".data[].id" | tr -d \" `
# Delete target on service
if [ "$ROUTE_ID" != "" ] ; then
    curl -i -X DELETE "http://$ADMIN_ADDR/routes/$ROUTE_ID"
else
    echo "Route to $API-service was not found"
fi
# Delete service
curl -i -X DELETE "http://$ADMIN_ADDR/services/$API-service"
# Delete UPSTREAM
curl -i -X DELETE "http://$ADMIN_ADDR/upstreams/$API-service"

