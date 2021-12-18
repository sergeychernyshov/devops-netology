#1

##1.1

###Какой тип инфраструктуры будем использовать для этого проекта: изменяемый или не изменяемый?

Предполагаю, что инфраструктура будет двух вариантов.
Например, докеры будут неизменяемы, а виртуальные машины можно будет менять.
Это обусловлено большим количеством изменений на начальном этапе разработки в условиях нечеткого тех задания
и большого количества релизов.

###Будет ли центральный сервер для управления инфраструктурой?

На первом этапе предполагаю что центрального сервера управления инфраструктуры не будет.
После понимания какой у нас будет сервис, решение может поменяться. В условиях неопределенности 
буду стараться использовать точечное управление.

###Будут ли агенты на серверах?

Предполагаю, что под агентами понимаются процессы собирающие телеметрию сервера, 
такие агенты будут использоваться. Так как мы находимся в условиях не определенности и не можем знать
насколько наш сервис будет нагружен.  

###Будут ли использованы средства для управления конфигурацией или инициализации ресурсов?

Да, буду использовать.

##1.2

###Какие инструменты из уже используемых вы хотели бы использовать для нового проекта?

Буду использовать Docker контейнеры для решения задач. 
Kubernetes для размещения контейнеров.
Packer для конфигурации образов виртуальных машин, Terraform для инициализации ресурсов.

##1.3

###Хотите ли рассмотреть возможность внедрения новых инструментов для этого проекта?

Нет, буду использовать проверенные инструменты. Т.к. проект новый и нет возможности
понять что нам может потребоваться. После первых этапов проекта можно будет сравнивать удобства
надежность, производительность текущих инструментов с новыми.

#2

Создаю публичный ключ для репозитория terraform
>add-apt-repository ppa:webupd8team/y-ppa-manager
>apt-get update
>ssh-keygen

Инструкция по установке terraform на ubuntu https://www.terraform.io/downloads


>curl -fsSL https://apt.releases.hashicorp.com/gpg |  apt-key add -
    
    OK

>apt-add-repository "deb [arch=amd64] https://apt.releases.hashicorp.com $(lsb_release -cs) main"
    
    Hit:1 http://ru.archive.ubuntu.com/ubuntu focal InRelease
    Hit:2 http://ru.archive.ubuntu.com/ubuntu focal-updates InRelease
    Hit:3 http://ru.archive.ubuntu.com/ubuntu focal-backports InRelease
    Hit:4 http://ppa.launchpad.net/webupd8team/y-ppa-manager/ubuntu focal InRelease
    Hit:5 https://apt.releases.hashicorp.com focal InRelease
    Hit:6 http://security.ubuntu.com/ubuntu focal-security InRelease
    Reading package lists... Done

>apt-get update
>apt-get install terraform

>terraform --version

    Terraform v1.1.2
    on linux_amd64

#3

Установил утилиту tfswitch
https://tfswitch.warrensbox.com/Install/

>curl -L https://raw.githubusercontent.com/warrensbox/terraform-switcher/release/install.sh | bash
> tfswitch
Creating directory for terraform binary at: /root/.terraform.versions
✔ 0.12.31

    Downloading to: /root/.terraform.versions
    28441056 bytes downloaded
    Switched terraform to version "0.12.31"

>tfswitch 0.12.31
Switched terraform to version "0.12.31"

>terraform --version

    Terraform v1.1.2

Не вышло просто сменить версию...

Тут нашел решение
https://tfswitch.warrensbox.com/Troubleshoot/

>wget https://raw.githubusercontent.com/warrensbox/terraform-switcher/release/install.sh
>chmod 755 install.sh
>./install.sh -b $HOME/.bin 
>$HOME/.bin/tfswitch

    ✔ 0.12.27 *recent
    Switched terraform to version "0.12.27"

>export PATH=$PATH:$HOME/.bin 
>tfswitch -b $HOME/.bin/terraform 0.12.27

    Switched terraform to version "0.12.27"
    root@oracle-terraform:~/.terraform.versions# terraform --version
    Terraform v0.12.27
    
    Your version of Terraform is out of date! The latest version
    is 1.1.2. You can update by downloading from https://www.terraform.io/downloads.html
