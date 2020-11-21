# STMT Core backend

Golang serverless backend for STMT Application

## Requirements

- [ ] Build script & Makefile
- [ ] Serverless framework setup
- [ ] Service Logic
 
## API

### AUTH (Authentication is not required only in this part)
- [x] Signup & Signin with google login

### USER
- [x] Get user info by uuid (My page, Get user)
- [ ] Remove login user account
- [x] Update alarm setting
- [x] Get openness setting
- [ ] Update openness setting
- [ ] Search user by email
- [ ] Follow user by user email
- [ ] Unfollow user by user email

### TASK
- [x] Lock user's current task list (create current task list)
- [x] Get user's current task list (get locked task list)
- [x] Clear user's current task & create a record (when the current task list lock time passes)
- [ ] Update progress of current tasks (ex. Check the task complete)

### Record
- [ ] Get login user's dashboard data
- [ ] Get other user's a week's amount of record by specific date and user email (only if the user has given permission)

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

> **ACCOUNT OPENNESS (SK)** (default: 1)  
> | CODE  | SEARCH & BI-FOLLOW | OPEN FOLLOW |
> | :---: | :----------------: | :---------: |
> |   0   |         X          |      X      |
> |   1   |         O          |      X      |
> |   2   |         O          |      O      |

> **RECORD OPENNESS (SK)**  (default: 1)
> | CODE  | `RANK FRIENDS` CANDIDATE | `RANK ALL` CANDIDATE |
> | :---: | :----------------------: | :------------------: |
> |   0   |            X             |          X           |
> |   1   |            O             |          X           |
> |   2   |            O             |          O           |


>  **TASK OPENNESS (SK)**  (default: 1)
> | CODE  | FRIENDS |  ALL  |
> | :---: | :-----: | :---: |
> |   0   |    X    |   X   |
> |   1   |    O    |   X   |
> |   2   |    O    |   O   |

---

#### Current Task

User get only 6 tasks row

|    PK     |     SK     |                    todo                    |    memo    |    where     |   willStart   | estimatedMinutes |  completedAt  |   createdAt   |
| :-------: | :--------: | :----------------------------------------: | :--------: | :----------: | :-----------: | :--------------: | :-----------: | :-----------: |
| USER#uuid | TASK#index | [{"content": "todo", "isCompleted":false}] |            |              |
| USER#uuid | TASK#index |                     []                     | MemoString | hanyang univ | 1604343297363 |       300        | 1604343441719 | 1604343257363 |

---

#### Record

|    PK     |        SK        |     tasks     |
| :-------: | :--------------: | :-----------: |
| USER#uuid | RECORD#timestamp | `Array<Task>` |

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
