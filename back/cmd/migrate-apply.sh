#!/bin/bash

source ./.env

atlas migrate apply \
--url "mysql://$MYSQL_USER:$MYSQL_PASSWORD@$YSQL_HOST:$MYSQL_PORT/$MYSQL_DATABASE" --dir "file://migrations" --baseline 20250301132313
