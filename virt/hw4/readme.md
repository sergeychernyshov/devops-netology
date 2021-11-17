#1
###Установка docker-compose

>sudo curl -L https://github.com/docker/compose/releases/download/1.29.2/docker-compose-`uname -s`-`uname -m` -o /usr/local/bin/docker-compose

>docker-compose --version
docker-compose version 1.29.2, build 5becea4c


###Установка yandex cloud

Установка CLI в /opt/yandex-cloud, без изменений в файле .bashrc:
>curl https://storage.yandexcloud.net/yandexcloud-yc/install.sh | \bash -s -- -i /opt/yandex-cloud -n

Установка CLI в директорию по умолчанию, в файл .bashrc добавляются completion и PATH:

>curl https://storage.yandexcloud.net/yandexcloud-yc/install.sh | \bash -s -- -a

>yc --version
Yandex.Cloud CLI 0.83.0 linux/amd64


>yc config list

    token: AQAAAABaQHJWAATuwWzUJzeKT0eChEL-zPlFe7s
    cloud-id: b1gjvd13grcv951tkt06
    folder-id: b1gioaj7tlvmc44fvd40
    compute-default-zone: ru-central1-a

>yc compute image list

    +----+------+--------+-------------+--------+
    | ID | NAME | FAMILY | PRODUCT IDS | STATUS |
    +----+------+--------+-------------+--------+



>yc vpc network create --name net-netology

    id: enp46c3uot8c8gt5h2tc
    folder_id: b1g91q2tu6if6gi7erv4
    created_at: "2021-11-15T19:52:46Z"
    name: net-netology

>yc vpc subnet create --name subnet-netology --zone ru-central1-a --range 10.1.2.0/24 --network-name net-netology  

    id: e9bombsvud6o7fih9be2
    folder_id: b1g91q2tu6if6gi7erv4
    created_at: "2021-11-15T19:57:27Z"
    name: subnet-netology
    network_id: enp46c3uot8c8gt5h2tc
    zone_id: ru-central1-a
    v4_cidr_blocks:
    - 10.1.2.0/24

>yc vpc subnet list

    +----------------------+-----------------+----------------------+----------------+---------------+---------------+
    |          ID          |      NAME       |      NETWORK ID      | ROUTE TABLE ID |     ZONE      |     RANGE     |
    +----------------------+-----------------+----------------------+----------------+---------------+---------------+
    | e9bombsvud6o7fih9be2 | subnet-netology | enp46c3uot8c8gt5h2tc |                | ru-central1-a | [10.1.2.0/24] |
    +----------------------+-----------------+----------------------+----------------+---------------+---------------+

>yc vpc net list

    +----------------------+--------------+
    |          ID          |     NAME     |
    +----------------------+--------------+
    | enp46c3uot8c8gt5h2tc | net-netology |
    +----------------------+--------------+

###Установка packer

>sudo apt install packer

>packer --version

    1.7.8


###Создание образа centos
>packer validate centos-7-base.json

    The configuration is valid.

>packer build centos-7-base.json

    Wait completed after 2 minutes 6 seconds
    
    ==> Builds finished. The artifacts of successful builds are:
    --> yandex: A disk image was created: centos-7-base (id: fd85knoha6lb4o5tkt5i) with family name centos

>yc compute image list

    +----------------------+---------------+--------+----------------------+--------+
    |          ID          |     NAME      | FAMILY |     PRODUCT IDS      | STATUS |
    +----------------------+---------------+--------+----------------------+--------+
    | fd85knoha6lb4o5tkt5i | centos-7-base | centos | f2ebfhrshe5m6i4saf1j | READY  |
    +----------------------+---------------+--------+----------------------+--------+

#2

###Установка terraform

>curl -fsSL https://apt.releases.hashicorp.com/gpg | sudo apt-key add -
>sudo apt-add-repository "deb [arch=$(dpkg --print-architecture)] https://apt.releases.hashicorp.com $(lsb_release -cs) main"
>sudo apt install terraform

>terraform --version

    Terraform v1.0.11
    on linux_amd64

>yc vpc subnet delete e9bombsvud6o7fih9be2
>yc vpc net delete enp46c3uot8c8gt5h2tc


>terraform init

    Initializing the backend...
    
    Initializing provider plugins...
    - Finding latest version of yandex-cloud/yandex...
    - Installing yandex-cloud/yandex v0.66.0...
    - Installed yandex-cloud/yandex v0.66.0 (self-signed, key ID E40F590B50BB8E40)
    
    Partner and community providers are signed by their developers.
    If you'd like to know more about provider signing, you can read about it here:
    https://www.terraform.io/docs/cli/plugins/signing.html
    
    Terraform has created a lock file .terraform.lock.hcl to record the provider
    selections it made above. Include this file in your version control repository
    so that Terraform can guarantee to make the same selections by default when
    you run "terraform init" in the future.
    
    Terraform has been successfully initialized!
    
    You may now begin working with Terraform. Try running "terraform plan" to see
    any changes that are required for your infrastructure. All Terraform commands
    should now work.
    
    If you ever set or change modules or backend configuration for Terraform,
    rerun this command to reinitialize your working directory. If you forget, other
    commands will detect it and remind you to do so if necessary.


>yc iam service-account create --name chernyshov-sergey --folder-id b1g91q2tu6if6gi7erv4
>yc iam key create --folder-name chernyshov-sergey --service-account-name chernyshov-sergey --output key.json

Далее аккаунт настроил по мануалу
https://cloud.yandex.ru/docs/cli/operations/authentication/service-account

>terraform validate

    Success! The configuration is valid.

> terraform plan
> terraform apply -auto-approve

    Apply complete! Resources: 3 added, 0 changed, 0 destroyed.

#3
>ansible-playbook provision.yml

>docker-compose ps

    Name               Command               State                Ports
    --------------------------------------------------------------------------------
    alertmanager   /bin/alertmanager          Up             9093/tcp
                   --config ...
    caddy          /sbin/tini -- caddy        Up             0.0.0.0:3000->3000/tcp,
                   -agree ...                                0.0.0.0:9090->9090/tcp,
                                                             0.0.0.0:9091->9091/tcp,
                                                             0.0.0.0:9093->9093/tcp
    cadvisor       /usr/bin/cadvisor          Up (healthy)   8080/tcp
                   -logtostderr
    grafana        /run.sh                    Up             3000/tcp
    nodeexporter   /bin/node_exporter         Up             9100/tcp
                   --path. ...
    prometheus     /bin/prometheus            Up             9090/tcp
                   --config.f ...
    pushgateway    /bin/pushgateway           Up             9091/tcp
 
Скриншот задания во вложении