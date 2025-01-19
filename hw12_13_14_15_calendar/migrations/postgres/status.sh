#!/bin/bash

LOCAL_MIGRATION_DIR=./
if [ "$PG_MIGRATIONS_DIR" != "" ]; then
    LOCAL_MIGRATION_DIR=$PG_MIGRATIONS_DIR
fi
echo "Using migrations dir: $LOCAL_MIGRATION_DIR"

LOCAL_MIGRATION_DSN="postgres://${PG_USER}:${PG_PASSWORD}@${PG_HOST_LOCAL}:${PG_PORT_LOCAL}/${PG_DATABASE}?sslmode=disable"

goose -dir ${LOCAL_MIGRATION_DIR} postgres ${LOCAL_MIGRATION_DSN} status -v