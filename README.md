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

4. airをインストール(再ビルド・再起動) 
- $ go install github.com/cosmtrek/air@latest 

5. サーバーを起動
- $ air

6. 動作確認
- $ curl -X POST http://localhost:8080/game

7. Dockerを停止
- $ docker compose down


