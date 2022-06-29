Установка kubeadm.

apt-get update

apt-get install -y apt-transport-https ca-certificates curl

curl -fsSLo /usr/share/keyrings/kubernetes-archive-keyring.gpg https://packages.cloud.google.com/apt/doc/apt-key.gpg

echo "deb [signed-by=/usr/share/keyrings/kubernetes-archive-keyring.gpg] https://apt.kubernetes.io/ kubernetes-xenial main" | sudo tee /etc/apt/sources.list.d/kubernetes.list

sudo apt-get install -y kubelet kubeadm kubectl

snap install k9s


# Задание 1: Запуск пода из образа в деплойменте
Для начала следует разобраться с прямым запуском приложений из консоли. Такой подход поможет быстро развернуть инструменты отладки в кластере. Требуется запустить деплоймент на основе образа из hello world уже через deployment. Сразу стоит запустить 2 копии приложения (replicas=2).

### Требования:

* пример из hello world запущен в качестве deployment
* количество реплик в deployment установлено в 2
````
kubectl create deployment hello-node --image=k8s.gcr.io/echoserver:1.4 -r 2
deployment.apps/hello-node created
````

* наличие deployment можно проверить командой kubectl get deployment

````
kubectl get deployment
NAME         READY   UP-TO-DATE   AVAILABLE   AGE
hello-node   2/2     2            2           52s
````

* наличие подов можно проверить командой kubectl get pods

````
kubectl get pods
NAME                          READY   STATUS    RESTARTS   AGE
hello-node-6b89d599b9-7xj8w   1/1     Running   0          79s
hello-node-6b89d599b9-8gbwj   1/1     Running   0          79s
````

# Задание 2: Просмотр логов для разработки

Разработчикам крайне важно получать обратную связь от штатно работающего приложения и, еще важнее, об ошибках в его работе. Требуется создать пользователя и выдать ему доступ на чтение конфигурации и логов подов в app-namespace.

Требования:

* создан новый токен доступа для пользователя

* * создаем неймспейс 
````
kubectl create namespace app-namespace
namespace/app-namespace created

kubectl create deployment hello-node --image=k8s.gcr.io/echoserver:1.4 --namespace=app-namespace
deployment.apps/hello-node created
````

* * создаем каталог пользователя
    
````
mkdir developer
cd developer    
````

* * генерируем ключ
    
````
openssl genrsa -out developer.key 2048
Generating RSA private key, 2048 bit long modulus (2 primes)
..............................+++++
.....................+++++
e is 65537 (0x010001)

openssl req -new -key developer.key -out developer.csr -subj "/CN=developer"
````

````
openssl x509 -req -in developer.csr -CA /root/.minikube/ca.crt -CAkey /root/.minikube/ca.key -CAcreateserial -out developer.crt -days 365
Signature ok
subject=CN = developer
Getting CA Private Key
````

* пользователь прописан в локальный конфиг (~/.kube/config, блок users)

````
kubectl config set-credentials developer --client-certificate=/home/developer/developer.crt --client-key=/home/developer/developer.key
User "developer" set.
````

* * создаю контекст

````
kubectl config set-context app-namespace-developer --namespace=app-namespace --cluster=minikube --user=developer
Context "app-namespace-developer" created.
````

* *ключ и сертификат

````
ls -la
total 20
drwxr-xr-x 2 root root 4096 июн 29 22:48 .
drwxr-xr-x 4 root root 4096 июн 29 22:40 ..
-rw-r--r-- 1 root root  993 июн 29 22:45 developer.crt
-rw-r--r-- 1 root root  891 июн 29 22:45 developer.csr
-rw------- 1 root root 1675 июн 29 22:41 developer.key
````

* Создаю роль

````
vim role.yml
````

````
apiVersion: rbac.authorization.k8s.io/v1
kind: Role
metadata:
  namespace: app-namespace
  name: developer-role
rules:
- apiGroups: [""]
  resources: ["pods", "pods/log"]
  verbs: ["get", "list"]
````

````
kubectl apply -f role.yml
role.rbac.authorization.k8s.io/developer-role created
````

* Создаю rolebinding

````
vim rolebinding.yml
````

````
apiVersion: rbac.authorization.k8s.io/v1
kind: RoleBinding
metadata:
  name: developer-rolebinding
  namespace: app-namespace
subjects:
- kind: User
  name: developer
  apiGroup: rbac.authorization.k8s.io
roleRef:
  kind: Role
  name: developer-role
  apiGroup: rbac.authorization.k8s.io
````

````
kubectl apply -f rolebinding.yml
rolebinding.rbac.authorization.k8s.io/developer-rolebinding created
````

* пользователь может просматривать логи подов и их конфигурацию (kubectl logs pod <pod_id>, kubectl describe pod <pod_id>)

* * смена контекста
````
kubectl config use-context app-namespace-developer
Switched to context "app-namespace-developer".
````    

* * проверка
````    
root@k8s-VirtualBox:/home/developer# kubectl get pods
NAME                          READY   STATUS    RESTARTS   AGE
hello-node-6b89d599b9-bwd2h   1/1     Running   0          34m
````

````
root@k8s-VirtualBox:/home/developer# kubectl describe pods | head -n 3
Name:         hello-node-6b89d599b9-bwd2h
Namespace:    app-namespace
Priority:     0
````

````
root@k8s-VirtualBox:/home/developer# kubectl logs pods/hello-node-6b89d599b9-bwd2h
root@k8s-VirtualBox:/home/developer#
````

````
root@k8s-VirtualBox:/home/developer# kubectl delete pod hello-node-6b89d599b9-bwd2h
Error from server (Forbidden): pods "hello-node-6b89d599b9-bwd2h" is forbidden: User "developer" cannot delete resource "pods" in API group "" in the namespace "app-namespace"
````

````
root@k8s-VirtualBox:/home/developer# kubectl create deployment hello-node --image=k8s.gcr.io/echoserver:1.4
error: failed to create deployment: deployments.apps is forbidden: User "developer" cannot create resource "deployments" in API group "apps" in the namespace "app-namespace"
````

````
root@k8s-VirtualBox:/home/developer# kubectl get nodes
Error from server (Forbidden): nodes is forbidden: User "developer" cannot list resource "nodes" in API group "" at the cluster scope
````

# Задание 3: Изменение количества реплик

Поработав с приложением, вы получили запрос на увеличение количества реплик приложения для нагрузки. Необходимо изменить запущенный deployment, увеличив количество реплик до 5. Посмотрите статус запущенных подов после увеличения реплик.

Требования:

в deployment из задания 1 изменено количество реплик на 5

````
kubectl scale deployment hello-node --replicas=5
deployment.apps/hello-node scaled
````

проверить что все поды перешли в статус running (kubectl get pods)

````
kubectl get pods
NAME                          READY   STATUS    RESTARTS   AGE
hello-node-6b89d599b9-7xj8w   1/1     Running   0          54m
hello-node-6b89d599b9-8gbwj   1/1     Running   0          54m
hello-node-6b89d599b9-gf5hf   1/1     Running   0          19s
hello-node-6b89d599b9-kq9lw   1/1     Running   0          19s
hello-node-6b89d599b9-szmk8   1/1     Running   0          10m
````

