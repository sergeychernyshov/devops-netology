#1

    Мой Ip адрес 194.15.119.144 получен на сайте 2ip.ru (task1_1.png)
    telnet route-views.routeviews.org
    Username: rviews
    >show ip route 194.15.119.144
 
    Routing entry for 194.15.116.0/22, supernet
      Known via "bgp 6447", distance 20, metric 0
      Tag 6939, type external
      Last update from 64.71.137.241 2d07h ago
      Routing Descriptor Blocks:
      * 64.71.137.241, from 64.71.137.241, 2d07h ago
          Route metric is 0, traffic share count is 1
          AS Hops 3
          Route tag 6939
          MPLS label: none

#2

    >echo "dummy" >> /etc/modules
    >echo "options dummy numdummies=2" > /etc/modprobe.d/dummy.conf
    >echo "options dummy numdummies=2" > /etc/modprobe.d/dummy.conf
    root@oracle-VirtualBox:~# ifconfig
    enp0s3: flags=4163<UP,BROADCAST,RUNNING,MULTICAST>  mtu 1500
            inet 192.168.56.102  netmask 255.255.255.0  broadcast 192.168.56.255
            inet6 fe80::ada:9550:623f:fc18  prefixlen 64  scopeid 0x20<link>
            ether 08:00:27:e0:61:94  txqueuelen 1000  (Ethernet)
            RX packets 901  bytes 134849 (134.8 KB)
            RX errors 0  dropped 0  overruns 0  frame 0
            TX packets 11372  bytes 1187476 (1.1 MB)
            TX errors 0  dropped 0 overruns 0  carrier 0  collisions 0
    
    enp0s8: flags=4163<UP,BROADCAST,RUNNING,MULTICAST>  mtu 1500
            inet 10.0.3.15  netmask 255.255.255.0  broadcast 10.0.3.255
            inet6 fe80::fe20:627d:3e54:1967  prefixlen 64  scopeid 0x20<link>
            ether 08:00:27:f7:f8:66  txqueuelen 1000  (Ethernet)
            RX packets 74100  bytes 106072571 (106.0 MB)
            RX errors 0  dropped 0  overruns 0  frame 0
            TX packets 17678  bytes 1349257 (1.3 MB)
            TX errors 0  dropped 0 overruns 0  carrier 0  collisions 0
    
    lo: flags=73<UP,LOOPBACK,RUNNING>  mtu 65536
            inet 127.0.0.1  netmask 255.0.0.0
            inet6 ::1  prefixlen 128  scopeid 0x10<host>
            loop  txqueuelen 1000  (Local Loopback)
            RX packets 1030  bytes 103038 (103.0 KB)
            RX errors 0  dropped 0  overruns 0  frame 0
            TX packets 1030  bytes 103038 (103.0 KB)
            TX errors 0  dropped 0 overruns 0  carrier 0  collisions 0
    
    >vim /etc/network/interfaces
    

    auto dummy0
    iface dummy0 inet static
    address 10.10.0.1
    netmask 255.255.255.0

    настраиваем dummy0
    >ip link add dummy0 type dummy
    >ip link set dummy0 multicast on
    >ip addr add 10.10.0.1/24 dev dummy0
    >ip link set dummy0 up

    >ifconfig
    dummy0: flags=4291<UP,BROADCAST,RUNNING,NOARP,MULTICAST>  mtu 1500
            inet 10.10.0.1  netmask 255.255.255.0  broadcast 0.0.0.0
            inet6 fe80::ccae:7ff:fe77:93ca  prefixlen 64  scopeid 0x20<link>
            ether ce:ae:07:77:93:ca  txqueuelen 1000  (Ethernet)
            RX packets 0  bytes 0 (0.0 B)
            RX errors 0  dropped 0  overruns 0  frame 0
            TX packets 49  bytes 7035 (7.0 KB)
            TX errors 0  dropped 0 overruns 0  carrier 0  collisions 0
    
    enp0s3: flags=4163<UP,BROADCAST,RUNNING,MULTICAST>  mtu 1500
            inet 192.168.56.102  netmask 255.255.255.0  broadcast 192.168.56.255
            inet6 fe80::ada:9550:623f:fc18  prefixlen 64  scopeid 0x20<link>
            ether 08:00:27:e0:61:94  txqueuelen 1000  (Ethernet)
            RX packets 1794  bytes 222809 (222.8 KB)
            RX errors 0  dropped 0  overruns 0  frame 0
            TX packets 12823  bytes 1462548 (1.4 MB)
            TX errors 0  dropped 0 overruns 0  carrier 0  collisions 0
    
    enp0s8: flags=4163<UP,BROADCAST,RUNNING,MULTICAST>  mtu 1500
            inet 10.0.3.15  netmask 255.255.255.0  broadcast 10.0.3.255
            inet6 fe80::fe20:627d:3e54:1967  prefixlen 64  scopeid 0x20<link>
            ether 08:00:27:f7:f8:66  txqueuelen 1000  (Ethernet)
            RX packets 74304  bytes 106114207 (106.1 MB)
            RX errors 0  dropped 0  overruns 0  frame 0
            TX packets 17904  bytes 1379964 (1.3 MB)
            TX errors 0  dropped 0 overruns 0  carrier 0  collisions 0
    
    lo: flags=73<UP,LOOPBACK,RUNNING>  mtu 65536
            inet 127.0.0.1  netmask 255.0.0.0
            inet6 ::1  prefixlen 128  scopeid 0x10<host>
            loop  txqueuelen 1000  (Local Loopback)
            RX packets 1251  bytes 126535 (126.5 KB)
            RX errors 0  dropped 0  overruns 0  frame 0
            TX packets 1251  bytes 126535 (126.5 KB)
            TX errors 0  dropped 0 overruns 0  carrier 0  collisions 0

    >route
    Kernel IP routing table
    Destination     Gateway         Genmask         Flags Metric Ref    Use Iface
    default         _gateway        0.0.0.0         UG    101    0        0 enp0s8
    10.0.3.0        0.0.0.0         255.255.255.0   U     101    0        0 enp0s8
    10.10.0.0       0.0.0.0         255.255.255.0   U     0      0        0 dummy0
    link-local      0.0.0.0         255.255.0.0     U     1000   0        0 enp0s3
    192.168.56.0    0.0.0.0         255.255.255.0   U     100    0        0 enp0s3

