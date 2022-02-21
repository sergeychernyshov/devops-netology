
#1 
Пробрасываю ключи, чтобы по SSH можно было подключиться без пароля.

Генерирую ключ на контролере Ansible
>ssh-keygen -t rsa

При помощи команды отправляю ключ на удаленные сервер к которому буду подключаться.

>ssh-copy-id elastic@192.168.56.118
>ssh-copy-id kibana@192.168.56.115
>ssh-copy-id app@192.168.56.117

#2
Запускаю на Ubuntu 

#3

Раздаю возможность делать SUDO без пароля

    echo "elastic  ALL=(ALL) NOPASSWD:ALL" | sudo tee /etc/sudoers.d/elastic
    echo "kibana  ALL=(ALL) NOPASSWD:ALL" | sudo tee /etc/sudoers.d/kibana
    echo "app  ALL=(ALL) NOPASSWD:ALL" | sudo tee /etc/sudoers.d/app

#4

Ссылка на [playbook](https://github.com/sergeychernyshov/ansible-netology/tree/main/hw3)
