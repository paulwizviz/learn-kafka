# Single Broker Cluster

This cluster is a reproduction of [confluence quick start](https://developer.confluent.io/quickstart/kafka-docker/). There is a working version here [./deployment/single_broker.yml](../deployment/single_broker.yml).

To work with this deployment, you will have to install docker and bash shell.

To help you start and stop the cluster, use this script [./scripts/sb-devops.sh](../scripts/sb-devops.sh). Here are commands:

* `./scripts/sb-devops.sh start` - start cluster
* `./scripts/sb-devops.sh stop`  - stop cluster 
* `./scripts/sb-devops.sh topic` - create topic name quick start

Once you have this running cluster, you can interact with it as per [confluence quick start](https://developer.confluent.io/quickstart/kafka-docker/).