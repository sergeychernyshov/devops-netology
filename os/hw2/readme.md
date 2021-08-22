#1
    
    Установил node_exporter
    >node_exporter --version
 
    node_exporter, version 1.2.2 (branch: HEAD, revision: 26645363b486e12be40af7ce4fc91e731a33104e)
    build user:       root@b9cb4aa2eb17
    build date:       20210806-13:43:35
    go version:       go1.16.7
    platform:         linux/386
    
    управление сервисом

    sudo nano /etc/systemd/system/node-exporter.service
    
    [Unit]
    Description=node-exporter
    
    [Service]
    Restart=always
    User=vagrant
    ExecStart=/usr/local/bin/node_exporter
    ExecReload=/bin/kill -HUP $MAINPID
    EnvironmentFile=/etc/default/node_exporter
    TimeoutStopSec=20s
    SendSIGKILL=no
    
    [Install]
    WantedBy=multi-user.target

    Перезапускаю демона для чтения файла /etc/systemd/system/node-exporter.service 
    >sudo systemctl daemon-reload
    Добавляю в автозагрузку
    >sudo systemctl enable node-exporter.service

    Стартую
    >sudo systemctl start node-exporter.service
    
    >sudo systemctl status node-exporter.service

    node-exporter.service - node-exporter
     Loaded: loaded (/etc/systemd/system/node-exporter.service; enabled; vendor preset: enabled)
     Active: active (running) since Sat 2021-08-21 20:57:58 UTC; 5min ago
     Main PID: 1496 (node_exporter)

    останавливаю
    >sudo systemctl stop node-exporter.service
    >sudo systemctl status node-exporter.service
    node-exporter.service - node-exporter
     Loaded: loaded (/etc/systemd/system/node-exporter.service; enabled; vendor preset: enabled)
     Active: inactive (dead) since Sat 2021-08-21 21:05:14 UTC; 1min 54s ago
    Main PID: 1496 (code=killed, signal=TERM)
    
    после перезагрузки все упало выполнил
    >chmod 777 node_exporter

    После перезагрузки процесс успешно запустился
    >sudo systemctl status node-exporter.service

    node-exporter.service - node-exporter
     Loaded: loaded (/etc/systemd/system/node-exporter.service; enabled; vendor preset: enabled)
     Active: active (running) since Sat 2021-08-21 21:37:47 UTC; 20s ago
    Main PID: 641 (node_exporter)

#2

    Зашел на страницу http://localhost:9100/metrics
    
    cpu:
    node_cpu_seconds_total{cpu="0",mode="idle"} 67.55
    node_cpu_seconds_total{cpu="0",mode="iowait"} 0.95
    node_cpu_seconds_total{cpu="0",mode="irq"} 0
    node_cpu_seconds_total{cpu="3",mode="nice"} 0
    node_cpu_seconds_total{cpu="3",mode="softirq"} 0.11
    node_cpu_seconds_total{cpu="3",mode="steal"} 0
    node_cpu_seconds_total{cpu="3",mode="system"} 1.44
    node_cpu_seconds_total{cpu="3",mode="user"} 0.47

    disk:
    node_disk_read_bytes_total{device="dm-0"} 1.92918528e+08
    node_disk_read_bytes_total{device="dm-1"} 3.342336e+06
    node_disk_write_time_seconds_total{device="dm-0"} 0.932
    node_disk_write_time_seconds_total{device="dm-1"} 0
    node_disk_write_time_seconds_total{device="sda"} 0.76

    memory:
    node_memory_Active_anon_bytes 4.3450368e+07
    node_memory_Committed_AS_bytes 3.58764544e+08

    Network:
    node_network_address_assign_type{device="eth0"} 0
    node_network_address_assign_type{device="lo"} 0
    node_network_carrier{device="eth0"} 1
    node_network_carrier{device="lo"} 1

#3

    Установил netdata.
    Скриншот во вложении netdata.png

#4
    
    для определения окружения запуска ОС использовал команду
    >dmesg | grep -i virtual

    [    0.000000] DMI: innotek GmbH VirtualBox/VirtualBox, BIOS VirtualBox 12/01/2006
    [    0.001022] CPU MTRRs all blank - virtualized system.
    [    0.089888] Booting paravirtualized kernel on KVM
    [    2.472185] systemd[1]: Detected virtualization oracle.

#5

    sysctl fs.nr_open - показывает максимальное количество открытых файлов в одном процессе
    по умолчанию значение 1048576 = 1024*1024 (значение должно быть кратно 1024)
    
    sysctl fs.file-max - максимальное количество открыты файлов в системе
    9223372036854775807
    
    ulimit -Sn -мягкий лимит(может быть увеличен пользователем), находится в пределах [0, жесткий лимит]
    1024
    
    ulimit -Hn - жесткий лимит(может изменять только root)
    1048576
 
#6
    
    Запускаю процесс в новом namespace
    >unshare -f --pid --mount-proc sleep 1h
    
    >ps -e |grep sleep
    2294 pts/1    00:00:00 sleep

    подключаюсь к namespace по PID
    >nsenter --target 2294 --pid --mount

    >ps -e
    PID TTY          TIME CMD
      1 pts/1    00:00:00 sleep
      2 pts/2    00:00:00 bash
     40 pts/2    00:00:00 ps
    
#7

    Выполнил
    >:(){ :|:& };:
    Система подвисла, netdata показал жуткое поглощение ресурсов (скрин create_many_process.png) 
     
    :() - определение функции
    {...} - тело функции
    
    : вызывает себя
    | означает передачу выходных данных в команду
    : после | означает pipe к функции
    & означает выполнение команды лева в фоновом режиме
    Затем ; разделитель команд

    правый символ : запускает вызов бесконечной рекурсии, активируя "бомбу" fork.

    
    Запустил dmesg

    29 622 413 023 000 000 ns HostLast=1 629 591 067 147 000 000 ns)
    [ 5276.943693] 08:53:43.042657 timesync vgsvcTimeSyncWorker: Radical guest time change: 31 345 876 831 000ns (GuestNow=1 629 622 423 042 650 000 ns GuestLast=1 629 591 077 165 819 000 ns fSetTimeLastLoop=true )
    [ 8762.913927] cgroup: fork rejected by pids controller in /user.slice/user-1000.slice/session-24.scope
    [ 9479.833604] cgroup: fork rejected by pids controller in /user.slice/user-1000.slice/session-26.scope
    
    с определённого момента невозможно сделать fork процесса "fork rejected".
    Предполагаю, что это связано с ограничениями одновременно открытых процессов в сессии.
    
    > ulimit -u
    15389
    
    Можно установить ограничение на чиcло процессов
    ulimit -u 100 -  100 процессов для пользователя
    
    установил 100 процессов и запустил :(){ :|:& };:
    систему отпустило через несколько секунд.