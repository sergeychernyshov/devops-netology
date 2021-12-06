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


