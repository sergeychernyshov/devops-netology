#1
    Linux
    >ip -o a show | cut -d ' ' -f 2,7 
    
    lo 127.0.0.1/8
    lo ::1/128
    eth0 10.0.2.15/24
    eth0 fe80::a00:27ff:fe73:60cf/64
    
    window
    >netsh interface show interface

    Admin State    State          Type             Interface Name
    -------------------------------------------------------------------------
    Enabled        Connected      Dedicated        Ethernet 2
    Enabled        Connected      Dedicated        VirtualBox Host-Only Network
    Enabled        Connected      Dedicated        VirtualBox Host-Only Network #2
    Enabled        Connected      Dedicated        Беспроводная сеть
    Enabled        Disconnected   Dedicated        Ethernet

#2
    LLDP – протокол для обмена информацией между соседними устройствами. 
    К сожалению не нашел соседства устройств с соседями по коммутатору
    
    >yum install lldpd
    >systemctl enable lldpd
    >systemctl start lldpd
    >systemctl status lldpd
    ● lldpd.service - LLDP daemon
       Loaded: loaded (/usr/lib/systemd/system/lldpd.service; enabled; vendor preset: disabled)
       Active: active (running) since Mon 2021-09-20 17:21:20 MSK; 3min 8s ago

    >lldpctl
    -------------------------------------------------------------------------------
    LLDP neighbors:

#3
    VLAN(Virtual Local Area Network) – виртуальное разделение коммутатора на несколько виртуальных интерфейсов на одном 
    физическом сетевом интерфейсе. С помощью VLAN можно разделять или объединять сегменты локальной сети.
    
    пакет для работы с VLAN
    >apt install vlan

#4
    Bonding – это объединение сетевых интерфейсов по определенному типу агрегации. 
    Служит для увеличения пропускной способности и/или отказоустойчивость сети.

    Типы агрегации интерфейсов в Linux:

    # Mode-0(balance-rr) – Данный режим используется по умолчанию. 
    Balance-rr обеспечивается балансировку нагрузки и отказоустойчивость. 
    В данном режиме сетевые пакеты отправляются “по кругу”, от первого интерфейса к последнему. 
    Если выходят из строя интерфейсы, пакеты отправляются на остальные оставшиеся. 
    Дополнительной настройки коммутатора не требуется при нахождении портов в одном коммутаторе. 
    При разностных коммутаторах требуется дополнительная настройка.

    # Mode-1(active-backup) – Один из интерфейсов работает в активном режиме, остальные в ожидающем. 
    При обнаружении проблемы на активном интерфейсе производится переключение на ожидающий интерфейс. 
    Не требуется поддержки от коммутатора.

    # Mode-2(balance-xor) – Передача пакетов распределяется по типу входящего и исходящего трафика по формуле 
    ((MAC src) XOR (MAC dest)) % число интерфейсов. Режим дает балансировку нагрузки и отказоустойчивость. 
    Не требуется дополнительной настройки коммутатора/коммутаторов.

    # Mode-3(broadcast) – Происходит передача во все объединенные интерфейсы, тем самым обеспечивая отказоустойчивость. 
    Рекомендуется только для использования MULTICAST трафика.

    # Mode-4(802.3ad) – динамическое объединение одинаковых портов. 
    В данном режиме можно значительно увеличить пропускную способность входящего так и исходящего трафика. 
    Для данного режима необходима поддержка и настройка коммутатора/коммутаторов.

    # Mode-5(balance-tlb) – Адаптивная балансировки нагрузки трафика. 
    Входящий трафик получается только активным интерфейсом, исходящий распределяется в зависимости 
    от текущей загрузки канала каждого интерфейса.
    Не требуется специальной поддержки и настройки коммутатора/коммутаторов.

    # Mode-6(balance-alb) – Адаптивная балансировка нагрузки. 
    Отличается более совершенным алгоритмом балансировки нагрузки чем Mode-5). 
    Обеспечивается балансировку нагрузки как исходящего так и входящего трафика. 
    Не требуется специальной поддержки и настройки коммутатора/коммутаторов.

    
    Пример настройки интерфейсов eth0 и eth1 в режиме active-backup в файле «/etc/network/interfaces»:

    auto bond0
    iface bond0 inet dhcp
       bond-slaves eth0 eth1
       bond-mode active-backup
       bond-miimon 100
       bond-primary eth0 eth1

#5
    В подсети с маской /29 свободно 3 бита.
    3^2 = 8 cоотвественно сеть состоит всего из 8-ми адресов. 
    8-2 = 6 узловых адресов.

    2^(8-(32-29)) = 2^5 = 32 подсети /29  можно получить из сети с маской /24

    10.10.10.0/24
    подсеть
    10.10.10.248/29
    10.10.10.240/29
    10.10.10.232/29

#6
    использовать сеть 100.64.0.0
    Можно использовать подсеть
    100.64.1.192/26

#7
    Windows
    просмотр 
    arp -a
    
    удалить один хост
    arp -d 192.168.100.25

    очистить кеш в Windows
    netsh interface ip delete arpcache

    Linux
    arp -v

    удалить один хост
    arp -d 10.0.0.2

    очистить кеш в Linux
    ip -s -s neigh flush all