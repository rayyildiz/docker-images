Build
---
> CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .
> docker build -t rayyildiz/gin-going-simple . 

Deploy 
---
>docker push rayyildiz/gin-going-simple


Run
---
>docker run -p 8086:8080 -d rayyildiz/gin-going-simple

