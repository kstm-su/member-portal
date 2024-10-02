# Member portalとは
kstmメンバーであることを確認し、また事務処理を簡潔化するためのポータルサイトです。

## Common
### Redoc(OpenAPI)にアクセスするには
[Redoc](redoc)

## Frontend
### Dockerでの起動方法
#### Localでビルドする場合
1. Dockerfileのあるディレクトリに移動
2. `docker build -t <image_name> .` を実行（image_nameは自由に設定）
3. `docker run -p 3000:3000 -d <image_name>` を実行
4. `http://localhost:3000` にアクセス

#### GitHub container registryからpullする場合
1. `docker pull ghcr.io/kstm-su/member-portal/frontend:latest` を実行
2. `docker run -p 3000:3000 -d ghcr.io/kstm-su/member-portal/frontend:latest` を実行
3. `http://localhost:3000` にアクセス

### Nodeでローカルで起動する場合
1. `cd member-portal-frontend` でディレクトリに移動
2. `npm install` を実行
3. `npm run dev` を実行
4. `http://localhost:3000` にアクセス


## Backend

### Dockerでの起動方法
#### Localでビルドする場合

1. Dockerfileのあるディレクトリに移動
2. `docker build -t <image_name> .` を実行（image_nameは自由に設定）
3. `docker run -p 3001:8080 -d <image_name>` を実行
4. `http://localhost:3001` にアクセス

#### GitHub container registryからpullする場合
1. `docker pull ghcr.io/kstm-su/member-portal/backend:latest` を実行
2. `docker run -p 3001:8080 -d ghcr.io/kstm-su/member-portal/backend:latest` を実行
3. `http://localhost:3001` にアクセス

### Goでローカルで起動する場合
1. `cd member-portal-backend` でディレクトリに移動
2. `go run main.go` を実行
3. `http://localhost:8080` にアクセス
