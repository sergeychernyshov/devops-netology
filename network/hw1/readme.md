#1
    telnet stackoverflow.com 80
    Trying 151.101.65.69...
    Connected to stackoverflow.com.
    Escape character is '^]'.
    GET /questions HTTP/1.0
    HOST: stackoverflow.com
    
    HTTP/1.1 301 Moved Permanently
    cache-control: no-cache, no-store, must-revalidate
    location: https://stackoverflow.com/questions
    x-request-guid: ef0ca6fe-1bbd-477a-ae45-6a7da6ca3d4c
    feature-policy: microphone 'none'; speaker 'none'
    content-security-policy: upgrade-insecure-requests; frame-ancestors 'self' https://stackexchange.com
    Accept-Ranges: bytes
    Date: Mon, 06 Sep 2021 19:37:45 GMT
    Via: 1.1 varnish
    Connection: close
    X-Served-By: cache-ams21075-AMS
    X-Cache: MISS
    X-Cache-Hits: 0
    X-Timer: S1630957065.312267,VS0,VE74
    Vary: Fastly-SSL
    X-DNS-Prefetch-Control: off
    Set-Cookie: prov=23d7f690-3331-657d-f3c2-830537b9abd3; domain=.stackoverflow.com; expires=Fri, 01-Jan-2055 00:00:00 GMT; path=/; HttpOnly

    получил 301 код: страница переадресовала меня с 80 на 443 порт, т.е. с HTTP на HTTPS.
    301 Moved Permanently - постоянный редирект

#2
    Google Chrome показывает 307 редирект (task2_screen_1.png).
    Однако, Link Redirect Trace подсказвает следующее(task2_screen_2.png):
    
    `The server has previously indicated this domain should always be accessed via HTTPS 
    (HSTS Policy per https://tools.ietf.org/html/rfc6797). Chrome has cached this internally, 
    and did not connect to any server for this redirect. Chrome reports this redirect as a "307 Internal Redirect" 
    which simply does not exist per https://tools.ietf.org/html/rfc7231#section-6.4.7 - 
    however this probably would have been a "301 Permanent redirect" originally and the Google guys made 
    fun of the webmaster community maybe. You can verify this by clearing your browser cache and visiting 
    the original URL again. Please note that this is kind of a weird behavior and that Google even calls 307 
    redirects "a lie" in a post by John Muller titled "A search-engine guide to 301, 302, 307, 
    & other redirects" at https://plus.google.com/+JohnMueller/posts/E4PqAhRJB2V - 
    However server side 307 redirects do exist and we will show them. ;-) - for further details on 
    redirects we recommend you checkout the CEMPER.Academy or the LinkResearchTools LRT Associate Training 
    as all those details are trained in there, and covered in the full LinkResearchTools suite of course.`

    ну и шуточки у Google.

    Браузер ATOM показал 301 редирект (task2_screen_3.png).

    307 редирект является временным перенаправлением

    
    самый долгий запрос на странице 430мс - получение html страницы

#3
    мой ip адрес в интернете 194.15.119.144 (task3_screen_1.png)
    проверял на сайте 2ip.ru

