#Подготовка к выполнению

##1

Создал репозитории 
[kibana-role](https://github.com/sergeychernyshov/kibana-role) 
[filebeat-role](https://github.com/sergeychernyshov/filebeat-role)

##2 

Готово
 
#Основная часть

##1

Создал файл requirements.yml в корне playbook

##2

Проверяю сертификат

>ssh -T git@github.com
Warning: Permanently added the ECDSA host key for IP address '140.82.121.3' to the list of known hosts.
Hi sergeychernyshov! You've successfully authenticated, but GitHub does not provide shell access.

>ansible-galaxy install -r requirements.yml

    Starting galaxy role install process
    - extracting elastic to /home/oracle/.ansible/roles/elastic
    - elastic (2.1.4) was installed successfully

------------------------------------------------------------
Не смог использовать в своем Playbook эту роль. Создал свою

>ansible-galaxy role init elastic-role

    - Role elastic-role was created successfully

##3

>ansible-galaxy role init kibana-role

    - Role kibana-role was created successfully

>tree

    .
    ├── files
    │   └── jdk-11.0.14_linux-x64_bin.tar.gz
    ├── inventory
    │   └── prod
    │       ├── group_vars
    │       │   ├── app.yml
    │       │   ├── elasticsearch.yml
    │       │   └── kibana.yml
    │       └── hosts.yml
    ├── kibana-role
    │   ├── defaults
    │   │   └── main.yml
    │   ├── files
    │   ├── handlers
    │   │   └── main.yml
    │   ├── meta
    │   │   └── main.yml
    │   ├── README.md
    │   ├── tasks
    │   │   └── main.yml
    │   ├── templates
    │   ├── tests
    │   │   ├── inventory
    │   │   └── test.yml
    │   └── vars
    │       └── main.yml
    ├── requirements.yml
    ├── site.yml
    └── templates
        ├── elasticsearch.yml.j2
        ├── filebeat.yml.j2

Структура каталогов для роли kibana-role создана

##4

Разнес переменные между vars и default

##5

Перенес шаблоны в templates

##6

>ansible-galaxy role init filebeat-role
    
    - Role filebeat-role was created successfully

##7

Разнес переменные между vars и default

##8

Перенес шаблоны в templates

##9

Описал в README.md роли и их параметры

##10

Выложил все roles в репозитории. Проставил тэги.
 
##11

Добавил файл requirements.yml

>ansible-galaxy install -r requirements.yml

    Starting galaxy role install process
    - extracting kibana-role to /home/oracle/.ansible/roles/kibana-role
    - kibana-role (1.0.0) was installed successfully
    - extracting elastic-role to /home/oracle/.ansible/roles/elastic-role
    - elastic-role (1.0.0) was installed successfully
    - extracting filebeat-role to /home/oracle/.ansible/roles/filebeat-role
    - filebeat-role (1.0.0) was installed successfully

##12

Переработал playbook с использованием ролей

##13

Выложил [playbook](https://github.com/sergeychernyshov/ansible-netology/tree/main/hw4)

##14

Роли 

[elastic-role](https://github.com/sergeychernyshov/elastic-role)

[kibana-role](https://github.com/sergeychernyshov/kibana-role)

[filebeat-role](https://github.com/sergeychernyshov/filebeat-role)
