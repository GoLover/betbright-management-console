# 5. Using replication feature for syncing Postgres to ElasticSearch


## Status

Accepted

## Context
 Syncing data and achieving consistency between to heterogeneous datastore is possible.
but achieving strong consistency will increase cost and not a business requirement and most scenarios.
one of solutions which works fine for postgres is logical replication in postgres.

## Decision
implementing WAL processor for use logical replication as syncing interface between postgres and elasticsearch.
## Consequences
more coding.