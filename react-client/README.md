# Dockerizing Go To Do Client Frontend

## Prerequisites
* Tutorial assumes you have completed this tutorial: https://levelup.gitconnected.com/build-a-todo-app-in-golang-mongodb-and-react-e1357b4690a6
* Tutorial assumes you have completed the previous Go middleware step: https://github.com/ironspur5/todo-app-dockerized-golang-react-mongodb/tree/master/golang-server

## Steps:

### 1. Build Docker image from included Dockerfile

```
$ cd go-todo

$ cd react-client 

$ docker build -t yourname/todoclient .
```

### 2. Run Docker image as a container locally

To test this, you first need to run the Docker image for the previous step (Golang middleware)

```
$ docker run -d -p 8080:8080 patrickguha/todoserver:latest
```

Then run the Docker image for the client 

```
$ docker run -d -p 3000:3000 patrickguha/todoclient:latest
```

### 3. Finally, with both dockerized client and server running, go to localhost:3000 to see the app running!























