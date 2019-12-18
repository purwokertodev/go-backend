# Go Backend

## Ngapak :blue_heart: Golang

[![Build Status](https://travis-ci.org/purwokertodev/go-backend.svg?branch=master)](https://travis-ci.org/purwokertodev/go-backend)
[![codecov](https://codecov.io/gh/purwokertodev/go-backend/branch/master/graph/badge.svg)](https://codecov.io/gh/purwokertodev/go-backend)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://github.com/purwokertodev/go-backend/blob/master/LICENSE)

Scaffolding, Boilerplate with Clean Architecture :facepunch:, Testable :facepunch:, for Build backend in pure Go.
Base on Uncle Bob's awesome post https://8thlight.com/blog/uncle-bob/2012/08/13/the-clean-architecture.html

# This is What you need,,,,

  - Golang version 1.10+
  - Go Coverage Merge; click https://github.com/bookerzzz/gocovmerge

# Test

  - Mocking
    - we use https://github.com/vektra/mockery
      ```shell
        $ go get github.com/vektra/mockery
      ```
    - and mock some `interface`
      ```shell
      $ mockery -name=InterfaceName
      ```

  - Unit test
    ```shell
    $ make test
    ```
# Code Coverage

  ```shell
  $ make cover
  ```

# Format Code
  ```shell
  $ make format
  ```

# Authors
  - Lone Wolf https://github.com/wuriyanto48

##
Purwokerto Dev 2017 :droplet:
