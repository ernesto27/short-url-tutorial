#bin/bash

docker exec mysql-short-url mysqldump -uroot -p1111 short-url > bkp_$(date +\%Y\%m\%d\%H\%M\%S).sql

# TODO get only latest 10 files

# TODO upload backup to a bucket on S3