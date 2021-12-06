FROM centos:7

LABEL ElasticSearch Chernyshov

ENV PATH=/usr/lib:$PATH

RUN rpm --import https://artifacts.elastic.co/GPG-KEY-elasticsearch
RUN echo "[elasticsearch]" >>/etc/yum.repos.d/elasticsearch.repo &&\
    echo "name=Elasticsearch repository for 7.x packages" >>/etc/yum.repos.d/elasticsearch.repo &&\
    echo "baseurl=https://artifacts.elastic.co/packages/7.x/yum">>/etc/yum.repos.d/elasticsearch.repo &&\
    echo "gpgcheck=1">>/etc/yum.repos.d/elasticsearch.repo &&\
    echo "gpgkey=https://artifacts.elastic.co/GPG-KEY-elasticsearch">>/etc/yum.repos.d/elasticsearch.repo &&\
    echo "enabled=0">>/etc/yum.repos.d/elasticsearch.repo &&\
    echo "autorefresh=1">>/etc/yum.repos.d/elasticsearch.repo &&\
    echo "type=rpm-md">>/etc/yum.repos.d/elasticsearch.repo

RUN yum install -y --enablerepo=elasticsearch elasticsearch

ADD elasticsearch.yml /etc/elasticsearch/

RUN mkdir /usr/share/elasticsearch/snapshots
RUN chown elasticsearch:elasticsearch /usr/share/elasticsearch/snapshots
RUN mkdir /var/lib/logs
RUN chown elasticsearch:elasticsearch /var/lib/logs
RUN mkdir /var/lib/data
RUN chown elasticsearch:elasticsearch /var/lib/data

USER elasticsearch

CMD ["/usr/sbin/init"]
CMD ["/usr/share/elasticsearch/bin/elasticsearch"]
~
