# member-portal
kstmメンバーであることを確認し、また事務処理を簡潔化するためのポータルサイト

## Frontend
### Docker起動方法
1. Dockerfileのあるディレクトリに移動
2. `docker build -t <image_name> .` を実行（image_nameは自由に設定）
3. `docker run -p 3000:3000 -d <image_name>` を実行
4. `http://localhost:3000` にアクセス

※ ローカルで起動する場合は、`npm run dev` を実行

## Backend 

### Docker起動方法
1. Dockerfileのあるディレクトリに移動
2. `docker build -t <image_name> .` を実行（image_nameは自由に設定）
3. `docker run -p 3001:8080 -d <image_name>` を実行
4. `http://localhost:3001` にアクセス