#4
    whois 194.15.119.144
    % This is the RIPE Database query service.
    % The objects are in RPSL format.
    %
    % The RIPE Database is subject to Terms and Conditions.
    % See http://www.ripe.net/db/support/db-terms-conditions.pdf
    
    % Note: this output has been filtered.
    %       To receive output for a database update, use the "-B" flag.
    
    % Information related to '194.15.116.0 - 194.15.119.255'
    
    % Abuse contact for '194.15.116.0 - 194.15.119.255' is 'noc@svsreut.ru'
    
    inetnum:        194.15.116.0 - 194.15.119.255
    netname:        SVS-TELECOM-NET
    country:        RU
    org:            ORG-LCL15-RIPE
    admin-c:        KK6450-RIPE
    tech-c:         KK6450-RIPE
    status:         ASSIGNED PI
    mnt-by:         RIPE-NCC-END-MNT
    mnt-by:         LEKSTAR-COMMUNICATION-MNT
    mnt-routes:     LEKSTAR-COMMUNICATION-MNT
    mnt-domains:    LEKSTAR-COMMUNICATION-MNT
    created:        2011-04-18T07:12:02Z
    last-modified:  2019-04-03T08:09:38Z
    source:         RIPE # Filtered
    sponsoring-org: ORG-CS216-RIPE
    
    organisation:   ORG-LCL15-RIPE
    org-name:       Lekstar Communication Ltd.
    org-type:       OTHER
    address:        Russia, Reutov, Yubileyny prospekt, 66
    abuse-c:        ACRO1187-RIPE
    mnt-ref:        LEKSTAR-COMMUNICATION-MNT
    mnt-by:         LEKSTAR-COMMUNICATION-MNT
    mnt-by:         MNT-NETUP
    created:        2016-09-19T08:39:44Z
    last-modified:  2018-02-22T14:39:31Z
    source:         RIPE # Filtered
    
    person:         Konstantin Kalinichenko
    address:        Russia, Reutov, Yubileyny prospekt, 66
    phone:          +7 903-566-09-93
    nic-hdl:        KK6450-RIPE
    mnt-by:         LEKSTAR-COMMUNICATION-MNT
    created:        2016-09-19T08:06:44Z
    last-modified:  2018-02-22T14:38:00Z
    source:         RIPE
    
    % Information related to '194.15.116.0/22AS49261'
    
    route:          194.15.116.0/22
    descr:          SVS-Telecom Ltd.
    descr:          Ru
    origin:         AS49261
    mnt-by:         MNT-SVS-TELECOM
    created:        2011-04-19T10:11:11Z
    last-modified:  2011-04-19T10:11:11Z
    source:         RIPE
    
    % This query was served by the RIPE Database Query Service version 1.101 (HEREFORD)
    
    
    ip адрес принадлежит провайдеру SVS-Telecom Ltd. (Lekstar Communication Ltd.)
    автономной системе AS49261

#5

    на WINDOWS
    tracert 8.8.8.8

    Трассировка маршрута к dns.google [8.8.8.8]
    с максимальным числом прыжков 30:
    
      1     1 ms     1 ms     1 ms  XiaoQiang [192.168.31.1]
      2     3 ms     3 ms     3 ms  10.172.93.1
      3     3 ms     5 ms    18 ms  172.16.10.1
      4     *        *        *     Превышен интервал ожидания для запроса.
      5     5 ms     5 ms     4 ms  91.211.105.62
      6     5 ms     5 ms     5 ms  108.170.250.33
      7     5 ms     7 ms     5 ms  108.170.250.34
      8    22 ms     *       20 ms  172.253.66.116
      9    22 ms    22 ms    32 ms  172.253.65.159
     10    20 ms    19 ms    20 ms  209.85.254.179
     11     *        *        *     Превышен интервал ожидания для запроса.
     12     *        *        *     Превышен интервал ожидания для запроса.
     13     *        *        *     Превышен интервал ожидания для запроса.
     14     *        *        *     Превышен интервал ожидания для запроса.
     15     *        *        *     Превышен интервал ожидания для запроса.
     16     *        *        *     Превышен интервал ожидания для запроса.
     17     *        *        *     Превышен интервал ожидания для запроса.
     18     *        *        *     Превышен интервал ожидания для запроса.
     19     *        *        *     Превышен интервал ожидания для запроса.
     20    21 ms    21 ms    21 ms  dns.google [8.8.8.8]
    
    на виртуалке UBUNTU  выполнена команда 
    traceroute 8.8.8.8 -A
    (task5_screen_1.png)

#6  
    выполнил команду
    mtr -nc 1000 --psise 1000 8.8.8.8
    результат task6_screen_1.png
    
    проблемы в подсетях 172.253.*.* и 209.85.*.*

