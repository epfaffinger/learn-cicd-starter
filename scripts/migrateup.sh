#!/bin/bash

if [ -f .env ]; then
    source .env
fi

cd sql/schema
goose turso "libsql://notely-db-epfaffinger.aws-us-east-2.turso.io?authToken=eyJhbGciOiJFZERTQSIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE3NTk5MzU0MTQsImlkIjoiYjU0MGUwZTUtNjRkNi00M2JmLWFmZWMtMTE5MjhmYWYwYTI1IiwicmlkIjoiNzcwNDBhNzEtNmZjYi00MjA2LWIzOTYtYjIwOWJhZTgzOWRkIn0.Z4AtWE2gcosxLUuipAeD64pbnC2q-ISjdU0apqY8upBgCIufpNlDDUpFDdTsQ4fdKFp4IlfVfoquwGxO8iLQDQ" up
