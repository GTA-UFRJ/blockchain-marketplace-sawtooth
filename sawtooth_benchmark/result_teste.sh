#!/bin/bash

#docker stop $(docker ps -a -q)  
#docker rm -f $(docker ps -aq)
#docker rmi -f $(docker images | grep dev | awk '{print $3}') 
#docker volume prune -f
#docker ps -a
#docker images -a
#docker volume ls
docker-compose -f docker-poet-1.yaml down --remove-orphans -v
./get_result.sh
