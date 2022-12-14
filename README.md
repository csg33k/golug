# What is This?

The purpose of this repo is to show case a really basic REST Crud service.

# Why are you using X / doing Y / Whyyyyyyy ?

Cause I want to.  They're all opinionated choices and not the best or worst but should work for the use case I'm presenting. 

License: I like MIT license, feel free to use anything you like. 

Logging:

[Logrus](https://github.com/sirupsen/logrus) I like it since it's simple to setup and requires basically no configuration.

Database:

[PostgresSQL](https://www.postgresql.org/) Same, I'm more familiar with it than MySQL. It's more SQL standard compliant and more feature rich.  Also, just simply cause I want to. 

[Chi](https://github.com/go-chi/chi) GoLang router.  I've been meaning to look at this for a while so this is a good excuse.  Therefore, I'm embracing my inner Chi.

[Viper](https://github.com/spf13/viper) At this point it feels silly not to use this. It supports, YAML, JSON, ini, hcl, toml, javaproperties, ENV overrides.

More DataBase Opinions: 

[DBMate](https://github.com/amacneil/dbmate) Schema management, requires more writing of raw sql, but I'm a bit fan of Schema first design, so there we go.

[SQLC](https://github.com/kyleconroy/sqlc) ORM that generates data models.

[SQLX](https://github.com/jmoiron/sqlx) An extension that allows for data-mapping queries into structs

## Install Packages

```sh
##Needed for code generation
go install github.com/amacneil/dbmate.git@latest
go install github.com/kyleconroy/sqlc/cmd/sqlc@latest
```

See [quickguide](docs/00_quick_guide.md) for more.
