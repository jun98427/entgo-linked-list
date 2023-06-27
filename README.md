# entgo-linked-list

### edit .env file example
```yaml
POSTGRES_USER=postgres
POSTGRES_PASSWORD=postgres
POSTGRES_DB=postgres
HOST = localhost
```

### DATABASE START
```bash
docker-compose --env-file ./.env up -d
```

### start go file
```bash
go run main.go
```

### 동작 방식
#### add node
1. node 가 존재하지 않을 경우, head node 로 추가
2. node 가 존재하고 가장 처음에 추가하고 싶은 경우, prevID 를 0 으로 설정
3. node 가 존재하고 중간에 추가하고 싶은 경우, prevID 를 추가하고 싶은 node 의 ID 로 설정

#### delete node
1. node 가 존재하지 않을 경우, 에러 발생

#### print list
1. node 가 존재하지 않을 경우, empty list 출력
2. node 가 존재할 경우, linked list head 부터 출력