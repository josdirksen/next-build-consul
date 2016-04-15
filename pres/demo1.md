# Cleanup beforehand

```
. dm-env nb1 --swarm
docker-compose -f docker-compose-agents.yml down
docker-compose -f docker-compose-backend.yml down
docker-compose -f docker-compose-frontend.yml down
```

* Cleanup registry on: http://nb-consul.local:8500


# Demo 1

This demo shows a basic setup of docker / swarm / compose together with Consul.

1. Show and explain some general stuff about containers and show the consul one.

```
$ dm ls

// one Container has consul running
$ . dm-env consul
$ docker ps
```

- Show UI: http://nb-consul.local:8500, and see that there is nothing in there.

2. Connect to the swarm and start the consul agents

```
// create consul agents
docker-compose -f docker-compose-agents.yml up -d

docker ps
 or 
docker ps --format '{{ .ID }}\t{{ .Image }}\t{{ .Command }}\t{{ .Names}}'

// show cluster status
docker exec -ti consul_agent_1 consul members
```

- Optionally: Show docker config in `docker-compose-agents.yml`
- Show UI: http://nb-consul.local:8500, and see that the agents have registered themselves.

3. Now start the backend services

```
docker-compose -f docker-compose-backend.yml up -d
docker ps --format '{{ .ID }}\t{{ .Image }}\t{{ .Command }}\t{{ .Names}}'
```

- Show UI: http://nb-consul.local:8500, and see the services are there
- Show `entrypoint.sh`

```
// stop an instance and show ui: http://nb-consul.local:8500
docker stop nb3/Backend3
docker start nb3/Backend3
```

4. Show how to access them

```
// what is your ip address
dig @nb-consul.local consul.service.consul +short
// what are the ip addresses of the backend services
dig @nb-consul.local backend-service.service.consul +short
// or just query one of the agents directly
curl -s http://nb1.local:8500/v1/catalog/service/backend-service | jq
```

5. The frontend services

Start the frontend services:
```
docker-compose -f docker-compose-backend.yml up -d
docker ps --format '{{ .ID }}\t{{ .Image }}\t{{ .Command }}\t{{ .Names}}'
```

They communicate with the backend services through golang.
- http://nb1.local:8090/
- http://nb2.local:8090/
- http://nb3.local:8090/
- Show how this works

```
// bring down one of the backend services
docker stop Backend1
// show in consul: http://nb-consul.local:8500
```
- Show that http://nb1.local:8090/ switches to a different host.
```
// bring it back up
docker start Backend1
// drop the other two
docker stop Backend2
docker stop Backend3
// and see the result: http://nb1.local:8090/ 
```
