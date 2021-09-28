#1

    Установка и использование Bitwarden на скриншотах task1_*.png

#2

    Установка и использование в Bitwarden двухфакторной авторизации
    при помощи приложения Google Authenticator на скриншотах task2_*.png

#3
    
    Установка apache
    >apt install apache2
    >ufw app list
 
    Available applications:
      Apache
      Apache Full
      Apache Secure
      CUPS
      OpenSSH

    >ufw allow 'Apache'
    Rule added
    Rule added (v6)

    >ufw status
    Status: active
    
    To                         Action      From
    --                         ------      ----
    22/tcp                     ALLOW       Anywhere
    22/tcp                     ALLOW       192.168.56.1
    Apache                     ALLOW       Anywhere
    22/tcp (v6)                ALLOW       Anywhere (v6)
    Apache (v6)                ALLOW       Anywhere (v6)

    >systemctl status apache2
    
    apache2.service - The Apache HTTP Server
     Loaded: loaded (/lib/systemd/system/apache2.service; enabled; vendor prese>
     Active: active (running) since Sun 2021-09-26 15:47:45 MSK; 29min ago
    
    включаем возможность иcпользовать сертификат
    >a2enmod ssl
    >systemctl restart apache2
    
    Генерируем сертификат
    >openssl req -x509 -nodes -days 365 -newkey rsa:2048 \-keyout /etc/ssl/private/apache-selfsigned.key \-out /etc/ssl/certs/apache-selfsigned.crt \-subj "/C=RU/ST=Moscow/L=Moscow/O=Company Name/OU=Org/CN=www.chernyshov.com"
    Generating a RSA private key
    ..............................................................................................+++++
    ...........................................+++++
    writing new private key to '/etc/ssl/private/apache-selfsigned.key'
    -----
    
    >vim /etc/apache2/sites-available/192.168.56.104.conf
    >mkdir /var/www/192.168.56.104
    >vim /var/www/192.168.56.104/index.html

    >systemctl reload apache2
    >a2ensite 192.168.56.104.conf
    >apache2ctl configtest
    AH00558: apache2: Could not reliably determine the server's fully qualified domain name, using 127.0.1.1. Set the 'ServerName' directive globally to suppress this message
    Syntax OK

    >ufw disable

    Перехожу на сайт https://192.168.56.104/index.html 
    результат на скриншоте task3_1.png


###3.1
    Добавил в /etc/hosts запись 192.168.56.104  www.chernyshov.com
    перехожу по сслыке https://www.chernyshov.com
    Смотрю инфо о сертификате task3.1_1.png
    Сайт выглядит как на скрине task3.1_2.png

#4

    >apt install git
    >git clone --depth 1 https://github.com/drwetter/testssl.sh.git

     >./testssl.sh -U --sneaky https://192.168.56.104

###########################################################
    testssl.sh       3.1dev from https://testssl.sh/dev/
    (b8bff80 2021-09-24 14:21:04 -- )

      This program is free software. Distribution and
             modification under GPLv2 permitted.
      USAGE w/o ANY WARRANTY. USE IT AT YOUR OWN RISK!

       Please file bugs @ https://testssl.sh/bugs/

###########################################################

     Using "OpenSSL 1.0.2-chacha (1.0.2k-dev)" [~183 ciphers]
     on oracle-VirtualBox:./bin/openssl.Linux.x86_64
     (built: "Jan 18 17:12:17 2019", platform: "linux-x86_64")
    
    
     Start 2021-09-26 17:17:22        -->> 192.168.56.104:443 (192.168.56.104) <<--
    
     rDNS (192.168.56.104):  oracle-VirtualBox. oracle-VirtualBox.local.
     Service detected:       HTTP
    
    
     Testing vulnerabilities
    
     Heartbleed (CVE-2014-0160)                not vulnerable (OK), no heartbeat extension
     CCS (CVE-2014-0224)                       not vulnerable (OK)
     Ticketbleed (CVE-2016-9244), experiment.  not vulnerable (OK)
     ROBOT                                     not vulnerable (OK)
     Secure Renegotiation (RFC 5746)           supported (OK)
     Secure Client-Initiated Renegotiation     not vulnerable (OK)
     CRIME, TLS (CVE-2012-4929)                not vulnerable (OK)
     BREACH (CVE-2013-3587)                    no gzip/deflate/compress/br HTTP compression (OK)  - only supplied "/" tested
     POODLE, SSL (CVE-2014-3566)               not vulnerable (OK)
     TLS_FALLBACK_SCSV (RFC 7507)              No fallback possible (OK), no protocol below TLS 1.2 offered
     SWEET32 (CVE-2016-2183, CVE-2016-6329)    not vulnerable (OK)
     FREAK (CVE-2015-0204)                     not vulnerable (OK)
     DROWN (CVE-2016-0800, CVE-2016-0703)      not vulnerable on this host and port (OK)
                                               no RSA certificate, thus certificate can't be used with SSLv2 elsewhere
     LOGJAM (CVE-2015-4000), experimental      common prime with 2048 bits detected: RFC3526/Oakley Group 14 (2048 bits),
                                               but no DH EXPORT ciphers
     BEAST (CVE-2011-3389)                     not vulnerable (OK), no SSL3 or TLS1
     LUCKY13 (CVE-2013-0169), experimental     potentially VULNERABLE, uses cipher block chaining (CBC) ciphers with TLS. Check patches
     Winshock (CVE-2014-6321), experimental    not vulnerable (OK)
     RC4 (CVE-2013-2566, CVE-2015-2808)        no RC4 ciphers detected (OK)
    
    
     Done 2021-09-26 17:17:45 [  24s] -->> 192.168.56.104:443 (192.168.56.104) <<--

************************************************************************************************************************

    >./testssl.sh -U --sneaky https://mail.ru
    Start 2021-09-26 17:20:01        -->> 217.69.139.202:443 (mail.ru) <<--

     Further IP addresses:   217.69.139.200 94.100.180.201 94.100.180.200 2a00:1148:db00:0:b0b0::1
     rDNS (217.69.139.202):  mail.ru.
     Service detected:       HTTP
    
    
     Testing vulnerabilities
    
     Heartbleed (CVE-2014-0160)                not vulnerable (OK), timed out
     CCS (CVE-2014-0224)                       not vulnerable (OK)
     Ticketbleed (CVE-2016-9244), experiment.  not vulnerable (OK)
     ROBOT                                     not vulnerable (OK)
     Secure Renegotiation (RFC 5746)           supported (OK)
     Secure Client-Initiated Renegotiation     not vulnerable (OK)
     CRIME, TLS (CVE-2012-4929)                not vulnerable (OK)
     BREACH (CVE-2013-3587)                    potentially NOT ok, "gzip" HTTP compression detected. - only supplied "/" tested
                                               Can be ignored for static pages or if no secrets in the page
     POODLE, SSL (CVE-2014-3566)               not vulnerable (OK)
     TLS_FALLBACK_SCSV (RFC 7507)              Downgrade attack prevention supported (OK)
     SWEET32 (CVE-2016-2183, CVE-2016-6329)    VULNERABLE, uses 64 bit block ciphers
     FREAK (CVE-2015-0204)                     not vulnerable (OK)
     DROWN (CVE-2016-0800, CVE-2016-0703)      not vulnerable on this host and port (OK)
                                               make sure you don't use this certificate elsewhere with SSLv2 enabled services
                                               https://censys.io/ipv4?q=73CE7337E1FE4D5E6CBAB304B5E401B21C006CCEC612092AD83209BBABEED18B could help you to find out
     LOGJAM (CVE-2015-4000), experimental      not vulnerable (OK): no DH EXPORT ciphers, no common prime detected
     BEAST (CVE-2011-3389)                     TLS1: ECDHE-RSA-AES256-SHA ECDHE-RSA-AES128-SHA ECDHE-RSA-DES-CBC3-SHA DHE-RSA-AES256-SHA
                                                     DHE-RSA-AES128-SHA EDH-RSA-DES-CBC3-SHA AES256-SHA AES128-SHA DES-CBC3-SHA
                                               VULNERABLE -- but also supports higher protocols  TLSv1.1 TLSv1.2 (likely mitigated)
     LUCKY13 (CVE-2013-0169), experimental     potentially VULNERABLE, uses cipher block chaining (CBC) ciphers with TLS. Check patches
     Winshock (CVE-2014-6321), experimental    not vulnerable (OK) - CAMELLIA or ECDHE_RSA GCM ciphers found
     RC4 (CVE-2013-2566, CVE-2015-2808)        no RC4 ciphers detected (OK)
    
    
     Done 2021-09-26 17:20:42 [ 127s] -->> 217.69.139.202:443 (mail.ru) <<--

#5
    
    > ssh-keygen
    Generating public/private rsa key pair.
    Enter file in which to save the key (/root/.ssh/id_rsa):
    
    Created directory '/root/.ssh'.
    Enter passphrase (empty for no passphrase):
    Enter same passphrase again:
    Your identification has been saved in /root/.ssh/id_rsa
    Your public key has been saved in /root/.ssh/id_rsa.pub
    The key fingerprint is:
    SHA256:AEB0RnhN0oXv1tpAPmwuI4Hb1RISAudzwyBrgy/cdrI root@oracle-VirtualBox
    The key's randomart image is:
    +---[RSA 3072]----+
    |o+*+*+.o.        |
    |.=o=oo+          |
    |oooo+...         |
    |o.oo....o        |
    |...= o BS.       |
    | .o = o X .      |
    |   E o = =       |
    |  . o o o .      |
    |     . o         |
    +----[SHA256]-----+

    >cd /root/.ssh/
    >ls -la
    total 20
    drwx------ 2 root root    4 сен 26 17:34 .
    drwx------ 5 root root    8 сен 26 17:33 ..
    -rw------- 1 root root 2655 сен 26 17:34 id_rsa
    -rw-r--r-- 1 root root  576 сен 26 17:34 id_rsa.pub

    >vim id_rsa.pub
    ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABgQC8SiVR8nYbFzzyDl/McR+W62UayRnUc2zG6jxTYEeBxSpOvtxnYsfTd+etGpRz957kpp3m50b6YKer1XS5hO7v7XFRJCUWj/hkWcQy60X1AClN4QOfxGZQG8/TkdSqqU22ntue4r61dALACOlXsVxBJfOulVb/hYXm1ZseYEl0NMEJ5Ko1ZFw9C/vS68dbaX13F1+mzox2fxq3i5qDnojNA1REGEng99I3nR7Dpi3IFqyXyD43q2cSVWMRjCo/jWMWz8u6SxhAG03GU5PPWJ2zE4+27uXtyC/8DyRBoS4JY7Y8wSNxiTQnt5/tUELXTwLzIYadBEiMkv4bfhSlVzbFPXhkN2k478hUzFixBAH3kaKCLwlAlNOXoLa/61rRZiV9C4ONeKM/f4rQAuhbdG+NhGzTpTjTbqNfgy74Mo5d+yM9lNxJFdc+5yafKP1hYoCljzUNEFgwVlpgjjTbM/qD82tkOZGdBf/3QDY8JYJP0EBV87E0Um87w/jcuSvMXsU= root@oracle-VirtualBox

    >ssh-copy-id oracle@192.168.56.102
    
    >ssh oracel@192.168.56.104


###5.1
    
    хост 192.168.56.104
    > ssh-keygen
    Generating public/private rsa key pair.
    Enter file in which to save the key (/root/.ssh/id_rsa):
    /root/.ssh/id_rsa already exists.
    Overwrite (y/n)?
    root@oracle-VirtualBox:~#
    root@oracle-VirtualBox:~# ssh-keygen
    Generating public/private rsa key pair.
    Enter file in which to save the key (/root/.ssh/id_rsa):
    /root/.ssh/id_rsa already exists.
    Overwrite (y/n)? y
    Enter passphrase (empty for no passphrase):
    Enter same passphrase again:
    Your identification has been saved in /root/.ssh/id_rsa
    Your public key has been saved in /root/.ssh/id_rsa.pub
    The key fingerprint is:
    SHA256:TQloCp9StbGyaeisbSRg7TFgX0d6TL50RI3esJmIQzI root@oracle-VirtualBox
    The key's randomart image is:
    +---[RSA 3072]----+
    |     .++ooo      |
    | o.E.oB=.+ o     |
    |. ++*===o.O      |
    |...*+*o.oB .     |
    |o o.* ..S .      |
    |.o.o             |
    | oo              |
    | o.              |
    |...              |
    +----[SHA256]-----+

    >cd /root/.ssh/
    >ls -la
    total 38
    drwx------ 2 root root    6 сен 28 23:20 .
    drwx------ 7 root root   13 сен 28 23:17 ..
    -rw-r--r-- 1 root root  576 сен 26 21:59 authorized_keys
    -rw------- 1 root root 2610 сен 28 23:22 id_rsa
    -rw-r--r-- 1 root root  576 сен 28 23:22 id_rsa.pub
    -rw-r--r-- 1 root root  444 сен 28 23:20 known_hosts
    root@oracle-VirtualBox:~/.ssh# cat id_rsa.pub
    ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABgQDGy/6xHhV+1EH1fZAXDNlW3PBpkNCE16ZPV+DQtIJK                                                                                                                                                             P75e/VY2FfXq9aIfRkFaf7Ss9tUWjVon7NH6wsIRxsCxtF5avj6wNkhi4JfacqDkY7gRAWBpzXObmQBH                                                                                                                                                             o7c/OLXWLcwDeS6tgG1OulF9QPnX+n0SzCzSNGJG8gpZZaOs2Ws+G/eee3bXuW9VKuAifRJM+Jsxzy18                                                                                                                                                             7Tdj97WPGW9Es+0rgDOg2DHxR+qt2i7ik0JLo9Gnl6z6C/Gk39cpI3UghkOJFMyUkQtKKLAst5quVsOC                                                                                                                                                             +T2kKiWoDUYI7kttiTpad4X07SLStoW6/29FNvlG26bgufvcOYFlXcy4dsyj8HE5kUQmPY/S8spWThP8                                                                                                                                                             91il3JGZ+WN/abO5TvySFUKmto6def2tfX81oldsRy52kmlUBVxkFBYrtU7QhzP9rIvaugRRWwa/ORCf                                                                                                                                                             jS0GZta5WEothGZpzaDTUVdvPm3G8q84CEX6qAbWjJnJu3H9LfcxVtdEfwIFqjx7y0MCAL8= root@or                                                                                                                                                             acle-VirtualBox
    root@oracle-VirtualBox:~/.ssh# ssh-copy-id oracle@192.168.56.102
    /usr/bin/ssh-copy-id: INFO: Source of key(s) to be installed: "/root/.ssh/id_rsa                                                                                                                                                             .pub"
    /usr/bin/ssh-copy-id: INFO: attempting to log in with the new key(s), to filter                                                                                                                                                              out any that are already installed
    /usr/bin/ssh-copy-id: INFO: 1 key(s) remain to be installed -- if you are prompt                                                                                                                                                             ed now it is to install the new keys
    oracle@192.168.56.102's password:
    
    Number of key(s) added: 1
    
    Now try logging into the machine, with:   "ssh 'oracle@192.168.56.102'"
    and check to make sure that only the key(s) you wanted were added.

    >ssh oracle@192.168.56.102


#6

    oracle@oracle-VirtualBox2:>ssh-keygen -C oracle@oracle-VirtualBox2
    
    oracle@oracle-VirtualBox:>sudo vim /etc/hosts
    set 192.168.56.102 oracle-VirtualBox2

    скопировал открытый ключ с VirtualBox2 на VirtualBox

    oracle@oracle-VirtualBox:~$ ssh oracle-VirtualBox2

###6.1
    
    хост 192.168.56.104
    >cd /root/.ssh/
    >ls -la
    total 38
    drwx------ 2 root root    6 сен 28 23:54 .
    drwx------ 7 root root   13 сен 28 23:54 ..
    -rw-r--r-- 1 root root  576 сен 26 21:59 authorized_keys
    -rw------- 1 root root 2610 сен 28 23:22 id_rsa
    -rw-r--r-- 1 root root  576 сен 28 23:22 id_rsa.pub
    -rw-r--r-- 1 root root  444 сен 28 23:20 known_hosts
    
    >mv id_rsa id_rsa2
    >ls -la
    total 38
    drwx------ 2 root root    6 сен 29 00:18 .
    drwx------ 7 root root   13 сен 28 23:54 ..
    -rw-r--r-- 1 root root  576 сен 26 21:59 authorized_keys
    -rw------- 1 root root 2610 сен 28 23:22 id_rsa2
    -rw-r--r-- 1 root root  576 сен 28 23:22 id_rsa.pub
    -rw-r--r-- 1 root root  444 сен 28 23:20 known_hosts

    >cat /root/.ssh/config
    Host connHost
            Hostname 192.168.56.102
            Port 22
            User oracle
            IdentityFile /root/.ssh/id_rsa2

    >ssh connHost
    Welcome to Ubuntu 20.04.3 LTS (GNU/Linux 5.11.0-36-generic x86_64)
    
     * Documentation:  https://help.ubuntu.com
     * Management:     https://landscape.canonical.com
     * Support:        https://ubuntu.com/advantage
    
    25 updates can be applied immediately.
    To see these additional updates run: apt list --upgradable
    
    Your Hardware Enablement Stack (HWE) is supported until April 2025.
    Last login: Tue Sep 28 23:53:06 2021 from 192.168.56.104

#7

     >tcpdump -w dump0001.pcap -c 100
     >apt install wireshark
     Дамп отрыл на скриншоте task7_1.png

#8
    
    >apt  install nmap

    Открытые порты
    >nmap scanme.nmap.org
    Starting Nmap 7.80 ( https://nmap.org ) at 2021-09-26 23:41 MSK
    Nmap scan report for scanme.nmap.org (45.33.32.156)
    Host is up (0.19s latency).
    Other addresses for scanme.nmap.org (not scanned): 2600:3c01::f03c:91ff:fe18:bb2f
    Not shown: 996 filtered ports
    PORT      STATE SERVICE
    22/tcp    open  ssh
    80/tcp    open  http
    9929/tcp  open  nping-echo
    31337/tcp open  Elite
    
    Nmap done: 1 IP address (1 host up) scanned in 11.98 seconds


    Запущенные сервисы
    >nmap -sV scanme.nmap.org
    Starting Nmap 7.80 ( https://nmap.org ) at 2021-09-26 23:42 MSK
    Nmap scan report for scanme.nmap.org (45.33.32.156)
    Host is up (0.19s latency).
    Other addresses for scanme.nmap.org (not scanned): 2600:3c01::f03c:91ff:fe18:bb2f
    Not shown: 996 filtered ports
    PORT      STATE SERVICE    VERSION
    22/tcp    open  ssh        OpenSSH 6.6.1p1 Ubuntu 2ubuntu2.13 (Ubuntu Linux; protocol 2.0)
    80/tcp    open  http       Apache httpd 2.4.7 ((Ubuntu))
    9929/tcp  open  nping-echo Nping echo
    31337/tcp open  tcpwrapped
    Service Info: OS: Linux; CPE: cpe:/o:linux:linux_kernel
    
    Service detection performed. Please report any incorrect results at https://nmap.org/submit/ .
    Nmap done: 1 IP address (1 host up) scanned in 20.23 seconds

#9

    >ufw status
    Status: inactive

    > ufw enable

    >ufw status
    Status: active
    
    To                         Action      From
    --                         ------      ----
    22/tcp                     ALLOW       Anywhere
    22/tcp                     ALLOW       192.168.56.1
    Apache                     ALLOW       Anywhere
    22/tcp (v6)                ALLOW       Anywhere (v6)
    Apache (v6)                ALLOW       Anywhere (v6)

    >ufw status numbered
    Status: active
    
         To                         Action      From
         --                         ------      ----
    [ 1] 22/tcp                     ALLOW IN    Anywhere
    [ 2] 22/tcp                     ALLOW IN    192.168.56.1
    [ 3] Apache                     ALLOW IN    Anywhere
    [ 4] 22/tcp (v6)                ALLOW IN    Anywhere (v6)
    [ 5] Apache (v6)                ALLOW IN    Anywhere (v6)

    >ufw delete 5
    >ufw delete 4
    >ufw delete 3
    >ufw delete 1

    оставил 22 порт с хоста 192.168.56.1, чтобы ssh не отвалилось
    >ufw status
    Status: active
    
    To                         Action      From
    --                         ------      ----
    22/tcp                     ALLOW       192.168.56.1
    
    открываем доступ по 80 порту
    >ufw allow 80/tcp comment 'Apache'
    
    открываем доступ по 443 порту    
    >ufw allow 443/tcp comment 'Apache SSL'

    открываем доступ по 22 порту c сервера 192.168.56.102    
    >ufw allow from 192.168.56.102 to any port 22 proto tcp


    >ufw status
    Status: active
    
    To                         Action      From
    --                         ------      ----
    22/tcp                     ALLOW       192.168.56.1
    80/tcp                     ALLOW       Anywhere                   # Apache
    443/tcp                    ALLOW       Anywhere                   # Apache SSL
    22/tcp                     ALLOW       192.168.56.102
    80/tcp (v6)                ALLOW       Anywhere (v6)              # Apache
    443/tcp (v6)               ALLOW       Anywhere (v6)              # Apache SSL
