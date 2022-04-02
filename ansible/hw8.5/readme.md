# Подготовка к выполнению

## 1

Установил molecule

>pip3 install "molecule==3.4.0"
>molecule --version

    molecule 3.4.0 using python 3.8
        ansible:2.12.2
        delegated:3.4.0 from molecule 

## 2
Собрал локальный образ

# Основная часть

## Molecule

### 1

>cd .ansible/roles/elastic-role/

>molecule test

    CRITICAL Failed to pre-validate.

    {'driver': [{'name': ['unallowed value docker']}]}

>pip3 install molecule-docker
>molecule test

-------------------------

>molecule init scenario docker --driver-name docker

    INFO     Initializing new scenario docker...
    INFO     Initialized scenario in /home/oracle/ansible/playbook_hw5/test_elastic_role/molecule/docker successfully.

>molecule test

    PLAY RECAP *********************************************************************
    localhost                  : ok=2    changed=2    unreachable=0    failed=0    skipped=1    rescued=0    ignored=0
    
    INFO     Pruning extra files from scenario ephemeral directory
 
### 2

Переходим в каталог с ролью kibana-role

>cd /home/oracle/.ansible/roles/kibana-role

>molecule init scenario --driver-name docker

    INFO     Initializing new scenario default...
    INFO     Initialized scenario in /home/oracle/.ansible/roles/kibana-role/molecule/default successfully.

>molecule init scenario -d docker kibana-role


# TOX

Yстановка TOX

>pip3 install tox
    
    Successfully installed distlib-0.3.4 filelock-3.6.0 platformdirs-2.5.1 py-1.11.0 toml-0.10.2 tox-3.24.5 virtualenv-20.13.4

# 1 

Запустил докер контейнер

 docker run --name netologyHW85 --privileged=True -v /home/oracle/.ansible/roles/kibana-role:/opt/kibana-role -v /home/oracle/.ansible/roles/kibana-role/molecule/alternative/containers.conf:/etc/containers/containers.conf -w /opt/kibana-role -it ansiblemolecula bash

# 2