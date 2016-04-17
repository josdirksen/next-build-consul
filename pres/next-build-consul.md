## Service Discovery in a [Microservice]() Architecture using [Consul]()



## Who Am I

* Currently doing Devops, Scala stuff
    * At Equeris, lean startup within Equens
    * Docker, Consul, Scala, Cloud and other buzzwords.
* Email me at: [jos.dirksen@gmail.com](mailto:jos.dirksen@gmail.com)
* I write at: http://www.smartjava.org
* Twitter: [@josdirksen](http://www.twitter.com/josdirksen)



## Follow along

* Sources, presentation etc:
    * https://github.com/josdirksen/next-build-consul
* Demo heavy session

![](../images/db-demo.png)



## Microservices?

![alt text](../images/microservices-hipsters.png )



### What are microservices?

* Small, fine-grained easy to replace components.
* Organized around capabilities.
* Different languages and backends (whatever fits best).
* Fault tolerant, resiliant, automated deployements.

> "Small Autonomous services that work together", Sam Newman



### From three tier to Microservices

![](../images/3-to-ms.png)



### Running Microservices is hard

* Where is my other service or database?
* Am I healthy, is the other one healthy?
* Where do I store configuration?
* How do I handle redundancy and failover?
* ...

>*["Distributed systems are hard"]()*, says everyone



# Service Discovery

![](../images/shouldnt_be_hard.png)



### Basic approach

Hardcoded IP Address or [DNS]() Lookup

![](../images/sd-1.png)

* [DNS]() Lookup is nice!
* Requires managing names (config files), DNS Server
* How to handle failover? 



### Now with failover

Point [DNS]() to a loadbalancer

![](../images/sd-2.png)

* Works nicely with [DNS]()!
* How to check health and register services?
* Programmatic access to LB?



### What would be nice

![](../images/sd-3.png)
* Does [Lookups](): Lightweight ([DNS](), REST) support failover
* Has flexible [Health checking]() and manages [Configuration]()



# Consul
> "Consul [..] provides an opinionated framework for service discovery and eliminates the guess-work and development effort. Clients simply register services and then perform discovery using a DNS or HTTP interface. Other systems require a home-rolled solution." 
> *-* [consul.io]()



## Main Features

* Service discovery through REST and DNS
* Simple registration using REST API
* Distributed KV store for configuration
* Provides extensive health checking

*All in one package*



## Good to know
 
* Multi DC-ready
* API for distributed locks
* Easy HA Setup
* Event system

> Consul = Zookeeper + Nagios + DNSMasq + Scriptings + ...



### Consul Architecture

![](../images/consul-arch.png)



### Service Registration Flow
1. `Service` calls `Consul Agent` with registration message: http://agent_host/v1/agent/service/register.
2. `Agent` communicates registration with `Consul Server`
3. `Agent` checks health of `Service`.
4. If check succeeds mark as `Healthy`, if not mark as `Unhealthy`, communicate results with `Server`
5. When a lookup for `Service` occurs, only return `Healthy` services



### Sample: registration message

```
{
  "Name": "service1",
  "address": "10.0.0.12",
  "port": 8080,
  "Check": {
     "http": "http://10.0.0.12:8080/health",
     "interval": "5s"
  }
}
```
* Send when a new service starts up
* Check types: `script`, `http`, `tcp`, `TTL`, `Docker`



## Sample: DNS Lookup

```
$dig @nb-consul.local backend-service.service.consul       

; <<>> DiG 9.8.3-P1 <<>> @nb-consul.local backend-service.service.consul
; (1 server found)
;; global options: +cmd
;; Got answer:
;; ->>HEADER<<- opcode: QUERY, status: NOERROR, id: 27716
;; flags: qr aa rd ra; QUERY: 1, ANSWER: 3, AUTHORITY: 0, ADDITIONAL: 0

;; QUESTION SECTION:
;backend-service.service.consul.        IN      A

;; ANSWER SECTION:
backend-service.service.consul. 0 IN    A       10.0.9.3
backend-service.service.consul. 0 IN    A       10.0.9.2
backend-service.service.consul. 0 IN    A       10.0.9.4
```
* Consul provides a DNS Server
* Works great with Docker ([teaser](): will show in Demo)



## Sample: REST Lookup

```
$ curl -s http://192.168.99.106:8500/v1/catalog/service/backend-service 
[{
    "Node": "cf2f293e423c",
    "Address": "192.168.99.111",
    "ServiceID": "backend-service",
    "ServiceName": "backend-service",
    "ServiceAddress": "10.0.9.2",
    "ServicePort": 8080
  },{
    "Node": "072b4ea1abc1",
    "Address": "192.168.99.112",
    "ServiceID": "backend-service",
    "ServiceName": "backend-service",
    "ServiceAddress": "10.0.9.3",
    "ServicePort": 8080
   }]
```



#DEMO

![](../images/demo1.png)



### Closer look: DNS Lookup

In code:
```
resp, err := http.Get("http://backend-service:8081/")
		if err != nil {
			fmt.Println(err)
		} else {
			defer resp.Body.Close()
			body, _ := ioutil.ReadAll(resp.Body)
			w.Header().Set("Content-Type",resp.Header.Get("Content-Type"))
			w.Write(body)
		}
```
For docker:
```
    dns: 192.168.99.106
    dns_search: service.consul
```



# Consul Ecosystem



## Consul template

* Render template based on Consul state
* Setup reverse proxy: `Nginx`, `Apache`, `haproxy`

```
global
    daemon
    maxconn {{key "service/haproxy/maxconn"}}

defaults
    mode {{key "service/haproxy/mode"}}{{range ls "service/haproxy/timeouts"}}
    timeout {{.Key}} {{.Value}}{{end}}

listen http-in
    bind *:8000{{range service "release.web"}}
    server {{.Node}} {{.Address}}:{{.Port}}{{end}}
```



## Demo

![](../images/demo2.png)



## Envconsul

>"III. Config
>Store config in the environment", http://12factor.net/

* More settings `>` More complexity

```
$ envconsul \
  -consul demo.consul.io \
  -prefix redis/config \
  redis-server [opts...]
```  
* Vault: help in managing `secrets`



## DEMO ENVCONSUL

![](../images/demo3.png)



### Prometheus & Grafana

![](../images/grafana.png)



## More information

* links:
    * https://www.consul.io/
    * https://github.com/hashicorp/consul-template
    * https://github.com/hashicorp/envconsul
* This presentation, sources and docker stuff:
    * https://github.com/josdirksen/next-build-consul



# Thank You!