#7  
    > dig dns.google any

    ; <<>> DiG 9.16.1-Ubuntu <<>> dns.google any
    ;; global options: +cmd
    ;; Got answer:
    ;; ->>HEADER<<- opcode: QUERY, status: NOERROR, id: 21872
    ;; flags: qr rd ra; QUERY: 1, ANSWER: 10, AUTHORITY: 0, ADDITIONAL: 1
    
    ;; OPT PSEUDOSECTION:
    ; EDNS: version: 0, flags:; udp: 65494
    ;; QUESTION SECTION:
    ;dns.google.                    IN      ANY
    
    ;; ANSWER SECTION:
    dns.google.             344     IN      RRSIG   AAAA 8 2 900 20211008160924 2021                                0908160924 1773 dns.google. W+SJL2jsOfj1fuvOZjMy6tWgKjMm/298xppqqmO+f2FeUMLGrb1X                                Ef2D J36y9Oild26xFpOdvGRCHy0THFvO8WnZQnv8Qb7g0bmwScyva1LWdajZ 6h9ppJKbep3jAjXagB                                A7DHYvCH5s6TBT6HDBb2UxyRbOmHq+1dF7pwQO K7M=
    dns.google.             344     IN      AAAA    2001:4860:4860::8844
    dns.google.             344     IN      AAAA    2001:4860:4860::8888
    dns.google.             1611    IN      NS      ns3.zdns.google.
    dns.google.             1611    IN      NS      ns2.zdns.google.
    dns.google.             1611    IN      NS      ns4.zdns.google.
    dns.google.             1611    IN      NS      ns1.zdns.google.
    dns.google.             765     IN      RRSIG   A 8 2 900 20211008160924 2021090                                8160924 1773 dns.google. TgNWkcjNAgttxafVK9iMX0kBHxLRQPwlK5cudnemTZtPvJdx+tBGxO1                                g RrFhYqPG8esw3jIFO31k61eazPL3fhhlA87unQrZ8DLr0HqG5lga/AVj PevY0qZOhVGWHsAIn4jgf                                WNd6CNQ2vW63h5L+zMkv97wnf6MT+V+R1+q zsk=
    dns.google.             765     IN      A       8.8.4.4
    dns.google.             765     IN      A       8.8.8.8
    
    ;; Query time: 12 msec
    ;; SERVER: 127.0.0.53#53(127.0.0.53)
    ;; WHEN: Ср сен 08 18:05:41 MSK 2021
    ;; MSG SIZE  rcvd: 544


    NS - доменные сервера
    A – адрес ipv4

#8
    dig -x 2001:4860:4860::8844 

    ;; QUESTION SECTION:
    ;4.4.8.8.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.6.8.4.0.6.8.4.1.0.0.2.ip6.arpa. IN PTR
    
    ;; ANSWER SECTION:
    4.4.8.8.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.6.8.4.0.6.8.4.1.0.0.2.ip6.arpa. 7056 IN PTR dns.google.

    dig -x 2001:4860:4860::8888 
    
    ;; QUESTION SECTION:
    ;8.8.8.8.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.6.8.4.0.6.8.4.1.0.0.2.ip6.arpa. IN PTR
    
    ;; ANSWER SECTION:
    8.8.8.8.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.6.8.4.0.6.8.4.1.0.0.2.ip6.arpa. 6903 IN PTR dns.google.

    dig -x ns3.zdns.google

     QUESTION SECTION:
    ;google.zdns.ns3.in-addr.arpa.  IN      PTR

    dig -x ns2.zdns.google.

    ;; QUESTION SECTION:
    ;google.zdns.ns2.in-addr.arpa.  IN      PTR

    dig -x ns4.zdns.google

    ;; QUESTION SECTION:
    ;google.zdns.ns4.in-addr.arpa.  IN      PTR


    dig -x ns1.zdns.google 

    ;; QUESTION SECTION:
    ;google.zdns.ns1.in-addr.arpa.  IN      PTR


    dig -x 8.8.4.4    

    ;; QUESTION SECTION:
    ;4.4.8.8.in-addr.arpa.          IN      PTR

    ;; ANSWER SECTION:
    4.4.8.8.in-addr.arpa.   30      IN      PTR     dns.google.

    dig -x 8.8.8.8              
    
    ;; QUESTION SECTION:
    ;8.8.8.8.in-addr.arpa.          IN      PTR

    ;; ANSWER SECTION:
    8.8.8.8.in-addr.arpa.   30      IN      PTR     dns.google.

    к IP адресам привязано доменное имя dns.google
    




    

    



    