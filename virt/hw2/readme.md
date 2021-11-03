#1

####Основные преимущества IaaC:
    
-упрощается процесс разработки. Появляется возможность быстро создавать dev среды на стенде и на локальных ПК.
*** мне как разработчику хранилищ данных этот момент очень тяжело принять, т.к. много терабайтную БД не 
просто перенести куда-нибудь, а разрабатывать на БД без значительных объемов данных не просто.
Предполагаю, что для больших БД в проектах существуют другие решения. Например, содержать копии продуктивной
БД на отдельных стендах.

-автоматизация переноса конфигураций позволяет содержать среды DEV,TEST,PROD в идентичном состоянии.

-легче находить ошибки в системе при помощи автоматизации тестирования. это позволяет ускорить процесс
    разработки.

    
####Главное преимущество IaaC является идемпотентность. 

В презентации указано, что это сво-во при котором
повторное выполнение операции получается тот же результат.
***мне доводилось слышать иное определение идемпотентности в программировании: 
Идемпотентность в информатике — действие, многократное повторение которого эквивалентно однократному.
Наверное, это больше относится к программированию, чем к системе.

#2

Ansible использует SSH инфраструктуру и не требует дополнительной установки специального окружения.
Также Ansible имеет простой декларативный язык описания. 



Pull-модель подходит для организации обновления системных компонентов на большом количестве кластеров
Push-модель подходит для выката приложений   

Модель Pull плюсы:
-внешний клиент не имеет прав на внесение изменений в кластер, все обновления накатываются изнутри;
-возможность обновления версий ПО из репозитория;
-инструменты могут быть распределены по разным пространствам имен с разными репозиториями Git и правами доступа;

Модель Pull минусы:
-сложно работы с ключами;

Модель Push плюсы:
-обновления версий контейнеров могут быть инициированы пайплайном сборки;

Модель Push минусы:
-данные для доступа к кластеру находятся внутри системы сборки;

#3

    apt-get update
    
    apt install virtualbox
    apt-get install virtualbox-dkms
    dpkg-reconfigure virtualbox-dkms 
    dpkg-reconfigure virtualbox
    
    vboxmanage --version
    6.1.26_Ubuntur145957
    
    
    install curl
    curl -O https://releases.hashicorp.com/vagrant/2.2.18/vagrant_2.2.18_x86_64.deb
    
    apt install ./vagrant_2.2.18_x86_64.deb
    
    vagrant --version
    Vagrant 2.2.18
    
    apt install ansible
    ansible --version
    ansible 2.9.6
      config file = /etc/ansible/ansible.cfg
      configured module search path = ['/root/.ansible/plugins/modules', '/usr/share/ansible/plugins/modules']
      ansible python module location = /usr/lib/python3/dist-packages/ansible
      executable location = /usr/bin/ansible
      python version = 3.8.10 (default, Jun  2 2021, 10:49:15) [GCC 9.4.0]


#4

