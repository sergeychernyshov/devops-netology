#1

###В чём отличие режимов работы сервисов в Docker Swarm кластере: replication и global?

    replication - режим работы при котором можно указать количество нод для выполенения одной задачи.
    Например, развернуть сервис HTTP с тремя репликами, каждая из которых обслуживает один и тот же контент.

    global - режим работы при котором запускается одна задача на каждой ноде.

###Какой алгоритм выбора лидера используется в Docker Swarm кластере?

    Отказоустойчивость сервиса гарантируется самим Docker. Это достигается за счет того, что в кластере могут 
    одновременно работать несколько управляющих нод, которые могут в любой момент заменить вышедшего из строя лидера. 
    Если говорить более подробно, то используется так называемый алгоритм поддержания распределенного консенсуса — Raft. 

###Что такое Overlay Network?

    Overlay Network - общий случай логической сети, создаваемой поверх другой сети. Узлы оверлейной сети могут быть 
    связаны либо физическим соединением, либо логическим, для которого в основной сети существуют один или несколько 
    соответствующих маршрутов из физических соединений.

#2

>yc init
>yc vpc network create --name net-netology
    
    id: enpmt19ci88f7sgod3q0
    folder_id: b1g91q2tu6if6gi7erv4
    created_at: "2021-11-20T18:48:36Z"
    name: net-netology

>yc vpc subnet create --name subnet-netology --zone ru-central1-a --range 10.1.2.0/24 --network-name net-netology

    id: e9bbm9b3thhcb7qt4ue1
    folder_id: b1g91q2tu6if6gi7erv4
    created_at: "2021-11-20T18:49:27Z"
    name: subnet-netology
    network_id: enpmt19ci88f7sgod3q0
    zone_id: ru-central1-a
    v4_cidr_blocks:
    - 10.1.2.0/24

>packer validate centos-7-base.json
    
    The configuration is valid.

>packer build centos-7-base.json

    ==> Builds finished. The artifacts of successful builds are:
    --> yandex: A disk image was created: centos-7-base (id: fd89tni1qcf7ks80phv5) with family name centos

>yc compute image list

    +----------------------+---------------+--------+----------------------+--------+
    |          ID          |     NAME      | FAMILY |     PRODUCT IDS      | STATUS |
    +----------------------+---------------+--------+----------------------+--------+
    | fd89tni1qcf7ks80phv5 | centos-7-base | centos | f2ebfhrshe5m6i4saf1j | READY  |
    +----------------------+---------------+--------+----------------------+--------+

>yc vpc subnet delete e9bbm9b3thhcb7qt4ue1
>yc vpc network delete enpmt19ci88f7sgod3q0 

>yc iam key create --folder-name chernyshov-sergey --service-account-name chernyshov-sergey --output key.json


>terraform init

>terraform validate
    
    Success! The configuration is valid.

>terraform plan

>terraform apply -auto-approve


    external_ip_address_node01 = "51.250.5.40"
    external_ip_address_node02 = "51.250.3.44"
    external_ip_address_node03 = "51.250.2.123"
    external_ip_address_node04 = "51.250.6.50"
    external_ip_address_node05 = "51.250.7.127"
    external_ip_address_node06 = "51.250.7.119"
    internal_ip_address_node01 = "192.168.101.11"
    internal_ip_address_node02 = "192.168.101.12"
    internal_ip_address_node03 = "192.168.101.13"
    internal_ip_address_node04 = "192.168.101.14"
    internal_ip_address_node05 = "192.168.101.15"
    internal_ip_address_node06 = "192.168.101.16"

> ssh centos@51.250.5.40
> sudo -i
> docker node ls

    ID                            HOSTNAME             STATUS    AVAILABILITY   MANAGER STATUS   ENGINE VERSION
    v8d9l8g7waetjg6ghrwnc29ma *   node01.netology.yc   Ready     Active         Leader           20.10.11
    yu7wbhokx4kdo8quogaqy26sn     node02.netology.yc   Ready     Active         Reachable        20.10.11
    gjtjtqtt9kwn87un5ub8wvzg3     node03.netology.yc   Ready     Active         Reachable        20.10.11
    s31ojkhygswhxe4wi2qopt6jh     node04.netology.yc   Ready     Active                          20.10.11
    vtwsusoka8at5bef3bi8yjhf6     node05.netology.yc   Ready     Active                          20.10.11
    piydm0wuux5rg5gr2vveiyonv     node06.netology.yc   Ready     Active                          20.10.11

#3

>docker service ls

    ID             NAME                                MODE         REPLICAS   IMAGE                                          PORTS
    li7t4k6h3hi3   swarm_monitoring_alertmanager       replicated   1/1        stefanprodan/swarmprom-alertmanager:v0.14.0
    sm4tgzh0fii1   swarm_monitoring_caddy              replicated   1/1        stefanprodan/caddy:latest                      *:3000->3000/tcp, *:9090->9090/tcp, *:9093-9094->9093-9094/tcp
    4zb2fqk3loqq   swarm_monitoring_cadvisor           global       6/6        google/cadvisor:latest
    uc5m6km77vci   swarm_monitoring_dockerd-exporter   global       6/6        stefanprodan/caddy:latest
    u57iuhw3yhyd   swarm_monitoring_grafana            replicated   1/1        stefanprodan/swarmprom-grafana:5.3.4
    moxo0oyrntyc   swarm_monitoring_node-exporter      global       6/6        stefanprodan/swarmprom-node-exporter:v0.16.0
    lbn6kpu202nd   swarm_monitoring_prometheus         replicated   1/1        stefanprodan/swarmprom-prometheus:v2.5.0
    j4xtf4da4h1x   swarm_monitoring_unsee              replicated   1/1        cloudflare/unsee:v0.8.0
