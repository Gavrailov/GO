## Задача 1

#### Go Web Application (стъпка 1)
Да се напише web application на Go, който:
- На адрес “/hello“ връща  “Hello <стойността на query parameter name>“ 
Пример за request “/hello?name=World!” връща отговор “Hello World!“
- На адрес “/healthz” връща HTTP Status 200
- Да има endpoint “/break”. След първата HTTP Post заявка на този адрес (“/break”) заявка от предходния endpoint (“/healthz”) да започва да връща HTTP Status 500
- Web server-ът да се build-не в docker image

За имплементацията може да се използва стандартната библиотека на Go или някоя от известните routers, web frameworks. Пример за такива са:
 - [gorilla/mux](https://github.com/gorilla/mux)
 - [chi](https://github.com/go-chi/chi)
 - [httprouter](https://github.com/julienschmidt/httprouter)
 - [gin](https://github.com/gin-gonic/gin)

Разбира се използването на дадена библиотека или framework трябва да може да се аргументира. Предоставянето на тестове също е плюс. За тестове може да се използва стандартната библиотека или [ginkgo](https://github.com/onsi/ginkgo).

#### Web Application Deployment (стъпка 2)
Да сe използва вградения Kubernetes в Docker или minikube, в който:
- Да се направи Deployment, който да run-ва Docker image-а
Deployment-a да се expose-не като Service (обърни внимание на selector-a)
- Да се дефинира Liveness Probe, който да сочи към “/healthz"
- Какво ще се случи, когато web application-a получи заявка на "/break" адреса?

Ето и малко полезни ресурси по темата:
 - [Go Documentation](https://go.dev/doc/)
 - [Pod](https://kubernetes.io/docs/concepts/workloads/pods/)
 - [Deployments](https://kubernetes.io/docs/concepts/workloads/controllers/deployment/)
 - [Liveness Probe](https://kubernetes.io/docs/tasks/configure-pod-container/configure-liveness-readiness-startup-probes/)
 - [Services](https://kubernetes.io/docs/concepts/services-networking/service/)
 - [kubectl Cheat Sheet](https://kubernetes.io/docs/reference/kubectl/cheatsheet/)

Hint: Ако има нужда може да се провери и youtube за tutorials на тема kubernetes.

**Успех!**