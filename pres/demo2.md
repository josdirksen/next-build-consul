# Consul-template

1. Run this in only one container
2. Hides the frontend endpoints
3. Hides the backend endpoints

```
. dm-env nb1 --swarm
docker-compose -f docker-compose-haproxy.yml up -d

// restart with configured to use haproxy
docker-compose -f docker-compose-frontend.yml down
docker-compose -f docker-compose-frontend-proxy.yml up -d

// show config file
docker exec -ti nb-haproxy cat /etc/haproxy/haproxy.cfg

// do some stuff in browser: http://nb1.local:1080

// stop some frontend servers and backend servers
docker stop Frontend2
docker stop Backend2

// show config file
docker exec -ti nb-haproxy cat /etc/haproxy/haproxy.cfg

// do some stuff in browser: http://nb1.local:1080
```


