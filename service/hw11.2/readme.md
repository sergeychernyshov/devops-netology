# Задача 1: API Gateway

Предложите решение для обеспечения реализации API Gateway. Составьте сравнительную таблицу возможностей различных программных решений. На основе таблицы сделайте выбор решения.

Решение должно соответствовать следующим требованиям:

Маршрутизация запросов к нужному сервису на основе конфигурации
Возможность проверки аутентификационной информации в запросах
Обеспечение терминации HTTPS
Обоснуйте свой выбор.

### Ответ

[Ссылка на источник](https://www.nginx.com/blog/choosing-the-right-api-gateway-pattern/?__cf_chl_captcha_tk__=ed4062235a4eed746839ffaa4456039c320f9dec-1624120905-0-AYcP8ha43drdyJNKyF-DIGvZ1DR9YkX22Cl4O1eqj6VlImek5TI0LGefDCYx2x8nlce-REfCRIUyoqCAkDCcqNbjEPJCtKlCWPqty1WO-_5wzjjqkhHOIIP4GOQc-WRiQC4-XSYDvk_4svL_UsxTXQXtnach30UaMfB5o7Gf1AWSQtMK4FuNSBTxryzUZ4JAFyVrZVzzUWvWihqdKzk9otDOoK_kSzkjOTisu3XqYuBH6SsaiivUTx2TE20_hYQk4_URyMlzH7RhybZzZQL_KZakNAx2Aa5rBJt-ps0ETa_PZsEOgBBOv2MFeQu0DzaSqiLqnmboU47i38a0dKZg23_sq2wz7ULK3JZTSKsj6fIO-PTA-k_yncNs7odQNoL-hbaUsWor_aJX3cTcoDvok11rxniKHSQtkTKaGKgmi9s4SNft7kq6u0fkANzo77qLQQ7RJuwEAb5QuRba8ulZPQADD8vVC16JyY3aRET9DskxQg2T-1n6MGHjOtvjrNjhUojPMI0VU4oFjU7YqCvJyFdW-r4b43R2D3S6uOeJVNZjn4DbU-lKshHet4jd45PjgA0C7fANadJ9bstolosruifQyN_rWlRJiytg-yRLpZpxkFKEhRuC3C6R9JVbivKhP1gSc2OkLvKH8Fa4wk6QwIr03PSzA78Dnsb4JSE5MhvqaK_0XYqb95ZWYUKaFF40Jw)

|Критерий| Центральный шлюз|Two-tier gateway|Microgateway|Per-pod|Sidecar, Service Mash|
|:---:|:---:|:---:|:---:|:---:|:---:|
|SSL/TLS termination          |+|+|+|+|+|
|Authentication               |+|+|+|+|+|
|Authorization                |+|+| | |+|
|Request routing              |+|+|+|+| |
|Rate limiting                |+|+|+|+| |
|Request/response manipulation|+| | | | |
|Facade routing               | |+| | |+|
|Load balancing               | |+|+| |+|
|                             |Подходит для монолита, не подходит при частых изменений|Не поддерживает распределенное управление,подходит для обеспечения гибкости|Подходит для отдельных команд разработки,может управлять трафиком между службами сложно добиться согласованности и контроля|Применяется при использовании с одним из предыдущих 3х решений. Использует прокси-сервер. Конфигурация статична, не требует модификации при изменении приложения|Сложность в управлении


Предполагаю что в нашем случае можно использовать Microgateway и Gateway Per-pod.
Эти шлюзы имеют маршрутизацию и позволяют обеспечить автономность для каждой команды разработки.


# Задача 2: Брокер сообщений

Составьте таблицу возможностей различных брокеров сообщений. На основе таблицы сделайте обоснованный выбор решения.

Решение должно соответствовать следующим требованиям:

Поддержка кластеризации для обеспечения надежности
Хранение сообщений на диске в процессе доставки
Высокая скорость работы
Поддержка различных форматов сообщений
Разделение прав доступа к различным потокам сообщений
Простота эксплуатации
Обоснуйте свой выбор.

### Ответ

[Ссылка на источник](https://www.okbsapr.ru/library/publications/shkola_kzi_chadov_mikhalchenko_2019/)

|Критерий|RabbitMQ|ActiveMQ|Qpid C++|SwiftMQ|Artemis|Apollo|Kafka|
|:---:|:---:|:---:|:---:|:---:|:---:|:---:|:---:|
|Кластеризация                          |+|+|+|+|+|-|+|
|Хранение сообщений                     |+|+|+|+|+|+|+|
|Скорость работы                        |+|-|-|-|+-|-|+|
|Поддержка различных форматов сообщений |+|-|-|-|-|-|+|
|Разделение прав доступа                |+|+|+|+|+|+|+|
|Простота эксплуатации                  |+|-|-|-|-|-|+|
 
RabbitMQ и Kafka более приоритетны в выборе.
RabbitMQ имеет меньше возможностей чем Kafka, но проще в администрировании.

Также можно рассмотреть Kubemq для работы в KUBERNETES.


# Задача 3: API Gateway * (необязательная)

Ссылка на файл [nginx.conf](nginx.conf)

Скачивать файл по такой ссылке
````
curl -O localhost/image/5115748a-08ee-4743-b768-3fb628fe5e46.jpg
````


С частью по скачиванию файла 2 ошибки в постановки задачи.

1) в задании указано

GET /v1/user/{image}

Проверка токена. Токен ожидается в заголовке Authorization. Токен проверяется через вызов сервиса security GET /v1/token/validation/
Запрос направляется в сервис minio GET /images/{image}

А далее проверка метода 

curl -X GET http://localhost/images/4e6df220-295e-4231-82bc-45e4b1484430.jpg

Без параметров запроса в заголовке.

2) в дополнительном материале указано, что проверить загрузку можно

curl localhost/image/<filnename> > <filnename>

а сразу ниже пример

````
$ curl localhost/images/c31e9789-3fab-4689-aa67-e7ac2684fb0e.jpg > c31e9789-3fab-4689-aa67-e7ac2684fb0e.jpg
````

image или images правильно не понятно.
Я выбрал вариант image

````
curl -O localhost/image/5115748a-08ee-4743-b768-3fb628fe5e46.jpg
````
