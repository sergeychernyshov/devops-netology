#1

    chdir("/tmp") - команда смены каталога

#2
    
    Информацию о файлах, содержащих сигнатуру, т.е. числовую или строковую константу, задающую тип, команда file получает из файлов:
        /home/vagrant/.magic.mgc
        /etc/magic.mgc
        /usr/share/misc/magic.mgc

#3
    
    Запустил в фоне процесс
    ping localhost>test_delete.txt &
    запустил процесс для просмотра
    tail -f  test_delete.txt
    
    из другого сеанса удалил файл
    rm test_delete.txt
    смотрю статус файла 
    lsof | grep test_delete.txt
    tail      2172                        vagrant    3r      REG              253,0     1893     191091 /home/vagrant/test_delete.txt (deleted)

    ищу файловый дескриптор
    find /proc/2172/fd -ls 2> /dev/null | grep test_delete.txt
    41153      0 lr-x------   1 vagrant  vagrant        64 Aug 21 13:45 /proc/2172/fd/3 -> /home/vagrant/test_delete.txt\ (deleted)
    
    удаляю содержимое дескриптора
    truncate -s 0 /proc/2172/fd/3

#4
    
    Процесс "зомби" не занимают ресурсы, но блокируют записи в таблице процессов, 
    размер которой ограничен для каждого пользователя и системы в целом. 

#5
    
    Команда /usr/sbin/opensnoop-bpfcc запустилать только под пользователм root
    root@vagrant:~# /usr/sbin/opensnoop-bpfcc
    815    vminfo              4   0 /var/run/utmp
    629    dbus-daemon        -1   2 /usr/local/share/dbus-1/system-services
    629    dbus-daemon        18   0 /usr/share/dbus-1/system-services
    629    dbus-daemon        -1   2 /lib/dbus-1/system-services
    629    dbus-daemon        18   0 /var/lib/snapd/dbus-1/system-services/
    427    systemd-udevd      14   0 /sys/fs/cgroup/unified/system.slice/systemd-udevd.service/cgroup.procs
    427    systemd-udevd      14   0 /sys/fs/cgroup/unified/system.slice/systemd-udevd.service/cgroup.threads
    637    irqbalance          6   0 /proc/interrupts
    637    irqbalance          6   0 /proc/stat
    637    irqbalance          6   0 /proc/irq/20/smp_affinity
    637    irqbalance          6   0 /proc/irq/0/smp_affinity
    637    irqbalance          6   0 /proc/irq/1/smp_affinity
    637    irqbalance          6   0 /proc/irq/8/smp_affinity
    637    irqbalance          6   0 /proc/irq/12/smp_affinity
    637    irqbalance          6   0 /proc/irq/14/smp_affinity
    637    irqbalance          6   0 /proc/irq/15/smp_affinity

    Показаны системные демоны linux.

#6
    
    Команда user -a использует системный вызов uname()
    
    uname({sysname="Linux", nodename="vagrant", ...}) = 0
    write(1, "Linux vagrant 5.4.0-80-generic #"..., 105Linux vagrant 5.4.0-80-generic #90-Ubuntu SMP Fri Jul 9 22:49:44 UTC 2021 x86_64 x86_64 x86_64 GNU/Linux
    ) = 105
    
    и man 
        Part of the utsname information is also accessible via
       /proc/sys/kernel/{ostype, hostname, osrelease, version,
       domainname}

#7
    
    && -  условный оператор, 
    ;  - разделитель последовательных команд

    set -e выход при любом не нулевом статусе команды.
    Предполагаю, что && set -e не имеет смысла

#8

    set -euxo pipefail
    
    -e  Exit immediately if a command exits with a non-zero status. (Немедленно выйти, если команда завершается с ненулевым статусом)
    -v  Print shell input lines as they are read. (Печатать строки ввода оболочки по мере их чтения)
    -x  Print commands and their arguments as they are executed.(Печатать команд и аргументов по мере их выполнения)
    -o pipefail the return value of a pipeline is the status of
                           the last command to exit with a non-zero status,
                           or zero if no command exited with a non-zero status
    (возвращаемое значение - статус последней команды для выхода с ненулевым статусом или ноль, если все с нулевым статусом)

    Использование такой команды повышает детализацию логирования.

#9

    vagrant@vagrant:/root$ ps -o stat
    STAT
    Ss
    S
    S
    R+

    S	Interruptible sleep (waiting for an event to complete) - процессы ожидающие завершения
    s	is a session leader (руководитель сессии)
    R	Running or runnable (on run queue) (запущенные процессы или процессы в очереди)
    +	is in the foreground process group (находится в группе процессов переднего плана)