# TanX.fi Infrastructure Engineer-Task

## Language and Tools Used:

- Golang
- Docker
- Docker Compose

## Initial Setup :

- First clone the repository .
```
git clone https://github.com/Kamalesh-Seervi/tanxfi-Infrastructure-Engineer-Task
```
- Now go to `tanxfi-Infrastructure-Engineer-Task` folder .
```
cd tanxfi-Infrastructure-Engineer-Task
```

## Docker Setup :

- In `tanxfi-Infrastructure-Engineer-Task` folder there will be a folder called `go`. In that there are two Dockerfiles .
```
cd go
```
- The  First `Dockerfile` is for the task, and another one `Dockerfile.test` for the test.
- To run these two file type the below command to use the application.

**For task service :**

> Build the Docker Image:
```
docker build -t golang-task .
```

> Run the Docker Container:
```
docker run -it golang-task
```

**For test service :**

> Build the Docker Image:
```
docker build -t golang-test -f Dockerfile.test .
```

> Run the Docker Container:
```
docker run -it golang-test
```

## Docker Compose Setup :

- Go to `tanxfi-Infrastructure-Engineer-Task` folder .
- There will a file called `docker-compose.yml`.
- To run the compose file and check the `task` and `test` file type the below commands.

```
docker compose up
```
- This will automatically build and run the containers

> Task service:
```
docker compose up app
```
> Test Service:
```
docker compose up test
```

## Clean up :

- These commands are for cleaning up of containers.
```
docker container rm <container_id or name>
```

### Output Images :

> For Task Service.

![image](https://github.com/Kamalesh-Seervi/tanxfi-Infrastructure-Engineer-Task/assets/107933310/fffc845c-f21d-4059-8ff3-b75f53fae942)

> For Test Service.
> 
![image](https://github.com/Kamalesh-Seervi/tanxfi-Infrastructure-Engineer-Task/assets/107933310/e565791c-367b-4042-ac25-93c7c975c885)

### Test Cases:

- To test the cases locally with docker type below command .

```
cd go
```
```
go test ./tests
```

- To test cases for particular functions.

```
go test -run TestCSVParser
```
```
go test -run TestRevenueCalculation
```

![image](https://github.com/Kamalesh-Seervi/tanxfi-Infrastructure-Engineer-Task/assets/107933310/c107639c-7c1a-4db7-8968-f51d8f968709)


## Note 
- If there is an error in Docker compose build try the below command.
```
docker compose build --no-cache
```


