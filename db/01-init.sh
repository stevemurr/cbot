#!/bin/bash
set -e
export PGPASSWORD=$POSTGRES_PASSWORD;
psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL
  GRANT ALL PRIVILEGES ON DATABASE $POSTGRES_DB TO $POSTGRES_USER;
  \connect $POSTGRES_DB $POSTGRES_USER
  BEGIN;
    CREATE TABLE IF NOT EXISTS prices(
      price TEXT NOT NULL,
      time date NOT NULL,
      productid TEXT
	);
  COMMIT;
EOSQL