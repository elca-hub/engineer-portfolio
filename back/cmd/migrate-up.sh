#!/bin/bash

source ./.env

migrate -database "mysql://$MYSQL_USER:$MYSQL_PASSWORD@tcp(localhost:$MYSQL_PORT)/$MYSQL_DATABASE" -path ./migrations up