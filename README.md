## Continuous testing workshop
* Frontend
	* ReactJS
* Backend
	* Golang
* Database
	* PostgreSQL
* Docker
* Continuous Integration with Jenkind
  * Build
  * Testing
  * Deploy

### Step to run with Docker
```
$docker-compose build
$docker-compose up -d db
$docker-compose up -d api
$docker-compose ps
$docker-compose logs --follow
```

Testing on local
* API
  * http://localhost:8000/accounts/1
  * http://localhost:8000/accounts/2
  * http://localhost:8000/accounts/3

* Web
  * http://localhost:7999 
  
Delete all 
```
$docker-compose down
```

### Working with Docker swarm mode

Create Docker Swarm
```
$docker swarm init
$docker node ls
```

See join-token command
```
$docker swarm join-token worker
$docker swarm join-token manager
```

Deploy stack with docker-compose file (v3+)
```
$docker stack deploy --compose-file docker-compose.yml demostack
$docker stack services demostack
```

Scaling service
```
docker service ls
$docker service scale <service name>=<number of replica(s)>
$docker service ls
$docker service ps <service name>
```

Example scale to 5
```
$docker service scale demostack_web=5

demostack_web scaled to 5
overall progress: 5 out of 5 tasks
1/5: running   [==================================================>]
2/5: running   [==================================================>]
3/5: running   [==================================================>]
4/5: running   [==================================================>]
5/5: running   [==================================================>]
verify: Service converged
```

Testing with curl and see result
```
$curl http://localhost:8080/
```

Delete services
```
$docker service rm <service name>
```

Delete stack
```
$docker stack rm demostack
```

Leave node from Swarm
```
$docker swarm leave --force
```
