## Service Discovery in a [Microservice]() Architecture using [Consul]()



## Who Am I

* Currently doing Devops, Scala stuff
    * At Equeris, lean startup within Equens
    * Docker, Consul, Scala, Cloud and other buzzwords.
* Email me at: [jos.dirksen@gmail.com](mailto:jos.dirksen@gmail.com)
* I write at: http://www.smartjava.org
* Twitter: [@josdirksen](http://www.twitter.com/josdirksen)



## Microservices?

![alt text](../images/microservices-hipsters.png )



### What are microservices?

* Small, fine-grained easy to replace components.
* Organized around capabilities.
* Different languages and backends (whatever fits best).
* Fault tolerant, resiliant, automated deployements.

> "Small Autonomous services that work together", Sam Newman



### We moved from this...

![alt text](../images/2-layer.png "Logo Title Text 1")



### To this...

![alt text](../images/n-layer.png "Logo Title Text 1")



### Or even this...

![alt text](../images/n-layer-container.png "Logo Title Text 1")



### Running Microservices is hard

* Where is my other service?
* Am I healthy, is the other one healthy?
* Where do I store configuration
* ...

>*["Distributed systems are hard"]()*, says everyone



### What we want...

![alt text](../images/registry-discovery.png "Logo Title Text 1")
* Lightweight and non intrusive
* Simple, so no UUDI or WSRR


### You can use 