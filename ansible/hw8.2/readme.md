#1 Приготовьте свой собственный inventory файл prod.yml.

>docker run --rm -d -i --name=elastic python:3.9 /bin/bash
>docker run --rm -d -i --name=kibana python:3.9 /bin/bash

Создал файл inventory 

  ---
    elasticsearch:
      hosts:
        elastic:
          ansible_connection: docker
    kibana:
      hosts:
        kibana:
          ansible_connection: docker      


#2 Допишите playbook: нужно сделать ещё один play, который устанавливает и настраивает kibana.

Дописал файл    [site.yml](https://github.com/sergeychernyshov/ansible-netology/blob/main/hw2/site.yml)

Создал template [kib.sh.j2](https://github.com/sergeychernyshov/ansible-netology/blob/main/hw2/templates/kib.sh.j2)

Создал group_vars  [kibana/vars.yml](https://github.com/sergeychernyshov/ansible-netology/blob/main/hw2/group_vars/kibana/vars.yml)

#3 При создании tasks рекомендую использовать модули: get_url, template, unarchive, file.

Использовал модули get_url, template, unarchive, file

#4 Tasks должны: скачать нужной версии дистрибутив, выполнить распаковку в выбранную директорию, сгенерировать конфигурацию с параметрами.

    - name: Install kibana
        hosts: kibana
        tasks:
          - name: Upload tar.gz kibana from remote URL
            get_url:
              url: "https://artifacts.elastic.co/downloads/kibana/kibana-{{ kibana_version }}-linux-x86_64.tar.gz"
              dest: "/tmp/kibana-{{ elastic_version }}-linux-x86_64.tar.gz"
              mode: 0755
              timeout: 60
              force: true
              validate_certs: false
            register: get_kibana
            until: get_kibana is succeeded
            tags: kibana
          - name: Create directrory for kibana
            file:
              state: directory
              path: "{{ kibana_home }}"
            tags: kibana
          - name: Extract Kibana in the installation directory
            #become: true
            unarchive:
              copy: false
              src: "/tmp/kibana-{{ kibana_version }}-linux-x86_64.tar.gz"
              dest: "{{ kibana_home }}"
              extra_opts: [--strip-components=1]
              creates: "{{ kibana_home }}/bin/kibana"
            tags:
              - skip_ansible_lint
              - kibana
          - name: Set environment Kibana
            #become: true
            template:
              src: templates/kib.sh.j2
              dest: /etc/profile.d/kib.sh
            tags: kibana

#5 Запустите ansible-lint site.yml и исправьте ошибки, если они есть.

Устанавливаю ansible-lint

>pip3 install "ansible-lint"
>ansible-lint site.yml

Ошибок нет

#6 Попробуйте запустить playbook на этом окружении с флагом --check.

>ansible-playbook -i inventory/prod.yml site.yml --check

    [WARNING]: Found both group and host with same name: kibana
    [WARNING]: While constructing a mapping from /home/oracle/ansible/playbook_hw2/site.yml, line 38, column 3, found
    a duplicate dict key (tasks). Using last defined value only.
    
    PLAY [Install Java] ***********************************************************************************************
    
    TASK [Gathering Facts] ********************************************************************************************
    ok: [kibana]
    ok: [elastic]
    
    TASK [Set facts for Java 11 vars] *********************************************************************************
    ok: [elastic]
    ok: [kibana]
    
    TASK [Upload .tar.gz file containing binaries from local storage] *************************************************
    changed: [kibana]
    changed: [elastic]
    
    TASK [Ensure installation dir exists] *****************************************************************************
    changed: [kibana]
    changed: [elastic]
    
    TASK [Extract java in the installation directory] *****************************************************************
    An exception occurred during task execution. To see the full traceback, use -vvv. The error was: NoneType: None
    fatal: [elastic]: FAILED! => {"changed": false, "msg": "dest '/opt/jdk/11.0.14' must be an existing dir"}
    An exception occurred during task execution. To see the full traceback, use -vvv. The error was: NoneType: None
    fatal: [kibana]: FAILED! => {"changed": false, "msg": "dest '/opt/jdk/11.0.14' must be an existing dir"}
    
    PLAY RECAP ********************************************************************************************************
    elastic                    : ok=4    changed=2    unreachable=0    failed=1    skipped=0    rescued=0    ignored=0                                                                                                                           
    kibana                     : ok=4    changed=2    unreachable=0    failed=1    skipped=0    rescued=0    ignored=0    

#7 Запустите playbook на prod.yml окружении с флагом --diff. Убедитесь, что изменения на системе произведены.

>ansible-playbook -i inventory/prod.yml site.yml --diff

    [WARNING]: Found both group and host with same name: kibana
    
    PLAY [Install Java] *************************************************************************************************************************************************************************************************************************
    
    TASK [Gathering Facts] **********************************************************************************************************************************************************************************************************************
    ok: [elastic]
    ok: [kibana]
    
    TASK [Set facts for Java 11 vars] ***********************************************************************************************************************************************************************************************************
    ok: [elastic]
    ok: [kibana]
    
    TASK [Upload .tar.gz file containing binaries from local storage] ***************************************************************************************************************************************************************************
    ok: [kibana]
    ok: [elastic]
    
    TASK [Ensure installation dir exists] *******************************************************************************************************************************************************************************************************
    ok: [kibana]
    ok: [elastic]
    
    TASK [Extract java in the installation directory] *******************************************************************************************************************************************************************************************
    skipping: [elastic]
    skipping: [kibana]
    
    TASK [Export environment variables] *********************************************************************************************************************************************************************************************************
    ok: [kibana]
    ok: [elastic]
    
    PLAY [Install Elasticsearch] ****************************************************************************************************************************************************************************************************************
    
    TASK [Gathering Facts] **********************************************************************************************************************************************************************************************************************
    ok: [elastic]
    
    TASK [Upload tar.gz Elasticsearch from remote URL] ******************************************************************************************************************************************************************************************
    changed: [elastic]
    
    TASK [Create directrory for Elasticsearch] **************************************************************************************************************************************************************************************************
    --- before
    +++ after
    @@ -1,4 +1,4 @@
     {
         "path": "/opt/elastic/7.10.1",
    -    "state": "absent"
    +    "state": "directory"
     }
    
    changed: [elastic]
    
    TASK [Extract Elasticsearch in the installation directory] **********************************************************************************************************************************************************************************
    changed: [elastic]
    
    TASK [Set environment Elastic] **************************************************************************************************************************************************************************************************************
    --- before
    +++ after: /home/oracle/.ansible/tmp/ansible-local-1317211m2vol6/tmpewb3v9oq/elk.sh.j2
    @@ -0,0 +1,5 @@
    +# Warning: This file is Ansible Managed, manual changes will be overwritten on next playbook run.
    +#!/usr/bin/env bash
    +
    +export ES_HOME=/opt/elastic/7.10.1
    +export PATH=$PATH:$ES_HOME/bin
    \ No newline at end of file
    
    changed: [elastic]
    
    PLAY [Install Kibana] ***********************************************************************************************************************************************************************************************************************
    
    TASK [Gathering Facts] **********************************************************************************************************************************************************************************************************************
    ok: [kibana]
    
    TASK [Upload tar.gz Kibana from remote URL] *************************************************************************************************************************************************************************************************
    changed: [kibana]
    
    TASK [Create directrory for Kibana (/opt/kibana/7.10.1)] ************************************************************************************************************************************************************************************
    --- before
    +++ after
    @@ -1,4 +1,4 @@
     {
         "path": "/opt/kibana/7.10.1",
    -    "state": "absent"
    +    "state": "directory"
     }
    
    changed: [kibana]
    
    TASK [Extract Kibana in the installation directory] *****************************************************************************************************************************************************************************************
    changed: [kibana]
    
    TASK [Set environment Kibana] ***************************************************************************************************************************************************************************************************************
    --- before
    +++ after: /home/oracle/.ansible/tmp/ansible-local-1317211m2vol6/tmpxauxqzrd/kib.sh.j2
    @@ -0,0 +1,5 @@
    +# Warning: This file is Ansible Managed, manual changes will be overwritten on next playbook run.
    +#!/usr/bin/env bash
    +
    +export KIBANA_HOME=/opt/kibana/7.10.1
    +export PATH=$PATH:$KIBANA_HOME/bin
    \ No newline at end of file
    
    changed: [kibana]
    
    PLAY RECAP **********************************************************************************************************************************************************************************************************************************
    elastic                    : ok=10   changed=4    unreachable=0    failed=0    skipped=1    rescued=0    ignored=0
    kibana                     : ok=10   changed=4    unreachable=0    failed=0    skipped=1    rescued=0    ignored=0

#8 Повторно запустите playbook с флагом --diff и убедитесь, что playbook идемпотентен.

> ansible-playbook -i inventory/prod.yml site.yml --diff
  
    [WARNING]: Found both group and host with same name: kibana
    
    PLAY [Install Java] *************************************************************************************************************************************************************************************************************************
    
    TASK [Gathering Facts] **********************************************************************************************************************************************************************************************************************
    ok: [kibana]
    ok: [elastic]
    
    TASK [Set facts for Java 11 vars] ***********************************************************************************************************************************************************************************************************
    ok: [elastic]
    ok: [kibana]
    
    TASK [Upload .tar.gz file containing binaries from local storage] ***************************************************************************************************************************************************************************
    ok: [elastic]
    ok: [kibana]
    
    TASK [Ensure installation dir exists] *******************************************************************************************************************************************************************************************************
    ok: [elastic]
    ok: [kibana]
    
    TASK [Extract java in the installation directory] *******************************************************************************************************************************************************************************************
    skipping: [elastic]
    skipping: [kibana]
    
    TASK [Export environment variables] *********************************************************************************************************************************************************************************************************
    ok: [kibana]
    ok: [elastic]
    
    PLAY [Install Elasticsearch] ****************************************************************************************************************************************************************************************************************
    
    TASK [Gathering Facts] **********************************************************************************************************************************************************************************************************************
    ok: [elastic]
    
    TASK [Upload tar.gz Elasticsearch from remote URL] ******************************************************************************************************************************************************************************************
    ok: [elastic]
    
    TASK [Create directrory for Elasticsearch] **************************************************************************************************************************************************************************************************
    ok: [elastic]
    
    TASK [Extract Elasticsearch in the installation directory] **********************************************************************************************************************************************************************************
    skipping: [elastic]
    
    TASK [Set environment Elastic] **************************************************************************************************************************************************************************************************************
    ok: [elastic]
    
    PLAY [Install Kibana] ***********************************************************************************************************************************************************************************************************************
    
    TASK [Gathering Facts] **********************************************************************************************************************************************************************************************************************
    ok: [kibana]
    
    TASK [Upload tar.gz Kibana from remote URL] *************************************************************************************************************************************************************************************************
    ok: [kibana]
    
    TASK [Create directrory for Kibana (/opt/kibana/7.10.1)] ************************************************************************************************************************************************************************************
    ok: [kibana]
    
    TASK [Extract Kibana in the installation directory] *****************************************************************************************************************************************************************************************
    skipping: [kibana]
    
    TASK [Set environment Kibana] ***************************************************************************************************************************************************************************************************************
    ok: [kibana]
    
    PLAY RECAP **********************************************************************************************************************************************************************************************************************************
    elastic                    : ok=9    changed=0    unreachable=0    failed=0    skipped=2    rescued=0    ignored=0
    kibana                     : ok=9    changed=0    unreachable=0    failed=0    skipped=2    rescued=0    ignored=0


#9 Подготовьте README.md файл по своему playbook. В нём должно быть описано: что делает playbook, какие у него есть параметры и теги.

Файл [README.md](https://github.com/sergeychernyshov/ansible-netology/blob/main/hw2/README.md)

#10 Готовый playbook выложите в свой репозиторий, в ответ предоставьте ссылку на него.

Ссылка на репозиторий [Playbook](https://github.com/sergeychernyshov/ansible-netology/tree/main/hw2)