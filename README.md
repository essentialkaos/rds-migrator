<p align="center"><a href="#readme"><img src=".github/images/card.svg"/></a></p>

<p align="center">
  <a href="https://kaos.sh/w/rds-migrator/ci"><img src="https://kaos.sh/w/rds-migrator/ci.svg" alt="GitHub Actions CI Status" /></a>
  <a href="https://kaos.sh/w/rds-migrator/codeql"><img src="https://kaos.sh/w/rds-migrator/codeql.svg" alt="GitHub Actions CodeQL Status" /></a>
  <a href="#license"><img src=".github/images/license.svg"/></a>
</p>

<p align="center"><a href="#installation">Installation</a> • <a href="#usage">Usage</a> • <a href="#ci-status">CI Status</a> • <a href="#contributing">Contributing</a> • <a href="#license">License</a></p>

<br/>

`rds-migrator` is a tool for migrating _Redis-Split_ metadata to [RDS](https://kaos.sh/rds) format.

### Installation

#### From source

To build the `rds-migrator` from scratch, make sure you have a working Go 1.21+ workspace (_[instructions](https://go.dev/doc/install)_), then:

```
go install github.com/essentialkaos/rds-migrator
```

#### Prebuilt binaries

You can download prebuilt binaries for Linux and macOS from [EK Apps Repository](https://apps.kaos.st/rds-migrator/latest):

```bash
bash <(curl -fsSL https://apps.kaos.st/get) rds-migrator
```

### Usage

<p align="center"><img src=".github/images/usage.svg"/></p>

### CI Status

| Branch | Status |
|--------|----------|
| `master` | [![CI](https://kaos.sh/w/rds-migrator/ci.svg?branch=master)](https://kaos.sh/w/rds-migrator/ci?query=branch:master) |
| `develop` | [![CI](https://kaos.sh/w/rds-migrator/ci.svg?branch=develop)](https://kaos.sh/w/rds-migrator/ci?query=branch:develop) |

### Contributing

Before contributing to this project please read our [Contributing Guidelines](https://github.com/essentialkaos/contributing-guidelines#contributing-guidelines).

### License

[Apache License, Version 2.0](http://www.apache.org/licenses/LICENSE-2.0)

<p align="center"><a href="https://essentialkaos.com"><img src="https://gh.kaos.st/ekgh.svg"/></a></p>
