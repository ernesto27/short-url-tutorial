# Short url - monolitic version

# Setup

### Setup key ssh
After log in on your VM, create a ssh key pair and add the public key to your github account.

```bash
$ ssh-keygen -t ed25519 -C "your_email@example.com" 
```

Clone the repository 

```bash
$ git clone git@github.com:yourrepo.git
```

Install docker

```bash
$ ./docker.sh
```
Response yes if prompt appears.
After scirpt finish, logout and login again on the server.



### Setup Mysql 

Go to the repo folder and run this command 
```bash
$ docker compose up mysql -d 
```

Create tables 
```bash
$ docker exec -i shorturl-mysql-1 mysql -uroot -p1111 < ./init.sql
```

### Init backend, nginx

Run this command to init backend and nginx on background
```bash
$ docker compose up backend nginx up -d
```

### Backup db 
```bash
$ docker exec mysql-short-url mysqldump -uroot -p1111 short-url > bkp.sql
```

### Import dump sql 
```bash
$ docker exec -i mysql-short-url mysql -uroot -p1111 short-url < bkp.sql
```

### Test endpoints

Create a short url
```bash
$ curl -X POST -H "Content-Type: application/json" -d '{"url":"http://site.com"}' http://localhost/create-url
```

Get short url 
```bash
$ curl -X GET http://localhost/hash-url
```