### Запускаю создание ВМ

    vagrant up
    Bringing machine 'server1.netology' up with 'virtualbox' provider...
    ==> server1.netology: Importing base box 'bento/ubuntu-20.04'...
    ==> server1.netology: Matching MAC address for NAT networking...
    ==> server1.netology: Checking if box 'bento/ubuntu-20.04' version '202107.28.0' is up to date...
    ==> server1.netology: Setting the name of the VM: server1.netology
    ==> server1.netology: Clearing any previously set network interfaces...
    ==> server1.netology: Preparing network interfaces based on configuration...
        server1.netology: Adapter 1: nat
        server1.netology: Adapter 2: hostonly
    ==> server1.netology: Forwarding ports...
        server1.netology: 22 (guest) => 20011 (host) (adapter 1)
        server1.netology: 22 (guest) => 2222 (host) (adapter 1)
    ==> server1.netology: Running 'pre-boot' VM customizations...
    ==> server1.netology: Booting VM...
    ==> server1.netology: Waiting for machine to boot. This may take a few minutes...
        server1.netology: SSH address: 127.0.0.1:2222
        server1.netology: SSH username: vagrant
        server1.netology: SSH auth method: private key
        server1.netology: Warning: Connection reset. Retrying...
        server1.netology: Warning: Remote connection disconnect. Retrying...
        server1.netology: Warning: Connection reset. Retrying...
        server1.netology: Warning: Remote connection disconnect. Retrying...
        server1.netology: Warning: Connection reset. Retrying...
        server1.netology: Warning: Connection reset. Retrying...
        server1.netology: Warning: Remote connection disconnect. Retrying...
        server1.netology:
        server1.netology: Vagrant insecure key detected. Vagrant will automatically replace
        server1.netology: this with a newly generated keypair for better security.
        server1.netology:
        server1.netology: Inserting generated public key within guest...
        server1.netology: Removing insecure key from the guest if it's present...
        server1.netology: Key inserted! Disconnecting and reconnecting using new SSH key...
    ==> server1.netology: Machine booted and ready!
    ==> server1.netology: Checking for guest additions in VM...
    ==> server1.netology: Setting hostname...
    ==> server1.netology: Configuring and enabling network interfaces...
    ==> server1.netology: Mounting shared folders...
        server1.netology: /vagrant => /etc/ansible/hw_virt_2/src/vagrant
    ==> server1.netology: Running provisioner: ansible...
        server1.netology: Running ansible-playbook...
    [WARNING]: While constructing a mapping from
    /etc/ansible/hw_virt_2/src/ansible/provision.yml, line 8, column 5, found a
    duplicate dict key (— name). Using last defined value only.
    ERROR! A playbook must be a list of plays, got a <class 'ansible.parsing.yaml.objects.AnsibleMapping'> instead
    
    The error appears to be in '/etc/ansible/hw_virt_2/src/ansible/provision.yml': line 2, column 3, but may
    be elsewhere in the file depending on the exact syntax problem.
    
    The offending line appears to be:
    
    ---
      -hosts: nodes
      ^ here
    Ansible failed to complete successfully. Any error output should be
    visible above. Please fix these errors and try again.
  
  
  ###Тут было еще много запусков с ошибкой в provision файле
  
  
    root@netology-ansible:/etc/ansible/hw_virt_2/src/vagrant# vagrant provision                              ==> server1.netology: Running provisioner: ansible...
        server1.netology: Running ansible-playbook...
    
    PLAY [nodes] *******************************************************************
    
    TASK [Gathering Facts] *********************************************************
    ok: [server1.netology]
    
    TASK [Create directory for ssh-keys] *******************************************
    changed: [server1.netology]
    
    TASK [Adding rsa-key in /root/.ssh/authorized_keys] ****************************
    An exception occurred during task execution. To see the full traceback, use -vvv. The error was: If you are using a module and expect the file to exist on the remote, see the remote_src option
    fatal: [server1.netology]: FAILED! => {"changed": false, "msg": "Could not find or access '~/.ssh/id_rsa.pub' on the Ansible Controller.\nIf you are using a module and expect the file to exist on the remote, see the remote_src option"}
    ...ignoring
    
    TASK [Checking DNS] ************************************************************
    changed: [server1.netology]
    
    TASK [Installing tools] ********************************************************
    ok: [server1.netology] => (item=['git', 'curl'])
    
    TASK [Installing docker] *******************************************************
    changed: [server1.netology]
    
    TASK [Add the current user to docker group] ************************************
    changed: [server1.netology]
    
    PLAY RECAP *********************************************************************
    server1.netology           : ok=7    changed=4    unreachable=0    failed=0    skipped=0    rescued=0    ignored=1
   
  
  
  ###финальная команда
  
    vagrant ssh
    Welcome to Ubuntu 20.04.2 LTS (GNU/Linux 5.4.0-80-generic x86_64)
    
     * Documentation:  https://help.ubuntu.com
     * Management:     https://landscape.canonical.com
     * Support:        https://ubuntu.com/advantage
    
      System information as of Wed 03 Nov 2021 04:07:08 PM UTC
    
      System load:  0.66              Users logged in:          0
      Usage of /:   3.2% of 61.31GB   IPv4 address for docker0: 172.17.0.1
      Memory usage: 25%               IPv4 address for eth0:    10.0.2.15
      Swap usage:   0%                IPv4 address for eth1:    192.168.192.11
      Processes:    104
    
    
    This system is built by the Bento project by Chef Software
    More information can be found at https://github.com/chef/bento
    Last login: Wed Nov  3 16:03:41 2021 from 10.0.2.2
    vagrant@server1:~$ docker ps
    CONTAINER ID   IMAGE     COMMAND   CREATED   STATUS    PORTS     NAMES

