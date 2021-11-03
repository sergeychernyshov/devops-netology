#1

###Полная (аппаратная) виртуализация

    В этом случае гипервизор (разделяющий ресурсы между виртуальными машинами) является самостоятельной
    операционной системой и не требует установки хостовой ОС.
    Гипервизор напрямую взаимодействует с железом сервера.

###Паравиртуализация

    Требует наличия хостовой ОС и также требует модификации ядра гостевых ОС, таким образом чтобы
    гостестевая ОС была встроена в ОС хоста. Гостевая ОС взаимодействует с гипервизором,
    который предоставляет гостевой ОС API для пользования ресурсами.

###Виртуализация на основе ОС

    В этом случае происходит встраивания контейнера с гостевой ОС. Контейнеры изолированы друг от друга.
    Контейнеры используют ядро хостовой ОС. Это ограничивает набор готовых ОС для установки. Можно развернуть
    только те гостевые ОС, которые имеют одинаковое ядро с хостовой ОС.
    Хостовая ОС выполняет ф-ции гиппервизора распределяя ресурсы.
    Гостевые ОС являются эмуляцией пользовательского окружения, а для пользователя выглядят как отдельный сервер.

#2

###Высоконагруженная база данных, чувствительная к отказу

    В этом случае использую физический сервер.
    Дополнительные слои в виде гиппервизора являются некими посредниками для получения ресурсов сервера -
    это замедляет работу БД. Размещение БД на отдельном сервере создает прогнозируемый отклик БД.
    Также гиппервизоры(контейнеры) это слой который может служить дополнительным источником отказов и ошибок.
    
    Рассмотрел вариант с паравиртуализацией. Для удобства администрирования сервера БД:
    простого создание резервной копии(в случае если БД делает плохо бекапы, требуется даунтайм для бекапа или вообще не умеет),
    легкого переноса БД в последующим на другие сервера или в облако.

###Различные web-приложения

    Виртуализация уровня ОС (контейнеры)
    Как правило web приложения решают небольшие точечные задачи, поэтому легко спрогнозировать какие ресурсы
    потребуются приложению для работы на определенное количество пользователей. Далее контейнеры легко масштабировать
    под необходимую нагрузку.

###Windows системы для использования бухгалтерским отделом

    В этом случае предполагаю, что основным критерием будет простота бекапирования системы и минимальное время поднятия
    системы из бекапа в случае сбоя. И т.к. это Windows системы, то хочется склониться к варианту конкретной
    виртуализации Hyper-V. Как я понимаю этот гиппервизор может работать в двух режимах:
    аппаратная виртуализация и паравиртуализация.

    Думаю, что выбрал тип паравиртуализации, т.к. в этом случае использование ресурсов происходит более рационально.

###Системы, выполняющие высокопроизводительные расчеты на GPU

    Предполагаю, что в случае сильной нагрузки на железо сервера, будет плохим вариантом заставлять работать через
    посредника в виде гипервизора. В этом случае использую физический сервер.

#3

    у меня нет опыта с вирутализацией, дать развернутый ответ мне будет трудно

###100 виртуальных машин на базе Linux и Windows, общие задачи, нет особых требований

    В этом случае выбор достаточно большой. Все должно зависеть от компетенций сотрудников.
    Там где админы больше разбираются то и нужно выбирать.
    VMWare,Hyper-V,KVM,Xen подойдет для решения.

###Требуется наиболее производительное бесплатное open source решение для виртуализации

    приоритетнее KVM - гостевые ос могут быть любыми, бесплатен и вполне производителен.
    Можно рассмотреть Xen, но он менее производителен.

###Необходимо бесплатное, максимально совместимое и производительное решение для виртуализации Windows инфраструктуры.

    В данном случае попробую использовать бесплатный Hyper-V, с Windows должен быть максимально совместим.
    Не знаю насколько производительное решение сделал Microsoft, но очень сильно надеюсь на то что их решение
    будет обладать хорошей производительностью в работе с Windows.

###Необходимо рабочее окружение для тестирования программного продукта на нескольких дистрибутивах Linux.

    Тут я бы пошел самым простым путем: поставил VirtualBox и vagrant.
    Такое решение можно развернуть на сервере для совместного пользования так и поставить у себя на локальной машине.
    Это будет удобнее для тестирования.

#4

    1) Разные решения виртуализации обладают своими преимуществами и недостатками. Для достижения удобства и
    производительность имеет смысл рассматривать разные системы виртуализации для разных задач. Если новый тип
    виртуализации будет лучше соответствовать критическим требованиям системы, то нужно использовать другой тип.
    Например, если надежность важнее производительности то нужно выбирать надежное решение. Если важна производительность 
    то соответственно выбор будет в наиболее производительной системе, а надежность будет обеспечиваться
    через масштабирование.
    Буду использовать разные системы виртуализации.
    
    2) Проблемы использования нескольких систем виртуализации:
    - нужно содержать несколько сотрудников на каждую систему виртуализации или обучать сотрудников для
    работы в разных системах
    - если prod/test/dev в разных системах виртуализации, то отлавливать "баги" сложнее, трудно поддерживать
    схожесть ландшафтов test/dev соответствующих prod среде.
    - проблемы миграции между разными системами виртуализации.
    
    3) Сложно сказать что делать для минимизации рисков гетерогенной среды виртуализации, но попробую:
    - обучать админов работать с разными системами.
    - не использовать напрасно разные системы виртуализации, все решение должны быть обдуманы и иметь доказательства
    превосходства в той или иной ситуации одного типа виртуализации над другой.
    Далее что-то сказать сложно...