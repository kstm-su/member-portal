# kstm-member-portal backend side

This is the backend part of kstm-member-portal. <br>
It is implemented by golang and [echo](https://github.com/gofiber/fiber).

## Getting Started
### Go
```bash 
go run main.go
```

or 

```bash
go build -o kstm-member-portal
./kstm-member-portal
```

### Docker
```bash
docker build -t kstm-member-portal .
docker run -p 8080:8080 kstm-member-portal
```

## Configuration
引数で設定ファイルを指定することができます。デフォルトは`/app/config.yaml`です。
```bash
./kstm-member-portal --config /path/to/config.yaml
```



