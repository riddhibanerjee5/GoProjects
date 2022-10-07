#!/bin/sh

docker logout
docker build -t myfalcotestcases .
docker tag myfalcotestcases riddhibanerjee/myfalcotestcases
docker login
docker push riddhibanerjee/myfalcotestcases
echo "Image has been updated and pushed to docker hub."
