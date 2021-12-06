#1
Ссылка на инструкцию по установке elasticsearch https://www.elastic.co/guide/en/elasticsearch/reference/current/rpm.html
С её помощью создавал Dockerfile

>docker build -t elasticsearch:1.0 .

    Successfully built 3803b03b91b4
    Successfully tagged elasticsearch:1.0

>docker images

    REPOSITORY      TAG       IMAGE ID       CREATED         SIZE
    elasticsearch   1.0       3803b03b91b4   8 minutes ago   938MB 


Пример запуска контейнера single node
https://www.elastic.co/guide/en/elasticsearch/reference/current/docker.html

> docker run --rm -d -p 127.0.0.1:9200:9200 -p 127.0.0.1:9300:9300 -e "discovery.type=single-node" --name esc-node1 elasticsearch:1.0
>docker ps
    
    CONTAINER ID   IMAGE               COMMAND                  CREATED          STATUS          PORTS                                                NAMES
    c41074e24dda   elasticsearch:1.0   "/usr/share/elastics…"   49 seconds ago   Up 48 seconds   127.0.0.1:9200->9200/tcp, 127.0.0.1:9300->9300/tcp   esc-node1


[Dockerfile](Dockerfile)

[DocHub elasticsearch](https://hub.docker.com/repository/docker/chernyshovnetology/elasticsearch)

>GET 127.0.0.1:9200/

    {
      "name" : "netology_test",
      "cluster_name" : "chernyshov",
      "cluster_uuid" : "XMyB3es8R9i1zOsrR5UOqg",
      "version" : {
        "number" : "7.15.2",
        "build_flavor" : "default",
        "build_type" : "rpm",
        "build_hash" : "93d5a7f6192e8a1a12e154a2b81bf6fa7309da0c",
        "build_date" : "2021-11-04T14:04:42.515624022Z",
        "build_snapshot" : false,
        "lucene_version" : "8.9.0",
        "minimum_wire_compatibility_version" : "6.8.0",
        "minimum_index_compatibility_version" : "6.0.0-beta1"
      },
      "tagline" : "You Know, for Search"
    }


#2

>curl -X PUT localhost:9200/ind-1 -H 'Content-Type: application/json' -d'{ "settings": { "number_of_shards": 1,  "number_of_replicas": 0 }}'

    {"acknowledged":true,"shards_acknowledged":true,"index":"ind-1"}r

>curl -X PUT localhost:9200/ind-2 -H 'Content-Type: application/json' -d'{ "settings": { "number_of_shards": 2,  "number_of_replicas": 1 }}'
    
    {"acknowledged":true,"shards_acknowledged":true,"index":"ind-2"}

>curl -X PUT localhost:9200/ind-3 -H 'Content-Type: application/json' -d'{ "settings": { "number_of_shards": 4,  "number_of_replicas": 2 }}';

    {"acknowledged":true,"shards_acknowledged":true,"index":"ind-3"}


###Индексы
Статус Yellow потому что у них указано число реплик, а у меня одинокий сервер (кластер из одной ноды)
>curl -X GET 'http://localhost:9200/_cluster/health/ind-1?pretty'

    {
      "cluster_name" : "chernyshov",
      "status" : "green",
      "timed_out" : false,
      "number_of_nodes" : 1,
      "number_of_data_nodes" : 1,
      "active_primary_shards" : 1,
      "active_shards" : 1,
      "relocating_shards" : 0,
      "initializing_shards" : 0,
      "unassigned_shards" : 0,
      "delayed_unassigned_shards" : 0,
      "number_of_pending_tasks" : 0,
      "number_of_in_flight_fetch" : 0,
      "task_max_waiting_in_queue_millis" : 0,
      "active_shards_percent_as_number" : 100.0
    }

>curl -X GET 'http://localhost:9200/_cluster/health/ind-2?pretty'

    {
      "cluster_name" : "chernyshov",
      "status" : "yellow",
      "timed_out" : false,
      "number_of_nodes" : 1,
      "number_of_data_nodes" : 1,
      "active_primary_shards" : 2,
      "active_shards" : 2,
      "relocating_shards" : 0,
      "initializing_shards" : 0,
      "unassigned_shards" : 2,
      "delayed_unassigned_shards" : 0,
      "number_of_pending_tasks" : 0,
      "number_of_in_flight_fetch" : 0,
      "task_max_waiting_in_queue_millis" : 0,
      "active_shards_percent_as_number" : 44.44444444444444
    }

>curl -X GET 'http://localhost:9200/_cluster/health/ind-3?pretty'

    {
      "cluster_name" : "chernyshov",
      "status" : "yellow",
      "timed_out" : false,
      "number_of_nodes" : 1,
      "number_of_data_nodes" : 1,
      "active_primary_shards" : 4,
      "active_shards" : 4,
      "relocating_shards" : 0,
      "initializing_shards" : 0,
      "unassigned_shards" : 8,
      "delayed_unassigned_shards" : 0,
      "number_of_pending_tasks" : 0,
      "number_of_in_flight_fetch" : 0,
      "task_max_waiting_in_queue_millis" : 0,
      "active_shards_percent_as_number" : 44.44444444444444
    }

###Состояние кластера

>curl -XGET localhost:9200/_cluster/health/?pretty=true

    {
      "cluster_name" : "chernyshov",
      "status" : "yellow",
      "timed_out" : false,
      "number_of_nodes" : 1,
      "number_of_data_nodes" : 1,
      "active_primary_shards" : 8,
      "active_shards" : 8,
      "relocating_shards" : 0,
      "initializing_shards" : 0,
      "unassigned_shards" : 10,
      "delayed_unassigned_shards" : 0,
      "number_of_pending_tasks" : 0,
      "number_of_in_flight_fetch" : 0,
      "task_max_waiting_in_queue_millis" : 0,
      "active_shards_percent_as_number" : 44.44444444444444
    }

###Удаляю индексы

>curl -X DELETE 'http://localhost:9200/ind-1?pretty'

    {
      "acknowledged" : true
    }

>curl -X DELETE 'http://localhost:9200/ind-2?pretty'

    {
      "acknowledged" : true
    }

>curl -X DELETE 'http://localhost:9200/ind-3?pretty'

    {
      "acknowledged" : true
    }

#3

###Создаю директорию
>curl -XPOST localhost:9200/_snapshot/netology_backup?pretty -H 'Content-Type: application/json' -d'{"type": "fs", "settings": { "location":"/usr/share/elasticsearch/snapshots" }}'
    
    {
      "acknowledged" : true
    }

> GET  http://localhost:9200/_snapshot/netology_backup?pretty

    {
      "netology_backup" : {
        "type" : "fs",
        "settings" : {
          "location" : "/usr/share/elasticsearch/snapshots"
        }
      }
    }

>curl -X PUT localhost:9200/test -H 'Content-Type: application/json' -d'{ "settings": { "number_of_shards": 1,  "number_of_replicas": 0 }}'
  
    {
      "acknowledged" : true,
      "shards_acknowledged" : true,
      "index" : "test"
    }

###Создаю индекс
>GET http://localhost:9200/test?pretty

    {
      "test" : {
        "aliases" : { },
        "mappings" : { },
        "settings" : {
          "index" : {
            "routing" : {
              "allocation" : {
                "include" : {
                  "_tier_preference" : "data_content"
                }
              }
            },
            "number_of_shards" : "1",
            "provided_name" : "test",
            "creation_date" : "1638815601421",
            "number_of_replicas" : "0",
            "uuid" : "LOFN_gsNT8SK4I2PHhAEmw",
            "version" : {
              "created" : "7150299"
            }
          }
        }
      }
    }

###Снимаю snapshot
>curl -X PUT localhost:9200/_snapshot/netology_backup/elasticsearch?wait_for_completion=true
    
    {
      "snapshot" : {
        "snapshot" : "elasticsearch",
        "uuid" : "eL_2z5gdQIy1dEPVdJTHxA",
        "repository" : "netology_backup",
        "version_id" : 7150299,
        "version" : "7.15.2",
        "indices" : ["snapshots","test",".geoip_databases"],
        "data_streams" : [],
        "include_global_state" : true,
        "state" : "SUCCESS",
        "start_time" : "2021-12-06T18:35:03.018Z",
        "start_time_in_millis" : 1638815703018,
        "end_time" : "2021-12-06T18:35:04.419Z",
        "end_time_in_millis" : 1638815704419,
        "duration_in_millis" : 1401,
        "failures" : [],
        "shards":{
          "total" : 3,
          "failed" : 0,
          "successful" : 3
        },
        "feature_states" : [{
          "feature_name" : "geoip",
          "indices":[".geoip_databases"]
        }]
      }
    }

>docker exec -it esc-node1 /bin/bash
>bash-4.2$ cd /usr/share/elasticsearch/snapshots
>bash-4.2$ pwd
  
    /usr/share/elasticsearch/snapshots

>bash-4.2$ ls -la 

    total 56
    drwxr-xr-x 1 elasticsearch elasticsearch  4096 Dec  6 18:35 .
    drwxr-xr-x 1 root          root           4096 Dec  4 19:26 ..
    -rw-r--r-- 1 elasticsearch elasticsearch  1081 Dec  6 18:35 index-0
    -rw-r--r-- 1 elasticsearch elasticsearch     8 Dec  6 18:35 index.latest
    drwxr-xr-x 5 elasticsearch elasticsearch  4096 Dec  6 18:35 indices
    -rw-r--r-- 1 elasticsearch elasticsearch 27629 Dec  6 18:35 meta-eL_2z5gdQIy1dEPVdJTHxA.dat
    -rw-r--r-- 1 elasticsearch elasticsearch   471 Dec  6 18:35 snap-eL_2z5gdQIy1dEPVdJTHxA.dat

###Удаляю индекс

>curl -X DELETE 'http://localhost:9200/test?pretty'

    {
      "acknowledged" : true
    }

###Создаю индекс 2
>curl -X PUT localhost:9200/test-2?pretty -H 'Content-Type: application/json' -d'{ "settings": { "number_of_shards": 1,  "number_of_replicas": 0 }}'
    
    {
      "acknowledged" : true,
      "shards_acknowledged" : true,
      "index" : "test-2"
    }

###Восстанавливаю индекс
>curl -X POST localhost:9200/_snapshot/netology_backup/elasticsearch/_restore?pretty -H 'Content-Type: application/json' -d'{"include_global_state":true}'
    
    {
      "accepted" : true
    }

###Проверяю индексы
>curl -X GET http://localhost:9200/_cat/indices?v

    health status index            uuid                   pri rep docs.count docs.deleted store.size pri.store.size
    green  open   .geoip_databases PDmgfwwxSlKWTS-omWYhiw   1   0         43           86    156.5mb        156.5mb
    yellow open   snapshots        su8GvgPTTaeymLQOy55xSQ   1   1          2            0      9.1kb          9.1kb
    green  open   test-2           vCAZN7uwTJe99h_Gr_RxPw   1   0          0            0       208b           208b
    green  open   test             5HSiKVN_R2ySv_e1EeaEIw   1   0          0            0       208b           208b

