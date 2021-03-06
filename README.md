# STMT Core backend

[English](./README_en.md)

STMT 애플리케이션을 위한 서버리스 Golang 코어 서버입니다.
 
## API

### AUTH (AUTH API에서만 인증 토큰이 필요하지 않습니다.)
- [x] 구글 인증으로 회원가입, 로그인하기

### USER
- [x] 로그인 한 유저 MyPage 가져오기
- [ ] 계정 삭제하기
- [x] 알람 설정 업데이트하기
- [x] 공개 설정 가져오기
- [ ] 공개 설정 업데이트

### SOCIAL
- [x] 이메일로 유저 검색
- [x] 유저 이메일로 팔로우하기
- [x] 유저 이메일로 언팔로우하기
- [ ] 팔로우 하는 유저들 랭킹 보기
- [x] 모든 공개 계정의 랭킹 확인하기
- [ ] 팔로우 하는 유저의 대시보드 데이터 가져오기
- [ ] 공개된 계정의 대시보드 데이터 가져오기

### TASK
- [x] 유저의 현재 태스크 리스트 잠금 (서버에 저장)
- [x] 유저의 현재 태스크 리스트 가져오기
- [x] 유저의 현재 태스크를 비우고, 레코드를 만들기
- [x] 현재 태스크의 진행도 업데이트하기

### RECORD
- [x] 로그인한 유저의 대시보드 데이터 가져오기
- [x] 로그인한 유저의 대시보드 자세한 기록 가져오기
- [ ] Get other user's a week's amount of record by specific date and user email (only if the user has given permission)

## AWS DynamoDB Core Table

AWS에서 권장한 [한 테이블에 설계하는 방법](https://changhoi.github.io/posts/backend/dynamodb-single-table-design/)을 따랐습니다.

### Secondary Indexes
- PK SK Inverted GSI 
- AppID, SK GSI (H: SK, SK: AppID)
 
### 테이블 스키마
---
#### Profile

|    PK     |      SK       |    AppID     | nickname | profile |
| :-------: | :-----------: | :----------: | :------: | :-----: |
| USER#uuid | PROFILE#email | Hashed AppID | Nickname | imgUrl  |
  
---

#### Openess
|    PK     |        SK         |
| :-------: | :---------------: |
| USER#uuid | OPEN#ACCOUNT#CODE |
| USER#uuid | OPEN#RECORD#CODE  |
| USER#uuid |  OPEN#TASK#CODE   |

>  **ACCOUNT OPENNESS (SK)** (default: 2)  
> | CODE  | SEARCH & FOLLOW REQUEST | OPEN FOLLOW |
> | :---: | :---------------------: | :---------: |
> |   0   |            X            |      X      |
> |   1   |            O            |      X      |
> |   2   |            O            |      O      |

> **RECORD OPENNESS (SK)**  (default: 2)
> | CODE  | `RANK FRIENDS` CANDIDATE | `RANK ALL` CANDIDATE |
> | :---: | :----------------------: | :------------------: |
> |   0   |            X             |          X           |
> |   1   |            O             |          X           |
> |   2   |            O             |          O           |


>  **TASK OPENNESS (SK)**  (default: 2)
> | CODE  | FRIENDS |  ALL  |
> | :---: | :-----: | :---: |
> |   0   |    X    |   X   |
> |   1   |    O    |   X   |
> |   2   |    O    |   O   |

---

#### Current Task

유저마다 6개의 row를 갖게 됩니다.

|    PK     |     SK     |                                  todo                                  |    memo    |    where     |   willStart   | estimatedMinutes |  completedAt  |   createdAt   |
| :-------: | :--------: | :--------------------------------------------------------------------: | :--------: | :----------: | :-----------: | :--------------: | :-----------: | :-----------: |
| USER#uuid | TASK#index |                                   []                                   |            |              |               |                  |               | 1604343057363 |
| USER#uuid | TASK#index | [{"content": "todo", "isCompleted":false, "createdAt": 1604343257363}] | MemoString | hanyang univ | 1604343297363 |       300        | 1604343441719 | 1604343257363 |

---

#### Record

|    PK     |                 SK                  | LockTime  |     Tasks     | Score | InComplete | Complete | Percent | Duration  | RecordOpenness | Nickname |
| :-------: | :---------------------------------: | :-------: | :-----------: | :---: | :--------: | :------: | :-----: | :-------: | :------------: | :------: |
| USER#uuid | RECORD#YYYY#MM#WeekOfYear#DayOfYear | timestamp | `Array<Task>` | 33.33 |     1      |    2     |  33.33  | timestamp |       2        |  string  |

---

### Follow

|      PK       |      SK       | ProfileUUID | FollowerEmail |
| :-----------: | :-----------: | :---------: | :-----------: |
| FOLLOWER#uuid | PROFILE#email |    uuid     |     email     |

---

### Follow Request

|       PK        |      SK       |
| :-------------: | :-----------: |
| REQ#FOLLOW#uuid | PROFILE#email |
