#1

    c = a+b
    d = 1+2
    e = 3

    По умолачанию в bash все переменные имеют тип строка по этому
    a+b это срока
    
    $a+$b это строка, в которой $a это знчение переменной а, $b значение переменной b
    поэтому мы получаем 1+2

    $(($a+$b)) это вычесление выражения $a+$b, т.е. 1+2 = 3

#2
    
    Мой вариант скрипта остановится, когда сервис станет доступен

    while [ 1==1 ]
    do 
    curl http://192.168.56.104:19999
    if (($? != 0))
    then 
    date >> curl.log 
    else 
    break
    fi 
    done

#3

    hosts=("192.168.0.1:80" "173.194.222.113:80" "87.250.250.242:80")
    while [ 1==1 ]
    do
    for host in "${hosts[@]}"
    do
    for i in 1 2 3 4 5
    do
    echo $host 
    curl $host 
    if (($? != 0))
    then
    echo service not work>>curl.log
    echo $host>> curl.log
    date>>curl.log
    fi
    done
    done
    done    

#4

    hosts=("192.168.0.1:80" "173.194.222.113:80" "87.250.250.242:80")
    while [ 1==1 ]
    do
    for host in "${hosts[@]}"
    do
    echo $host 
    curl $host 
    if (($? != 0))
    then
    echo service not work>>error.log
    echo $host>> error.log
    date>>error.log
    break 2
    fi 
    done
    done   

#*5
    prepare-commit-msg
    Текст коммита должен начинаться с [99_script-99-bash] далее может идти любой текст
    Не более 30 символов

    #!/bin/bash
    commitRegex='^\[[0-9]{2}-script-[0-9]{2}-bash\]*.+'
    if ! grep -qE "$commitRegex" "$1"; then
        echo "Aborting according commit message policy. Please specify start with mask:[99-script-99-bash] text"
        exit 1
    fi
     
    message=`cat $1`;  
    length=${#message} 
    
    if  (($length>30)) 
    then
        echo "string must be less than 30 characters"
        exit 1
    fi
