## Fitur / Arsitektur

- Clean Architecture
- MySQL ACID
- Concurrency safe
- JWT Auth
- Rate limiting
- Security headers
- Unit test
- Docker
- 50 request / second
- CSP
- SQL Injection safe

## Instalasi

1>. download repo 

https://github.com/dennis2908/financexyzgolang

2>. Jalankan dengan Docker
Start semua service
docker-compose up -d --build

Kalau sukses:

Listening and serving HTTP on :8080

<img width="1068" height="254" alt="image" src="https://github.com/user-attachments/assets/aa728621-000b-4805-8384-973a4a282907" />

3>. Setup database

Open terminal baru:

Masuk mysql:

docker exec -it xyz-finance-production-db-1 mysql -uroot -proot xyz

atau:

docker exec -it <container_name_db> mysql -uroot -proot xyz

Run schema

Copy isi file:

migrations/schema.sql

paste ke mysql

atau:

docker exec -i <db_container> mysql -uroot -proot xyz < migrations/schema.sql

4>. Insert dummy data (penting biar bisa test)

Masuk MySQL:

INSERT INTO limits(user_id, tenor, limit_amount, used_amount)
VALUES (1, 1, 1000000, 0);

Artinya:

user 1 limit = 1.000.000

5. Test API dan concurrency

Ke Postman dan import Postman dari folder postman ke postman

Hasil:

1 success
1 limit exceeded

Response sukses
{
  "message": "success"
}

Lakukan beberapa kali dengan hasil : 

Kalau limit habis
amount: 600000

Response:

{
  "error": "limit exceeded"
}

<img width="807" height="733" alt="image" src="https://github.com/user-attachments/assets/f8ad536e-617d-4f94-87d9-4616c717cf1f" />

Untuk UNIT Test

di cmd :

docker-compose run --rm app go test -v -race ./...

<img width="1062" height="245" alt="image" src="https://github.com/user-attachments/assets/6abb3f83-311c-47b6-ae76-55fc7f47126c" />
