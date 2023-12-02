# GONION

Generate Go project that implement Hexagonal Architecture (Onion)

## How to use

install CLI
```shell
go install github.com/ac-kurniawan/gonion
```

start generate new project
```shell
gonion generate --repository github.com/ac-kurniawan/gonion
```

shorthand
```shell
gonion g -r github.com/ac-kurniawan/gonion
```

## Result
```
/{{moduleName}}
    /core
        /model
        service.go
        repository.go
    /adaptor
    /interface
    mod.go
    Dockerfile
    .gitignore
    properties.yml
```