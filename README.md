# weblair ag7if

One Paragraph of project description goes here

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing 
purposes. See deployment for notes on how to deploy the project on a live system.

### Prerequisites

  - Go 1.14.2 (or later)
  - Postgres 11.4 (or later)
  - GNU Make 3.81 (or later)
  - [Swaggo CLI](https://github.com/swaggo/swag)
  - [Lair](https://github.com/weblair/lair.git) (Optional, but recommended).
  - [Air](https://github.com/cosmtrek/air) (Necessary only if you want to use `make run`)

### Dev Setup (with Lair)

  1. `git clone git@github.com:weblair/maricopa.git`
  2. `go mod tidy`
  3. `lair db create --seed`
  4. `make run`

### Dev Setup (without Lair)

  1. `git clone git@github.com:weblair/maricopa.git`
  2. `psql -U postgres -c "CREATE DATABASE ag7if_development;"`
  3. `migrate -source file://db/migrations -database postgres://localhost:5432/ag7if_development?sslmode=disable up`
  4. `go mod tidy`
  5. `make run`

## Running the tests

Explain how to run the automated tests for this system

### Break down into end to end tests

Explain what these tests test and why

```
Give an example
```

### And coding style tests

Explain what these tests test and why

```
Give an example
```

## Deployment

Add additional notes about how to deploy this on a live system

## Built With

  - [Gin](https://gin-gonic.com) - API framework
  - [Postgres](https://postgresql.org) - Database
  - [Migrate](https://github.com/golang-migrate/migrate) - Database Migrations
  - [Swaggo](https://github.com/swaggo/swag) - API Documentation

## Contributing

Before committing, be sure to:
  1. Use [Git Flow](https://www.atlassian.com/git/tutorials/comparing-workflows/gitflow-workflow)
  2. Use [conventional commit messages](https://www.conventionalcommits.org/en/v1.0.0-beta.2/) 
     (see below for the tags used in this repo).
  3. [Sign your commits](https://git-scm.com/book/ms/v2/Git-Tools-Signing-Your-Work)
  4. Run `gofmt -w -s .`
  5. Run `swag init`
  6. Use the [Go Reportcard CLI](https://github.com/gojp/goreportcard)

### Commit Message Guidelines

The rules of [conventional commit messages](https://www.conventionalcommits.org/en/v1.0.0-beta.2/) should be observed.
Observe to keep the first line of the commit mesage down to 50 characters and insert hard line-breaks at 72 characters
for the rest of the message body.

When working on `feature` or `hotfix` branches, the rules can be relaxed a bit. PRs should only be opened from your 
`develop` branch, and when wrapping up your `feature` branches, you should squash your commits.

#### Tags
  - fix
  - feat
  - BREAKING CHANGE
  - refactor
  - docs
  - chore

## Versioning

We use [SemVer](http://semver.org/) for versioning. For the versions available, see the 
[tags on this repository](https://github.com/your/project/tags). 

## Authors

  - **Robert Hawk** - *Initial work* - [DerHabicht](https://github.com/DerHabicht)

See also the list of [contributors](https://github.com/weblair/maricopa/contributors) who participated in this project.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details

## Acknowledgments

  - [PurpleBooth](https://github.com/PurpleBooth) for the README template
  - [Commissure](https://github.com/commissure) for the build meta-data Makefile
