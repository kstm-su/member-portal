# コンフィグの設定について

```yaml
database:
  type: sqlite  
  sqlite:
    path: /app/data/db.sqlite3
file:
  base: /app/data
jwt:
  issuer: localhost
  key:
    private_key: private.pem
    public_key: public.pem
  key_id: key
  realm: localhost
password:
  algorithm: argon2
  pepper: <random_pepper> #自動で生成されるため、設定不要
server:
  host: localhost
  port: 8080
```

## database
postgresqlを利用する場合、以下のように設定を変更してください。

```yaml
database:
  type: postgres
  postgres:
    host: db
    port: 5432
    user: <user>
    password: <password>
```