#3

    >netstat -at
    Active Internet connections (servers and established)
    Proto Recv-Q Send-Q Local Address           Foreign Address         State
    tcp        0      0 localhost:domain        0.0.0.0:*               LISTEN
    tcp        0      0 0.0.0.0:ssh             0.0.0.0:*               LISTEN
    tcp        0      0 localhost:ipp           0.0.0.0:*               LISTEN
    tcp        0      0 oracle-VirtualBox:49794 ec2-52-38-198-132:https ESTABLISHED
    tcp        0      0 oracle-VirtualBox:55176 server-13-33-246-:https ESTABLISHED
    tcp        0      0 oracle-VirtualBox:52932 ec2-35-155-229-13:https TIME_WAIT
    tcp        0      0 oracle-VirtualBox:52936 ec2-35-155-229-13:https TIME_WAIT
    tcp        0      0 oracle-VirtualBox:ssh   192.168.56.1:49678      ESTABLISHED
    tcp6       0      0 [::]:ssh                [::]:*                  LISTEN
    tcp6       0      0 ip6-localhost:ipp       [::]:*                  LISTEN

    используются протоколы SSH, IPP, HTTPS
    приложения oracle-VirtualBox   

#4
    
    > netstat sockets | grep udp
    udp        0      0 oracle-VirtualBo:bootpc _gateway:bootps         ESTABLISHED
    udp        0      0 oracle-VirtualBo:bootpc 192.168.56.100:bootps   ESTABLISHED

    протокол bootps
    приложения oracle-VirtualBox   

#5

    Схема домашней сети в файле home-network.html 

#6

    >apt install nginx

    Настройка TCP балансировки между 192.168.56.12:3002 и 192.168.56.13:3003
    >vim /etc/nginx/nginx.conf
    
    stream {
    upstream test_balance {
        server 192.168.56.12:3002 weight=5;
        server 192.168.56.13:3003;

    }
        server {
            listen 3001;
            proxy_pass test_balance;
        }
    }
    >systemctl restart nginx
    >systemctl status nginx
     nginx.service - A high performance web server and a reverse proxy server
     Loaded: loaded (/lib/systemd/system/nginx.service; enabled; vendor preset: enabled)
     Active: active (running) since Sun 2021-09-26 00:21:58 MSK; 15s ago
       Docs: man:nginx(8)

#7

#8