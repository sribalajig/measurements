# Overview

This is a Go implementation of a RESTful api for performing CRUD operations on user body measurements.

# How to run

### System requirements

* docker 
* docker-compose

### Run the container

* Navigate to /hpi/measurement
* Execute the following command -
	```
	docker-compose up --build
	```
### Graceful shutdown
* Execute the following command -
	```
	docker-compose down
	```
# Examples

### Create/Update body measurements for a user

```
HTTP POST

http://localhost:9090/users/124/bodyMeasurements

Body -

{
	"measurements" :
		[ 
			{
				"id" : 3,
				"type" : "temperature",
				"value" : 37,
				"units" : "celcius"
			},
			{
				"id" : 4,
				"type" : "weight",
				"value" : 70,
				"units" : "kg"
			}
		]	
}

```

### Get body measurements given user id

```
HTTP GET

http://localhost:9090/users/124/bodyMeasurements
```

### Get *specific* body measurements given user id

```
HTTP GET


http://localhost:9090/users/124/bodyMeasurements?id=4&id=1
```

# Running tests

* Make sure the hpi folder is under /src/github.com
* Run 

```
go test ./...
```

# Going through the code

The entry point is /hpi/measurement/pkg/cmd/api/main.go

