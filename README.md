# STMT Core backend

Golang serverless backend for STMT Application

## Requirements

- [ ] Swagger
- [ ] Build script & Makefile
- [ ] Serverless framework setup
- [ ] Service Logic
 
## API

### GET
- User info by appId & email (Login)
- User's latest tasks (appId & date (sort))
- Get Todo list & Task info by Task ID

### POST
- Add a friend with email
- Add today's task (appId, date)
- Add task todo (taskId)

### DELETE
- Remove Friend with appId
- Remove Account (then, remove all records & tasks & todos & friendlist)

## AWS DynamoDB Table
- PK SK Inverted GSI 
- Score LSI (Sparse Key)
 
|PK|SK|nickname|profile|todo|score|memo|where|willStart|estimatedMinutes|completedAt|createdAt|
|:-----:|:-----:|:-----:|:-----:|:-----:|:-----:|:-----:|:-----:|:-----:|:-----:|:-----:|:-----:|
|USER#appId|PROFILE#email|Nickname|imgUrl|-|-|-|-|-|-|-|-|-|-|
|USER#appId|REC#date|-|-|-|999|-|-|-|-|-|-|
|TASK#uuid|REC#date|-|-|[]|-|blah|hanyang univ|1604343297363|300|1604343441719|1604343257363|
|TASK#uuid|REC#date|-|-|[{content, isCompleted}]|-|-|-|-|-|-|1604343277363|