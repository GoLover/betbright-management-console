# 4. Use ElasticSearch as fulltext search database


## Status

Accepted

## Context
 handling search on main database will increase cost and reduce performance of database.
at the other hand features which achieved easily by fulltext search databases are harder to achieve in other type of databases.
so choosing a fulltext search specific database is an option for scale and ease of handling search queries.
## Decision
use ElasticSearch as search engine.
## Consequences
synchronization between postgres and elasticsearch is a challenge.
consistency between main database and elasticsearch will be done eventually.
searching will be easier and scaling is much easier to achieve.