#bin/bash

yes | sudo apt update

yes | sudo apt install apt-transport-https ca-certificates curl software-properties-common

yes | curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo gpg --dearmor -o /usr/share/keyrings/docker-archive-keyring.gpg

yes | echo "deb [arch=$(dpkg --print-architecture) signed-by=/usr/share/keyrings/docker-archive-keyring.gpg] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable" | sudo tee /etc/apt/sources.list.d/docker.list > /dev/null

yes | sudo apt update

yes | sudo apt install docker-ce

yes | sudo usermod -aG docker ${USER}


echo "Logout and login again on server to use docker without sudo"


# Add cron job to backup mysql database every minute on $HOME folder
crontab -l | { cat; echo "* * * * * ~/code/projects/shorturl/backupDb.sh"; } | crontab -