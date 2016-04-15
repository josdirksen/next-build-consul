# Env-consul

# For instance we can see what docker has stored
 ./envconsul -consul=nb-consul.local:8500 -prefix docker -once env
 
# Or easily pass in our own config
./envconsul -consul=nb-consul.local:8500 -prefix nb/config -once env