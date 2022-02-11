#1 Установите ansible версии 2.10 или выше. 
Для версии 2.10 и выше требуется Python 3.8 и выше 
Установка Python 3.9

>sudo apt update
>sudo apt install software-properties-common

Библиотеки для Python https://wiki.python.org/moin/UbuntuInstall 
>sudo apt install build-essential zlib1g-dev libncurses5-dev libgdbm-dev libnss3-dev libssl-dev libreadline-dev libffi-dev libsqlite3-dev wget libbz2-dev
>wget https://www.python.org/ftp/python/3.9.9/Python-3.9.9.tgz
>tar -xf Python-3.9.9.tgz
>cd Python-3.9.9 
>./configure --enable-optimizations
>make -j 12
>sudo make altinstall
>python3.9 --version
    
    Python 3.9.9
 
>sudo apt install python3-pip
>pip --version
 
    pip 20.0.2 from /usr/lib/python3/dist-packages/pip (python 3.8)

>pip3 --version

    pip 20.0.2 from /usr/lib/python3/dist-packages/pip (python 3.8)

>pip3 install ansible

>ansible --version
ansible [core 2.12.2]
  config file = None
  configured module search path = ['/home/oracle/.ansible/plugins/modules', '/usr/share/ansible/plugins/modules']
  ansible python module location = /home/oracle/.local/lib/python3.8/site-packages/ansible
  ansible collection location = /home/oracle/.ansible/collections:/usr/share/ansible/collections
  executable location = /home/oracle/.local/bin/ansible
  python version = 3.8.10 (default, Nov 26 2021, 20:14:08) [GCC 9.3.0]
  jinja version = 3.0.3
  libyaml = True

#2 готово
#3 готово

*******************************************************
#1 
>ansible-playbook site.yml -i inventory/test.yml

  PLAY [Print os facts] ***********************************************************************************************************************************************************************************************************************
  
  TASK [Gathering Facts] **********************************************************************************************************************************************************************************************************************
  ok: [localhost]
  
  TASK [Print OS] *****************************************************************************************************************************************************************************************************************************
  ok: [localhost] => {
      "msg": "Ubuntu"
  }
  
  TASK [Print fact] ***************************************************************************************************************************************************************************************************************************
  ok: [localhost] => {
      "msg": 12
  }
  
  PLAY RECAP **********************************************************************************************************************************************************************************************************************************
  localhost                  : ok=3    changed=0    unreachable=0    failed=0    skipped=0    rescued=0    ignored=0

Значение переменной some_fact = 12 

