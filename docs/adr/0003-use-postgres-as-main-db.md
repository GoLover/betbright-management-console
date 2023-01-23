# 2. Use PostgreSQL for main DB


## Status

Accepted

## Context
 using a relational database make us sure about relation between data and handling actions on them.
we can relate data to each other and handle consistency between them.
at the other hand, PostgreSQL has a good indexing and searching options, so it make our job easier.
## Decision
use PostgreSQL as an object relational database
## Consequences
It will depend project on PostgreSQL. data consistency will be outsourced to database instead handling in service itself.
unique constraint will achieve easily and deduplication will happen by nature of DB.
