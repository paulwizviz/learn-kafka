#!/bin/bash

export NETWORK=learn-kafka_simple-Network

COMMAND="$1"

case $COMMAND in
    "start")
        docker-compose -f ./deployment/single_broker.yml up
        ;;
    "stop")
        docker-compose -f ./deployment/single_broker.yml down
        ;;
    "topic")
        docker exec broker kafka-topics --bootstrap-server broker:9092 \ 
                                        --create \
                                        --topic quickstart
        ;;
    *)
        echo "$0 [start | stop]"
        ;;
esac