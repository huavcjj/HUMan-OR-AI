# Bot-or-Not
ハッカソン技育CAMP2024 Vol.19

forntend:
1. Node.jsのインストール
   
https://nodejs.org/en/download/prebuilt-installer
このURLからver20.18.1(LTS)を選択しインストール(wizardでは全てNextをクリックしてok)


2. インストールの確認
   vscodeのターミナルを開き,"node -v" "npm -v"とそれぞれ入力し
   バージョンが見れる事を確認してください。

3.　ディレクトリの移動
　　cd コマンドでfrontendディレクトリに移動してください。

4. プロジェクト立ち上げ
　　vscodeのターミナルで"npm run dev"を実行してください。
　　無理だった場合、以下をそれぞれターミナルで実行し、再度"npm run dev"を実行してください。

　　winget install Schniz.fnm

　　fnm env --use-on-cd | Out-String | Invoke-Expression
  
　　fnm use --install-if-missing 20

　　成功するとターミナルにlocalhost:3000のURLが発行されるのでクリックして開いてください。

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

6. Dockerを停止
- $ docker compose down


