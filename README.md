# go-clean-arch-app
Clean Architecture Golang App


# Get Started
... 1) make your database
... 2)

go run main.go

# Basic usage examples

http://localhost:8000/users
http://localhost:8000/user?id=1
http://localhost:8000/posts

# StoryTelling Docs Style

1) ... define two infrastructure objects: Logger e SQLHandler
infrastructure.looger.go ---> interfaces.logger.go
infrastructure.sqlHandler.go ---> interfaces.sqlHandler.go

2) ... in main.go ... get two infrastructure objects
infrastructure.NewLogger() --> (interfaces.Logger) <-- Logger struct
infrastructure.NewSQLHandler() --> (interfaces.SQLHandler) <-- SQLHandler struct

3) ... use Load function (defined in env.go) to set Logger with environment vars

4) ... In router.go, we call infrastructure.Dispatch(logger, sqlHandler)
   a) to create a controller for each entity in my app domain
   b) to define routes
   c) every route receive a controller required to operate
   d) start app in a given port, see env.go and *.env

    UserController --- has a --- usecases.UserInteractor
    UserController --- has a --- interface.Logger

5) ... programming to interface:
usecases.UserInteractor - has a - usecases.UserInteractor
usecases.UserInteractor ---> ( interface.UserRepository ) <--- usecases.UserRepository
(tip) you can change database vendor by rewriting user_repository.go (usecases)

usecases.PostInteractor - has a - usecases.PostInteractor
usecases.PostInteractor ---> ( interface.UserRepository ) <--- usecases.PostRepository
(tip) you can change database vendor by rewriting post_repository.go (usecases)

# go-clean-architecture-web-application-boilerplate
A web application boilerplate built with go and clean architecture.
Most of this application built by standard libray.

![image](https://user-images.githubusercontent.com/13291041/102681893-84326980-4208-11eb-8f84-2959e03b89d8.png)


# Architecture
```
app
├── database
│   ├── postgres-script
│   │   └── scritps.sql
├── domain
│   ├── post.go
│   └── user.go
├── infrastructure
│   ├── env.go
│   ├── logger.go
│   ├── router.go
│   └── sqlhandler.go
├── interfaces
│   ├── logger.go
│   ├── post_repository.go
│   ├── sqlhandler.go
│   └── user_repository.go
├── log
│   ├── access.log
│   └── error.log
├── main.go
└── usecases
    ├── post_controller.go
    ├── post_interactor.go
    ├── post_repository.go
    ├── user_controller.go
    ├── user_interactor.go
    └── user_repository.go

7 directories, 22 files
```

| Layer                | Directory      |
|----------------------|----------------|
| Frameworks & Drivers | infrastructure |
| Interface            | interfaces     |
| Usecases             | usecases       |
| Entities             | domain         |

# API

| ENDPOINT | HTTP Method    | Parameters    |
|----------|----------------|---------------|
| /users   | GET            |               |
| /user    | GET            | ?id=[int]     |
| /posts   | GET            |               |
| /post    | POST           |               |
| /post    | DELETE         | ?id=[int]     |

# Controller method naming rule

| Controller Method | HTTP Method | Description                                |
|-------------------|-------------|--------------------------------------------|
| Index             | GET         | Display a listing of the resource          |
| Store             | POST        | Store a newly created resource in storage  |
| Show              | GET         | Display the specified resource             |
| Update            | PUT/PATCH   | Update the specified resource in storage   |
| Destroy           | DELETE      | Remove the specified resource from storage |

# Repository method naming rule

| Repository Method | Description                                          |
|-------------------|------------------------------------------------------|
| FindByXX          | Returns the entity identified by the given XX        |
| FindAll           | Returns all entities                                 |
| Save              | Saves the given entity                               |
| SaveByXX          | Saves the given entity identified by the given XX    |
| DeleteByXX        | Deletes the entity identified by the given XX        |
| Count             | Returns the number of entities                       |
| ExistsBy          | Indicates whether an entity with the given ID exists |

cf. [Spring Data JPA - Reference Documentation](https://docs.spring.io/spring-data/data-jpa/docs/current/reference/html/#repositories.core-concepts)

# Tests
I have no tests because of my laziness, but I will prepare tests in future commits.

# References
- [github - manuelkiessling/go-cleanarchitecture](https://github.com/manuelkiessling/go-cleanarchitecture)
- [github - rymccue/golang-standard-lib-rest-api](https://github.com/rymccue/golang-standard-lib-rest-api)
- [github - hirotakan/go-cleanarchitecture-sample](https://github.com/hirotakan/go-cleanarchitecture-sample)
- [Recruit Technologies - Go言語とDependency Injection](https://recruit-tech.co.jp/blog/2017/12/11/go_dependency_injection/)
- [Clean ArchitectureでAPI Serverを構築してみる](https://qiita.com/hirotakan/items/698c1f5773a3cca6193e)
- [github - ponzu-cms/ponzu](https://github.com/ponzu-cms/ponzu)
- [クリーンアーキテクチャの書籍を読んだのでAPIサーバを実装してみた](https://qiita.com/yoshinori_hisakawa/items/f934178d4bd476c8da32)
- [Go × Clean Architectureのサンプル実装](http://nakawatch.hatenablog.com/entry/2018/07/11/181453)
- [Uncle Bob – Payroll Case Study (A full implementation)](http://cleancodejava.com/uncle-bob-payroll-case-study-full-implementation/)
