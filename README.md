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
```

Testing on local
* API
	* http://localhost:8000/accounts/1
	* http://localhost:8000/accounts/2
	* http://localhost:8000/accounts/3

Delete all 
```
$docker-compose down
```