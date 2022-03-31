# 은하수 프로젝트

## 백엔드

### 환경

- Go
- Fiber
- MongoDB

### 준비

- [Kakao REST API Key](https://developers.kakao.com)


- [소상공인시장진흥공단 상가(상권)정보](https://www.data.go.kr/data/15083033/fileData.do)

1. `/backend/.data/` 디렉터리에 압축해제
2. `/backend/` 디렉터리로 이동
3. 설정 파일 (`server.config.yml`) 작성
4. `go run cmd/setup/main.go` 명령어 실행
5. 데이터베이스에 추가된 인덱스, 데이터 확인

### 설정

- `/backend/server.config.yml`
    - 서버 기본 설정 파일
    - `/backend/server.config.yml.example` 참고

## 프론트엔드

### 환경

- Svelte
- TypeScript
