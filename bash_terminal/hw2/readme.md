#1.

    Команда cd типа stdin (standard input) так как она не возвращает результата.
    Команда может быть типа stdout, если будет возвращать результат перехода в каталог.
    Например, в случае успеха: Директория изменена.
    В случае не успеха: Указанной директории не существует. 
    
#2.

    Альтернатива команде grep <some_string> <some_file> | wc -l будет команда
    grep <some_string> <some_file> -c

#3.

    ps -A
    процесс с PID = 1 называется systemd

#4.

    В первом терминале выполнить
    vagrant@vagrant:~$ tty

    /dev/pts/1
    /dev/pts/1 имя терминала
    
    Во втором терминале выполнить
    ls -test>/dev/pts/1 2>&1

    мы перенаправляем поток stderr в stdout и уже stdout перенаправляем в нужный терминал /dev/pts/1

#5.

    Да, получится если использовать команду tee, например
    cat file_for_stdin.txt | tee file_output.txt

#6.

    Получилось вывести ошибку из PTY в TTY терминал.
    В PTY выполнил команду
    ls -test>/dev/tty1 2>&1
    
    в TTY терминале получил ошибку
    vagrant@vagrant:~$ ls: invalid option -- 'e'
    Try 'ls --help' for more information.

#7.
    
    Выполнение команды bash 5>&1 приведет к созданию файлового дескриптора с идентификатором 5, 
    который направляет вывод в стандартный поток вывода.
    
    По сути мы создали файл для хранения выходных данных.
    Команда 
    echo netology > /proc/$$/fd/5
    запишет в файл строку netology, а так как мы отправили вывод дескриптора 5 в стандартный поток,
    мы увидим на экране строку netology.

#8.
    Команды выводят содержимое в стандартный поток, при этом содержимое файла пустое
    bash 4>&1
    ls -la 2>&1 >/proc/$$/fd/4 | cat>file_out.text
    
    команды выводят ошибку в файл
    bash 4>&1
    ls -test 2>&1 >/proc/$$/fd/4 | cat>file_out.text
    
#9.

    Переменная environment указывает на массив указателей, которые называются «окружением».
    Аналогичный вывод можно получить командой
    env

#10.

    /proc/<PID>/cmdline В этом файле хранится командная строка, которой был запущен данный процесс.
    /proc/<PID>/exe представляет собой символическую ссылку на исполняемый файл, который инициировал запуск процесса.

#11.
    
    Не понимаю как читать иформацию о информацию, по этому копирую данные всего атрибута

    flags : fpu vme de pse tsc msr pae mce cx8 apic sep mtrr pge mca cmov pat pse36 clflush mmx fxsr 
    sse sse2 ht syscall nx rdtscp lm constant_tsc rep_good nopl xtopology nonstop_tsc cpuid tsc_known_freq pni 
    pclmulqdq ssse3 cx16 pcid sse4_1 sse4_2 x2apic movbe popcnt aes xsave avx rdrand hypervisor lahf_lm abm 
    3dnowprefetch invpcid_single fsgsbase avx2 invpcid rdseed clflushopt md_clear flush_l1d arch_capabilities
    
#12.

    при подключении ожидается пользователь, а не другой процесс, и нет локального tty в данный момент. 
    для запуска можно добавить -t - , и команда исполняется c принудительным созданием псевдо терминала
    ssh -t localhost 'tty'
    
#13.

    В первом терминале запустил в screen процесс
    screen sudo apt install postgresql postgresql-contrib

    Во втором терминале нашел процесс
    ps -aux | grep "postgre"
    
    перехватил выполнение командой
    sudo reptyr -T 5286

#14.

    tee создает pipe для записи в файл, при этом стандартный поток вывода также будет работать.
    Через pipe передается только stdout команды слева от | на stdin команды справа,
    т.е. у команды справа не будет прав для записи в файл, который принадлежит пользователю root.
    sudo tee /root/new_file позволит в отдельном потоке записать данные в файл.
    
    



