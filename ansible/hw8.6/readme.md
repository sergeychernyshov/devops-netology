# Подготовка к выполнению

## 1

Создал репозиторий cd my_own_collection

## 2

>git clone https://github.com/ansible/ansible.git

## 3

>cd ansible

## 4

>sudo apt install python3.9-venv
>python3.9 -m venv venv 

## 5

>. venv/bin/activate

## 6

>pip install -r requirements.txt

## 7

>. hacking/env-setup

    Traceback (most recent call last):
      File "setup.py", line 6, in <module>
        from setuptools import find_packages, setup
    ModuleNotFoundError: No module named 'setuptools'
    
    Setting up Ansible to run out of checkout...
    
    PATH=/home/oracle/my_own_collection/ansible/bin:/home/oracle/my_own_collection/venv/bin:/home/oracle/.local/bin:/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin:/usr/games:/usr/local/games:/snap/bin
    PYTHONPATH=/home/oracle/my_own_collection/ansible/test/lib:/home/oracle/my_own_collection/ansible/lib
    MANPATH=/home/oracle/my_own_collection/ansible/docs/man:/usr/local/man:/usr/local/share/man:/usr/share/man
    
    Remember, you may wish to specify your host file with -i
    
    Done!

## 8

>(venv) oracle@oracle-ansible:~/my_own_collection/ansible$ deactivate
oracle@oracle-ansible:~/my_own_collection/ansible$



# Основная часть

## 1

В виртуальном окружении создал новый my_own_module.py файл.
Файл в каталоге /home/oracle/my_own_collection/ansible/lib/ansible/modules

## 2

