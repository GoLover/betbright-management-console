# BetBright Challenge - 888Spectate

BetBright-Management-Console is a code challenge task for 888Spectate Company.

Architecture Decision Records are available under `docs` directory.

## How to build
there is a Dockerfile and you can build a ready image.
but if you want to build it outside docker you can go through these steps.
```
go get
go build
```

## How to start 
```
docker-compose up -d
./betbright-management-console [entity]
```