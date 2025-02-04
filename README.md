# Go Gin Boilerplate

## Setup Project

- Clone project

```bash
  git clone https://github.com/adimurianto/go-gin-boilerplate.git
  cd go-gin-boilerplate
```

- Get all dependency

```bash
  go get .
```

- Create file **.env** base on file **.env.example** and adjust the contents of the variable

-  Init for generate Swagger API documentation

```bash
  swag init
```

-  Init for live reload

```bash
  air init
```

- Access url http://127.0.0.1:5000/docs/index.html

```bash
  air // if using live reload

  go run main.go // test locally on your system
  go build main.go // build locally on your system then make sure run /main.exe or /main
```

## Folder Structure
Here are the folders to consider when adding a new endpoint

```
├── controllers/
│   ├── user_controller.go
│   └── ...
├── models/
│   ├── user_model.go
│   └── ...
├── routes/
|   ├── groups/
│       ├── user.go
|       └── ...
|   └── ...
```
