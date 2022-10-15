#!/bin/bash

export NETWORK=learn-kafka_simple-Network

COMMAND="$1"

case $COMMAND in
    "start")
        docker-compose -f ./deployment/single_instance.yml up
        ;;
    "stop")
        docker-compose -f ./deployment/single_instance.yml down
        ;;
    *)
        echo "$0 [start | stop]"
        ;;
esac