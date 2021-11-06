#1

>docker pull nginx

>docker run --name nginx-template -p 8080:80 -d nginx

>docker container ps

    CONTAINER ID   IMAGE     COMMAND                  CREATED          STATUS          PORTS                  NAMES
    e5d35e385a6a   nginx     "/docker-entrypoint.…"   42 seconds ago   Up 40 seconds   0.0.0.0:8080->80/tcp   nginx-template

>docker exec -it nginx-template bash

не знаю как редактировать правильно данные внутри контейнера, выбрал способ установки VIM
>apt-get update
>apt-get install vim 

Настроил домашнюю страницу
>docker commit nginx-template

>docker images
    
    REPOSITORY   TAG       IMAGE ID       CREATED          SIZE
    <none>       <none>    ab7c21f9a2cf   36 seconds ago   184MB
    nginx        latest    87a94228f133   3 weeks ago      133MB

>docker tag ab7c21f9a2cf chernyshovnetology/nginx-template

>docker tag ab7c21f9a2cf chernyshovnetology/nginx-template

>docker images

    REPOSITORY                          TAG       IMAGE ID       CREATED              SIZE
    chernyshovnetology/nginx-template   latest    ab7c21f9a2cf   About a minute ago   184MB
    nginx                               latest    87a94228f133   3 weeks ago          133MB

>docker push chernyshovnetology/nginx-template 

ссылка на репозиторий
https://hub.docker.com/repository/docker/chernyshovnetology/nginx-template

>docker pull chernyshovnetology/nginx-template

>docker run --name nginx-template-test -p 9090:80 -d chernyshovnetology/nginx-template

#2

###Высоконагруженное монолитное java веб-приложение
    Склоняюсь к виртуальной машине: будет легко масштабировать приложение (проще поднять ноду), легко переехать с 
    одного сервера на другой.
    Web приложение доступно многим пользователям, возможны различные атаки на приложение. В этом случае слой
    виртуальной машины будет дополнительным слоем защиты от злоумышленника.

    Монолитное приложение будет требовать частых изменений кода и реконфигураций сервера докер для такого случая
    использовать не хочется. 

###Nodejs веб-приложение
    Все зависит от функциональности: например, если это простое одностраничное приложение с графиком числа запросов 
    к серверу в секунду или простой web сервис выполняющий сжатие картинки, то можно поставить в Docker.
    
    Если это приложение требующее сложных вычислений или рендеринг картинки, в этом случае использую виртуализацию 
    или рассмотрю вариант выделенного сервера.

###Мобильное приложение c версиями для Android и iOS;
    Тут два варианта: если приложение имеет микросервисную архитектуру выберу докеры.
    Если приложение монолитное, то буду рассматривать виртуализацию.
    Выделенный сервер не буду рассматривать в этом случае, т.к. считаю что все что доступно во внешней сети ставить
    на выделенный сервер не безопасно.

###Шина данных на базе Apache Kafka;
    Предполагаю, что потеря данных в шине является критичной, то докер рассматривать не стану.
    Мгновенно отклика шина данных также не требует. Вируалка будет хорошим решением.

###Elasticsearch кластер для реализации логирования продуктивного веб-приложения - три ноды elasticsearch, два logstash и две ноды kibana;
    Не встречался с Elasticsearch, если я правильно понял: elasticsearch должен хранить логи, Logstash - собирает
    логи, Kibana - мониторинг и визуализация различных метрик. 
    В этом случае, логи хранить буду на виртуальной машине: elasticsearch ставим на виртуалку. Это обеспечит
    надежное хранение данных.
    Средства для обработки и визуализации логов поставлю в docker контейнеры: logstash и kibana в docker. 

###Мониторинг-стек на базе Prometheus и Grafana;
    Насколько я понял, системы Prometheus и Grafana не являются хранилищами данных логов, то можно использовать
    docker.

###MongoDB, как основное хранилище данных для java-приложения;
    Буду использовать физический сервер, как более производительное решение для хранилища.
    Если это будет вспомогательная БД для хранения информации сеанса или документов можно рассмотреть виртуализацию.  

###Gitlab сервер для реализации CI/CD процессов и приватный (закрытый) Docker Registry.
    Хранение данных внутри докера или проводить изменение данных внутри контейнера плохое решение. 
    Т.к. после перезапуска контейнера все данные внутри контейнера теряются.  
    В этом случае буду использовать виртуальную машину.

#3

>cd docker-volume
> 
###centos

>docker pull centos:latest

>docker images

    REPOSITORY                          TAG       IMAGE ID       CREATED       SIZE
    chernyshovnetology/nginx-template   latest    ab7c21f9a2cf   4 hours ago   184MB
    nginx                               latest    87a94228f133   3 weeks ago   133MB
    centos                              latest    5d0da3dc9764   7 weeks ago   231MB

>docker run -v %cd%:/data --name first-centos -d centos ping 8.8.8.8


###debian
>docker pull debian

>docker images

    REPOSITORY                          TAG       IMAGE ID       CREATED       SIZE
    chernyshovnetology/nginx-template   latest    ab7c21f9a2cf   4 hours ago   184MB
    nginx                               latest    87a94228f133   3 weeks ago   133MB
    debian                              latest    f776cfb21b5e   3 weeks ago   124MB
    centos                              latest    5d0da3dc9764   7 weeks ago   231MB 

>docker run -d -v %cd%:/data --name second-debian debian /bin/bash

>docker exec -it first-centos bash

>echo "test" > file_centos

>docker exec -it second-debian bash

    root@3fd191180390:/# cd /data/
    root@3fd191180390:/data# ls -la
    total 4
    drwxrwxrwx 1 root root  512 Nov  6 00:05 .
    drwxr-xr-x 1 root root 4096 Nov  5 23:58 ..
    -rw-r--r-- 1 root root    0 Nov  6 00:03 file_centos.txt
    -rwxrwxrwx 1 root root    0 Nov  6 00:05 file_host.txt


#4

    Использовал Dockerfile
    
    FROM ubuntu:20.04
    
    RUN apt-get update; \
        apt-get install -y gcc libffi-devel python3 epel-release; \
        apt-get install -y python3-pip; \
        apt-get install -y wget; \
        apt-get clean all
    
    RUN pip3 install --upgrade pip; \
        python3 -m pip install ansible; \

>docker>docker build . -t ansible

>docker>docker tag 22884ff6ba87 chernyshovnetology/ansible

>docker push chernyshovnetology/ansible

https://hub.docker.com/repository/docker/chernyshovnetology/ansible