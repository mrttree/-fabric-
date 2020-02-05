#!/bin/bash
echo "启动容器: orderer..."
docker-compose -f docker-orderer.yaml up -d
##firstparty启动
echo "启动容器: firstparty.peer0 and cli ..."
docker-compose -f docker-peer0-firstparty.yaml up -d
echo "启动容器: firstparty.peer1 ..."
docker-compose -f docker-peer1-firstparty.yaml up -d
##superv启动
echo "启动容器: superv.peer0 ..."
docker-compose -f docker-peer0-superviser.yaml up -d
echo "启动容器: superv.peer1 ..."
docker-compose -f docker-peer1-superviser.yaml up -d
##install启动
echo "启动容器: install.peer0 ..."
docker-compose -f docker-peer0-installer.yaml up -d
echo "启动容器: install.peer1 ..."
docker-compose -f docker-peer1-installer.yaml up -d
##produce启动
echo "启动容器: produce.peer0 ..."
docker-compose -f docker-peer0-producer.yaml up -d
echo "启动容器: produce.peer1 ..."
docker-compose -f docker-peer1-producer.yaml up -d
##partproduce启动
echo "启动容器: partproduce.peer0 ..."
docker-compose -f docker-peer0-partproducer.yaml up -d
echo "启动容器: partproduce.peer1 ..."
docker-compose -f docker-peer1-partproducer.yaml up -d
##transit启动
echo "启动容器: transit.peer0 ..."
docker-compose -f docker-peer0-transiter.yaml up -d
echo "启动容器: transit.peer1 ..."
docker-compose -f docker-peer1-transiter.yaml up -d
