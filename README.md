# Short url - monolitic version


## System achitecture diagram
![Diagram](./system_architecture.png)


## Setup virtual machine

Setup ssh key

After log in on your VM, create a ssh key pair and add the public key to your github account.

```bash
$ ssh-keygen -t ed25519 -C "your_email@example.com" 
```

Clone the repository 

```bash
$ git clone git@github.com:yourrepo.git
```

Install VM dependencies

```bash
$ ./installVM.sh
```
Response yes if prompt appears.
After scirpt finish, logout and login again on the server.

Go to the repo folder and run this command 
```bash
$ docker compose up -d mysql 
```

Create tables 
```bash
$ docker exec -i mysql-short-url mysql -uroot -p1111 < ./init.sql
```

Init services
```bash
$ docker compose up -d backend nginx redis 
```

---

## CI/CD
The project have a github actions configurations that runs two jobs define in this files.

**backend.yml**

This job runs on every commit, execute a static analysis check and tests.


**build-deploy.yml**

This job runs only when a new tag is created, it builds a docker image, push to docker hub and connect to a VM via ssh in order to run the lastest version of the service

---

## Setup local environment

Setup Mysql 

Go to the repo folder and run this command 
```bash
$ docker compose -f docker-compose-dev.yaml up -d mysql 
```

Create tables 
```bash
$ docker exec -i mysql-short-url mysql -uroot -p1111 < ./init.sql
```

Init backend, nginx

Run this command to init backend, redis and nginx on background
```bash
$ docker compose up -d backend nginx redis 
```

### Backup db 
```bash
$ docker exec mysql-short-url mysqldump -uroot -p1111 short-url > bkp.sql
```

Import dump sql 
```bash
$ docker exec -i mysql-short-url mysql -uroot -p1111 short-url < bkp.sql
```
--- 
### Test endpoints

Create a short url
```bash
$ curl -X POST -H "Content-Type: application/json" -d '{"url":"http://site.com"}' http://localhost/create-url
```

Get short url 
```bash
$ curl -X GET http://localhost/hash-url
```






