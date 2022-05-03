# Go Generate Server

WARNING : This project is still unstable, if interested please contribute to develop

Generate server code with simple clean architecture using go language

Using simple framework, just only `gin-gonic` and `GORM`

Support driver : `mysql`, `postgres`, `sqlite`

## installation
> still support for mac

- running Makefile in this root directory

```bash
make build
```

-  add path to bashrc / zshrc
```
export PATH=~/go-gen-server:$PATH
```

-  after than, make the program executable

```
chmod +x ~/go-gen-server/go-gen
```

## List command

### init

you must init service first using this command line

```bash
go-gen init [SERVICE-NAME]
```

example : `go-gen init .` or `go-gen init my-example`

> alias : `go-gen i`

this service will be created list of file like this :

```
.
|-- config
    |-- config.go
|-- utils
    |-- utils.go
|-- main.go
|-- go.mod
|-- go.sum
|-- .env
|-- example.env
|-- .gitignore
```


### init middleware

Command for initialize middleware

```bash
go-gen init-middleware
```

> alias : `go-gen init-mw`

### go-gen create [ENTITY-NAME] [PROPERTIES]

initialize service with simple layer clean architecture (repository, entities / model, service, controller)


```bash
go-gen create [NAME-ENTITY] [PROPERTIES]
```

example : `go-gen create user name:string,username:string,password:string,role:string`

> alias : `go-gen c` or `go-gen cr`

### go-gen remove

Remove initialize service based on entity name

```bash
go-gen remove [NAME-ENTITY]
```

example : `go-gen remove user`

> alias : `go-gen rm`
