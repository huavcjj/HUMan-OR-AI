# Bot-or-Not
ハッカソン技育CAMP2024 Vol.19

forntend:

backend :

1. ディレクトリに移動                　
- $ cd backend/ 

2. 必要なGoモジュールをインストール
- $ go mod tidy 

3. Dockerを起動
- $ docker compose up -d 

4. サーバーを起動
- $ go run cmd/main.go 

5. 動作確認
- $ curl -X POST http://localhost:8080/game/1

- $ curl -X POST http://localhost:8080/player/1 \
     -H "Content-Type: application/json" \
     -d '{"name": "Player1"}'

- $ curl -X GET http://localhost:8080/player/1


6. Dockerを停止
- $ docker compose down


