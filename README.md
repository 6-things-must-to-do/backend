# STMT Core backend

Golang serverless backend for STMT Application

## Requirements

- [ ] Swagger
- [ ] Build script & Makefile
- [ ] Serverless framework setup
- [ ] Service Logic
 
## API

### GET
- [x] User info by uuid (My page, Get user)
- User's latest tasks (appId & date (sort))
- Get Todo list & Task info by Task ID

### POST
- [x] Issue JWT by appId, provider, email
- [x] Create user if don't exist 
- Add a friend with email
- Add today's task (appId, date)
- Add task todo (taskId)

### DELETE
- Remove Friend with appId
- Remove Account (then, remove all records & tasks & todos & friendlist)

### PUT
- Complete Todo by task id
- Complete Task by task id

## AWS DynamoDB Table
- PK SK Inverted GSI 
- Score LSI (Sparse Key)
- AppID, SK GSI (H: SK, SK: AppID)
 
|PK|SK|AppID|nickname|profile|todo|score|memo|where|willStart|estimatedMinutes|completedAt|createdAt|
|:-----:|:-----:|:-----:|:-----:|:-----:|:-----:|:-----:|:-----:|:-----:|:-----:|:-----:|:-----:|:-----:|
|USER#uuid|PROFILE#email|Hashed AppID|Nickname|imgUrl|-|-|-|-|-|-|-|-|-|-|
|USER#uuid|REC#appId#date|-|-|-|-|999|-|-|-|-|-|-|
|TASK#uuid|REC#appId#date|-|-|-|[]|-|blah|hanyang univ|1604343297363|300|1604343441719|1604343257363|
|TASK#uuid|REC#appId#date|-|-|-|[{content, isCompleted}]|-|-|-|-|-|-|1604343277363|