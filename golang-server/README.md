# Dockerizing Go To Do Server Middleware

## Prerequisites
* Tutorial assumes you have completed this tutorial: https://levelup.gitconnected.com/build-a-todo-app-in-golang-mongodb-and-react-e1357b4690a6

## Notes:  

#### We have to talk about Go modules if want to Dockerize the Go middleware. Go modules are essentially to Go what package.json files are to Node apps. Modules track external dependencies to your Go code. Creating a module outputs a go.mod file that your Dockerfile will read for dependencies to create the image.

## Steps:

### 1. Create a private Github repo for your Go module. Strangely enough, Go modules have to be tracked with VCS, so that's what we'll do first.

```
$ cd go-todo

$ cd golang-server 

$ git init

$ git add .

$ git commit -m "initial commit"
```

// Create private repo in Github to host your module code 

```
$ git remote add origin https://github.com/<your_github_username>/<your-go-todo-server-private-repo>.git

$ git push -u origin master
```

// Provide Github credentials when prompted on command line 

### 2. Initialize Go module, which will output two new files to your folder: go.mod and go.sum

```
$ go mod init github.com/<your_github_username>/<your-go-todo-server-private-repo>

$ go build ./...
```

### 3. Build Docker image using the provided Dockerfile and run locally 

```
$ docker build -t yourname/todoserver .
```

// Get this error? “build ../middleware: cannot find module for path ../middleware” No worries, this is because your code doesn’t understand the relative path anymore since you made your project a module on Github.

// Simply change ''../middleware" and all other relative path import occurrences to “github.com/<your_github_username>/<your-go-todo-server-private-repo>/middleware“

```
$ go get -u 

$ docker build -t yourname/todoserver .

$ docker images 

$ docker run -d -p 8080:8080 yourname/todoserver:latest

$ docker ps
```

#### 4. Test API in Postman and stop running container instance of image when finished

GET 127.0.0.1:8080/api/task (will return all tasks in collection)

```
$ docker stop <running CONTAINER ID for todoserver image>
```

#### 5. Finally, don't forget to push those changes you made to fix the relative import issue to your Github repo

```
$ git add --all

$ git commit -m "updated imports to use module path instead"

$ git push -u origin master
```

## Helpful Links:

Package Management With Go Modules The Pragmatic Guide:
https://medium.com/@adiach3nko/package-management-with-go-modules-the-pragmatic-guide-c831b4eaaf31

Everything you need to know about Packages in Go:
https://medium.com/rungo/everything-you-need-to-know-about-packages-in-go-b8bac62b74cc

Importing other packages in the same module:
https://stackoverflow.com/questions/55442878/organize-local-code-in-packages-using-go-modules

Anatomy of Modules in Go:
https://medium.com/rungo/anatomy-of-modules-in-go-c8274d215c16

To create a go.mod for an existing project:
https://github.com/golang/go/wiki/Modules#gomod

Building Docker Containers for Go Applications:
https://www.callicoder.com/docker-golang-image-container-example/

















