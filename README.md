svc-master


Hal yang perlu dipersiapkan (tersedia dalam folder collction):

- Database schema
- Postman collection


Berikut adalah langkah-langkah untuk menjalankan service svc-master:

1. Masuk folder /cmd
2. Buat file .env dengan format berikut


```
DB_POSTGRES_USER=postgres
DB_POSTGRES_PASSWORD=postgres
DB_POSTGRES_HOST=localhost
DB_POSTGRES_PORT=5432
DB_POSTGRES_NAME=DiskhaTek
DB_POSTGRES_MAX_IDLE_CONNS=5
DB_POSTGRES_MAX_OPEN_CONNS=10
DB_POSTGRES_MAX_LIFE_TIME=3s
DB_POSTGRES_FLAVOR=
DB_POSTGRES_LOCATION=
DB_POSTGRES_TIMEOUT=20s

DB_MYSQL_USER=
DB_MYSQL_PASSWORD=
DB_MYSQL_HOST=
DB_MYSQL_PORT=
DB_MYSQL_NAME=
DB_MYSQL_MAX_IDLE_CONNS=
DB_MYSQL_MAX_OPEN_CONNS=
DB_MYSQL_MAX_LIFE_TIME=
DB_MYSQL_FLAVOR=
DB_MYSQL_LOCATION=
DB_MYSQL_TIMEOUT=

APP_HOST=localhost
APP_ENDPOINT_V=V1
APP_ENVIRONMENT=DEVELOPMENT
APP_PORT=3000

TOKEN_LIFESPAN=10
TOKEN_SECRET=DiskhaTek

```

3. Jalankan command go run main.go