# Boilerplate GO

[![Coverage Status](https://coveralls.io/repos/evermos/payment-reconciliation/badge.svg?branch=master)](https://coveralls.io/r/evermos/payment-reconciliation)
[![GitHub Stars](https://img.shields.io/github/stars/evermos/payment-reconciliation.svg)]()
[![Build Status](https://jenkinsx.evermosa2z.com/buildStatus/icon?job=be-payment-reconciliation%2Fmaster)](https://jenkinsx.evermosa2z.com/job/be-payment-reconciliation/job/master/)

## Description
This repo is boilerplate go with:
 - wire : dependency injection using [google wire](https://github.com/google/wire)
 - infras : infrastructure for connection to db
 - internal: main service package
   - domain
     - model
     - repo
     - service
   - handler
 - external : this package for external service

## System Requirement
- Go 1.20

## Setup and Installation
 - clone this repo
 - open file main.go
 - execute command go generate or copy command `go run github.com/google/wire/cmd/wire` in terminal
   - and then, file wire_gen was generate automation
 - now, you can running service

## Contributing
 - [dzul](https://github.com/dzulfiqarali)
 - [fikgatrh](https://github.com/fikrigatrh)
