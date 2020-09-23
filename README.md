# Virtual-Economy

A virtual economy written in go.

## Tech Stack

- Go
- MongoDB

## Features

1. Artificial Inflation
2. Reserves/exchange creation

## To be implemented

1. Full replication of monetarism/macroeconomic system
2. Exchange
3. APIs

## How to run the app

Run `./run.sh`

## [Learning resources I used to build the project](resources.txt)

## Database schema

### Current

1. All tokens in one document (1 token per object)

### Eventual

1. All tokens in one document (1 token per object)
2. Each user has their own document (wallet)
3. Currency (tokens) can be transferred to their wallet