#2
В файле [playbook/group_vars/all/examp.yml](https://github.com/sergeychernyshov/ansible-netology/blob/main/group_vars/all/examp.yml) установил значение

some_fact: all default fact

>ansible-playbook site.yml -i inventory/test.yml

PLAY [Print os facts] ***********************************************************************************************************************************************************************************************************************

TASK [Gathering Facts] **********************************************************************************************************************************************************************************************************************
ok: [localhost]

TASK [Print OS] *****************************************************************************************************************************************************************************************************************************
ok: [localhost] => {
    "msg": "Ubuntu"
}

TASK [Print fact] ***************************************************************************************************************************************************************************************************************************
ok: [localhost] => {
    "msg": "all default fact"
}

PLAY RECAP **********************************************************************************************************************************************************************************************************************************
localhost                  : ok=3    changed=0    unreachable=0    failed=0    skipped=0    rescued=0    ignored=0

#3 
Установил Docker 
запускаю контейнеры c python:3.9
>docker run --rm -d -i --name=centos7 python:3.9 /bin/bash
>docker run --rm -d -i --name=ubuntu python:3.9 /bin/bash


>docker ps

CONTAINER ID   IMAGE        COMMAND       CREATED              STATUS              PORTS     NAMES
20b905047ed1   python:3.9   "/bin/bash"   3 seconds ago        Up 3 seconds                  ubuntu
a353d57a382f   python:3.9   "/bin/bash"   About a minute ago   Up About a minute             centos7

#4
>ansible-playbook site.yml -i inventory/prod.yml

PLAY [Print os facts] ***********************************************************************************************************************************************************************************************************************

TASK [Gathering Facts] **********************************************************************************************************************************************************************************************************************
ok: [ubuntu]
ok: [centos7]

TASK [Print OS] *****************************************************************************************************************************************************************************************************************************
ok: [centos7] => {
    "msg": "Debian"
}
ok: [ubuntu] => {
    "msg": "Debian"
}

TASK [Print fact] ***************************************************************************************************************************************************************************************************************************
ok: [centos7] => {
    "msg": "el"
}
ok: [ubuntu] => {
    "msg": "deb"
}

PLAY RECAP **********************************************************************************************************************************************************************************************************************************
centos7                    : ok=3    changed=0    unreachable=0    failed=0    skipped=0    rescued=0    ignored=0
ubuntu                     : ok=3    changed=0    unreachable=0    failed=0    skipped=0    rescued=0    ignored=0


для centos7 some_fact = el
для ubuntu some_fact = deb

#5
Не совсем понял задание.

Добавил файлы [var_deb.yml](https://github.com/sergeychernyshov/ansible-netology/blob/main/group_vars/deb/examp.yml) и 
[var_deb.yml](https://github.com/sergeychernyshov/ansible-netology/blob/main/group_vars/el/examp.yml) 
Эти файлы идут после examp.yml, по этому значения из добавленных файлов читаются позже и перезаписывают значения из examp.yml.

#6
>ansible-playbook site.yml -i inventory/prod.yml

PLAY [Print os facts] ***********************************************************************************************

TASK [Gathering Facts] **********************************************************************************************
ok: [centos7]
ok: [ubuntu]

TASK [Print OS] *****************************************************************************************************
ok: [centos7] => {
    "msg": "Debian"
}
ok: [ubuntu] => {
    "msg": "Debian"
}

TASK [Print fact] ***************************************************************************************************
ok: [centos7] => {
    "msg": "el default fact"
}
ok: [ubuntu] => {
    "msg": "deb default fact"
}

PLAY RECAP **********************************************************************************************************
centos7                    : ok=3    changed=0    unreachable=0    failed=0    skipped=0    rescued=0    ignored=0
ubuntu                     : ok=3    changed=0    unreachable=0    failed=0    skipped=0    rescued=0    ignored=0

#7
>ansible-vault encrypt el/examp.yml
>ansible-vault encrypt el/var_el.yml
>ansible-vault encrypt deb/examp.yml
>ansible-vault encrypt deb/var_deb.yml

#8

>ansible-playbook site.yml -i inventory/prod.yml --ask-vault-password
Vault password:

PLAY [Print os facts] ***********************************************************************************************

TASK [Gathering Facts] **********************************************************************************************
ok: [ubuntu]
ok: [centos7]

TASK [Print OS] *****************************************************************************************************
ok: [centos7] => {
    "msg": "Debian"
}
ok: [ubuntu] => {
    "msg": "Debian"
}

TASK [Print fact] ***************************************************************************************************
ok: [centos7] => {
    "msg": "el default fact"
}
ok: [ubuntu] => {
    "msg": "deb default fact"
}

PLAY RECAP **********************************************************************************************************
centos7                    : ok=3    changed=0    unreachable=0    failed=0    skipped=0    rescued=0    ignored=0
ubuntu                     : ok=3    changed=0    unreachable=0    failed=0    skipped=0    rescued=0    ignored=0

#9
Список плагинов для подключения
>ansible-doc -t connection -l

ansible.netcommon.httpapi      Use httpapi to run command on network applia...
ansible.netcommon.libssh       (Tech preview) Run tasks using libssh for ss...
ansible.netcommon.napalm       Provides persistent connection using NAPALM
ansible.netcommon.netconf      Provides a persistent connection using the n...
ansible.netcommon.network_cli  Use network_cli to run command on network ap...
ansible.netcommon.persistent   Use a persistent unix socket for connection
community.aws.aws_ssm          execute via AWS Systems Manager
community.docker.docker        Run tasks in docker containers
community.docker.docker_api    Run tasks in docker containers
community.docker.nsenter       execute on host running controller container
community.general.chroot       Interact with local chroot
community.general.funcd        Use funcd to connect to target
community.general.iocage       Run tasks in iocage jails
community.general.jail         Run tasks in jails
community.general.lxc          Run tasks in lxc containers via lxc python l...
community.general.lxd          Run tasks in lxc containers via lxc CLI
community.general.qubes        Interact with an existing QubesOS AppVM
community.general.saltstack    Allow ansible to piggyback on salt minions
community.general.zone         Run tasks in a zone instance
community.libvirt.libvirt_lxc  Run tasks in lxc containers via libvirt
community.libvirt.libvirt_qemu Run tasks on libvirt/qemu virtual machines
community.okd.oc               Execute tasks in pods running on OpenShift
community.vmware.vmware_tools  Execute tasks inside a VM via VMware Tools
:...skipping...
ansible.netcommon.httpapi      Use httpapi to run command on network applia...
ansible.netcommon.libssh       (Tech preview) Run tasks using libssh for ss...
ansible.netcommon.napalm       Provides persistent connection using NAPALM
ansible.netcommon.netconf      Provides a persistent connection using the n...
ansible.netcommon.network_cli  Use network_cli to run command on network ap...
ansible.netcommon.persistent   Use a persistent unix socket for connection
community.aws.aws_ssm          execute via AWS Systems Manager
community.docker.docker        Run tasks in docker containers
community.docker.docker_api    Run tasks in docker containers
community.docker.nsenter       execute on host running controller container
community.general.chroot       Interact with local chroot
community.general.funcd        Use funcd to connect to target
community.general.iocage       Run tasks in iocage jails
community.general.jail         Run tasks in jails
community.general.lxc          Run tasks in lxc containers via lxc python l...
community.general.lxd          Run tasks in lxc containers via lxc CLI
community.general.qubes        Interact with an existing QubesOS AppVM
community.general.saltstack    Allow ansible to piggyback on salt minions
community.general.zone         Run tasks in a zone instance
community.libvirt.libvirt_lxc  Run tasks in lxc containers via libvirt
community.libvirt.libvirt_qemu Run tasks on libvirt/qemu virtual machines
community.okd.oc               Execute tasks in pods running on OpenShift
community.vmware.vmware_tools  Execute tasks inside a VM via VMware Tools
containers.podman.buildah      Interact with an existing buildah container
containers.podman.podman       Interact with an existing podman container
kubernetes.core.kubectl        Execute tasks in pods running on Kubernetes
local                          execute on controller
paramiko_ssh                   Run tasks via python ssh (paramiko)
psrp                           Run tasks over Microsoft PowerShell Remoting...
ssh                            connect via SSH client binary
winrm                          Run tasks over Microsoft's WinRM


local(execute on controller) - Этот подойдет для выполнения команд на локальном хосте.

#10
Добавил в inventory 

  local:
    hosts:
      localhost:
        ansible_connection: local

#11

ansible-playbook site.yml -i inventory/prod.yml --ask-vault-password
Vault password:

PLAY [Print os facts] ***********************************************************************************************************************************************************************************************************************

TASK [Gathering Facts] **********************************************************************************************************************************************************************************************************************
ok: [localhost]
ok: [centos7]
ok: [ubuntu]

TASK [Print OS] *****************************************************************************************************************************************************************************************************************************
ok: [localhost] => {
    "msg": "Ubuntu"
}
ok: [ubuntu] => {
    "msg": "Debian"
}
ok: [centos7] => {
    "msg": "Debian"
}

TASK [Print fact] ***************************************************************************************************************************************************************************************************************************
ok: [localhost] => {
    "msg": "all default fact"
}
ok: [centos7] => {
    "msg": "el default fact"
}
ok: [ubuntu] => {
    "msg": "deb default fact"
}

PLAY RECAP **********************************************************************************************************************************************************************************************************************************
centos7                    : ok=3    changed=0    unreachable=0    failed=0    skipped=0    rescued=0    ignored=0
localhost                  : ok=3    changed=0    unreachable=0    failed=0    skipped=0    rescued=0    ignored=0
ubuntu                     : ok=3    changed=0    unreachable=0    failed=0    skipped=0    rescued=0    ignored=0

#12
Этот файл

