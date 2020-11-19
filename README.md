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

## AWS DynamoDB Core Table

### Additional Indexes
- PK SK Inverted GSI 
- AppID, SK GSI (H: SK, SK: AppID)
 
### Table schema

all in one table

---
#### Profile

|    PK     |      SK       |    AppID     | nickname | profile |       |
| :-------: | :-----------: | :----------: | :------: | :-----: | :---: |
| USER#uuid | PROFILE#email | Hashed AppID | Nickname | imgUrl  |
  
---

#### Openess
|    PK     |        SK         |
| :-------: | :---------------: |
| USER#uuid | OPEN#ACCOUNT#CODE |
| USER#uuid | OPEN#RECORD#CODE  |
| USER#uuid |  OPEN#TASK#CODE   |

> **ACCOUNT OPENNESS (SK)**  
> | CODE  | SEARCH & BI-FOLLOW | OPEN FOLLOW |
> | :---: | :----------------: | :---------: |
> |   0   |         X          |      X      |
> |   1   |         O          |      X      |
> |   2   |         O          |      O      |

> **RECORD OPENNESS (SK)**  
> | CODE  | `RANK FRIENDS` CANDIDATE | `RANK ALL` CANDIDATE |
> | :---: | :----------------------: | :------------------: |
> |   0   |            X             |          X           |
> |   1   |            O             |          X           |
> |   2   |            O             |          O           |


>  **TASK OPENNESS (SK)**  
> | CODE  | FRIENDS |  ALL  |
> | :---: | :-----: | :---: |
> |   0   |    X    |   X   |
> |   1   |    O    |   X   |
> |   2   |    O    |   O   |

---

#### Current Task

User get only 6 tasks row

|    PK     |      SK      |                    todo                    |    memo    |    where     |   willStart   | estimatedMinutes |  completedAt  |   createdAt   |
| :-------: | :----------: | :----------------------------------------: | :--------: | :----------: | :-----------: | :--------------: | :-----------: | :-----------: |
| USER#uuid | TASK#index   | [{"content": "todo", "isCompleted":false}] |            |              |
| USER#uuid | TASK#index   |                     []                     | MemoString | hanyang univ | 1604343297363 |       300        | 1604343441719 | 1604343257363 |

---

#### Record

|    PK     |        SK         |     tasks     |
| :-------: | :---------------: | :-----------: |
| USER#uuid | RECORD#YYYY-MM-DD | `Array<Task>` |

---

### Follow

|      PK       |      SK       |
| :-----------: | :-----------: |
| FOLLOWER#uuid | PROFILE#email |

---

### Follow Request

|       PK        |      SK       |
| :-------------: | :-----------: |
| REQ#FOLLOW#uuid | PROFILE#email |
