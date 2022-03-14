# 은하수 프로젝트

## 준비

- [소상공인시장진흥공단 상가(상권)정보](https://www.data.go.kr/data/15083033/fileData.do)

1. `/server/.data/` 디렉터리에 압축해제
2. `/server/` 디렉터리로 이동 
3. 설정 파일 (`server.config.yml`) 작성
4. `go run cmd/setup/main.go` 명령어 실행
5. 데이터베이스에 설정된 인덱스, 데이터 확인


## 설정

- `/server/server.config.yml`
  - 서버 기본 설정 파일
  - `/server/server.config.yml.example` 참조