Наполнил его содержимым, как указано в [статье](https://docs.ansible.com/ansible/latest/dev_guide/developing_modules_general.html#creating-a-module)

В корне создал файл args.json

    {
        "ANSIBLE_MODULE_ARGS": {
            "name": "hello",
            "new": true
        }
    }

>python -m ansible.modules.my_own_module args.json

{"changed": true, "original_message": "hello", "message": "goodbye", "invocation": {"module_args": {"name": "hello", "new": true}}}

## 3 

Переписал код my_own_module.py, чтобы module должен создавал 
текстовый файл на удалённом хосте по пути, определённом 
в параметре path, с содержимым, определённым в 
параметре content 

Изменил файл args.json

    {
        "ANSIBLE_MODULE_ARGS": {
            "path": "/tmp/xxx.txt",
            "content": "mail.ru"
        }
    }

## 4

>python -m ansible.modules.my_own_module args.json

    {"changed": true, "original_message": "mail.ru", "message": "file created", "invocation": {"module_args": {"path": "/tmp/xxx.txt", "content": "mail.ru"}}}

>python -m ansible.modules.my_own_module args.json

    {"changed": false, "original_message": "mail.ru", "message": "file exists", "invocation": {"module_args": {"path": "/tmp/xxx.txt", "content": "mail.ru"}}}

## 5 

Создал простой playbook

    ---
    - name: Assert create file module
      hosts: localhost
      tasks:
      - name: Create file
        my_own_module:
          path: /tmp/xxx5.txt
          content: "hello netology"

## 6

>ansible-playbook playbook/site.yml

    [WARNING]: You are running the development version of Ansible. You should only run Ansible from "devel" if you are modifying the Ansible engine, or trying out features under development. This is a rapidly changing source of code and can
    become unstable at any point.
    [WARNING]: No inventory was parsed, only implicit localhost is available
    [WARNING]: provided hosts list is empty, only localhost is available. Note that the implicit localhost does not match 'all'
    
    PLAY [Assert create file module] ************************************************************************************************************************************************************************************************************
    
    TASK [Gathering Facts] **********************************************************************************************************************************************************************************************************************
    ok: [localhost]
    
    TASK [Create file] **************************************************************************************************************************************************************************************************************************
    changed: [localhost]
    
    PLAY RECAP **********************************************************************************************************************************************************************************************************************************
    localhost                  : ok=2    changed=1    unreachable=0    failed=0    skipped=0    rescued=0    ignored=0

>ansible-playbook playbook/site.yml

    [WARNING]: You are running the development version of Ansible. You should only run Ansible from "devel" if you are modifying the Ansible engine, or trying out features under development. This is a rapidly changing source of code and can
    become unstable at any point.
    [WARNING]: No inventory was parsed, only implicit localhost is available
    [WARNING]: provided hosts list is empty, only localhost is available. Note that the implicit localhost does not match 'all'
    
    PLAY [Assert create file module] ************************************************************************************************************************************************************************************************************
    
    TASK [Gathering Facts] **********************************************************************************************************************************************************************************************************************
    ok: [localhost]
    
    TASK [Create file] **************************************************************************************************************************************************************************************************************************
    ok: [localhost]
    
    PLAY RECAP **********************************************************************************************************************************************************************************************************************************
    localhost                  : ok=2    changed=0    unreachable=0    failed=0    skipped=0    rescued=0    ignored=0

playbook на идемпотентен

## 7

(venv) oracle@oracle-ansible:~/my_own_collection2/ansible$ deactivate
oracle@oracle-ansible:~/my_own_collection2/ansible$

Вышел из окружения

## 8

>ansible-galaxy collection init my_own_namespace.files

    [WARNING]: You are running the development version of Ansible. You should only
    run Ansible from "devel" if you are modifying the Ansible engine, or trying out
    features under development. This is a rapidly changing source of code and can
    become unstable at any point.
    - Collection my_own_namespace.files was created successfully

Создал коллекцию с названием files

## 9

>cd my_own_namespace/files/plugins/
>mkdir modules
>cp /home/oracle/my_own_collection/ansible/lib/ansible/modules/my_own_module.py /home/oracle/my_own_collection/ansible/my_own_namespace/files/plugins/modules//my_own_module.py

Коллекция скопирована

## 10

>ansible-galaxy init my_onw_module_role

    [WARNING]: You are running the development version of Ansible. You should only run Ansible from "devel" if you are
    modifying the Ansible engine, or trying out features under development. This is a rapidly changing source of code
    and can become unstable at any point.
    - Role my_onw_module_role was created successfully

Перенес Single task playbook в single task role

## 11

Создал site_role.yml

    ---
    - name: Assert create file module
      hosts: localhost
      roles: 
        - role: my_onw_module_role

>ansible-playbook playbook/site_role.yml

    [WARNING]: You are running the development version of Ansible. You should only run Ansible from "devel" if you are
    modifying the Ansible engine, or trying out features under development. This is a rapidly changing source of code
    and can become unstable at any point.
    [WARNING]: No inventory was parsed, only implicit localhost is available
    [WARNING]: provided hosts list is empty, only localhost is available. Note that the implicit localhost does not
    match 'all'
    
    PLAY [Assert create file module] ************************************************************************************
    
    TASK [Gathering Facts] **********************************************************************************************
    ok: [localhost]
    
    TASK [my_onw_module_role : run create file module] ******************************************************************
    changed: [localhost]
    
    PLAY RECAP **********************************************************************************************************
    localhost                  : ok=2    changed=1    unreachable=0    failed=0    skipped=0    rescued=0    ignored=0

## 12

Сохранил в свой репозиторий, поставил тег 1.0.0

## 13

Зашел в корень коллекции

>cd my_own_namespace/files/
>ansible-galaxy collection build

    [WARNING]: You are running the development version of Ansible. You should only run Ansible from "devel" if you are
    modifying the Ansible engine, or trying out features under development. This is a rapidly changing source of code
    and can become unstable at any point.
    Created collection for my_own_namespace.files at /home/oracle/my_own_collection/ansible/my_own_namespace/files/my_own_namespace-files-1.0.0.tar.gz

## 14 

Создал директорию files-copy. Скопировал содержимое из files

## 15

>ansible-galaxy collection install my_own_namespace-files-1.0.0.tar.gz
 
    [WARNING]: You are running the development version of Ansible. You should only run Ansible from "devel" if you are
    modifying the Ansible engine, or trying out features under development. This is a rapidly changing source of code
    and can become unstable at any point.
    Starting galaxy collection install process
    Process install dependency map
    Starting collection install process
    Installing 'my_own_namespace.files:1.0.0' to '/home/oracle/.ansible/collections/ansible_collections/my_own_namespace/files'
    my_own_namespace.files:1.0.0 was installed successfully

## 16

>ansible-playbook playbook/site_role.yml

    [WARNING]: You are running the development version of Ansible. You should only run Ansible from "devel" if you are
    modifying the Ansible engine, or trying out features under development. This is a rapidly changing source of code
    and can become unstable at any point.
    [WARNING]: No inventory was parsed, only implicit localhost is available
    [WARNING]: provided hosts list is empty, only localhost is available. Note that the implicit localhost does not
    match 'all'
    
    PLAY [Assert create file module] ************************************************************************************
    
    TASK [Gathering Facts] **********************************************************************************************
    ok: [localhost]
    
    TASK [my_onw_module_role : run create file module] ******************************************************************
    ok: [localhost]
    
    PLAY RECAP **********************************************************************************************************
    localhost                  : ok=2    changed=0    unreachable=0    failed=0    skipped=0    rescued=0    ignored=0

Проверил, все работает.

## 17

[Ссылка на репозиторий](https://github.com/sergeychernyshov/my_own_collection/tree/1.0.0) версия 1.0.0