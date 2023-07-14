# Tech Test : Cushon

This is my solution to the scenario of allowing customers
to create investments in an Isa fund.

## Command Options

To use this solution, you can use the following terminal
commands from the repository root:

- `make up` - bring up the `MySQL` database and setup the
schema
- `make down` - bring down the `MySQL` database
- `make client` - get a connection to the `MySQL` database so
you can review the contents of the tables
- `make run` - run the service (listens on port `8080`)
- `make sqlc` - regenerate `SqlC` database query entities

## Implementation

### Schema

For this solution, you can review the database schema I designed
in the following location: `sim/mysql/init.sql`

I have chosen to create associations between companies and isa
funds, allowing you to separate B2B offerings from B2C offerings
by tailoring which funds are offered to a company I created called
`GeneralPublic` - add or remove a fund from this company (in the
`companies_isa_funds` table) to add or remove funds from the public.

### Endpoints

You can interact with my solution with the following endpoints:

- `/funds` - get a list of funds currently available to the general public
- `/investment/customer/:customer/fund/:fund` - get the current level of investment in a fund for a particular customer
- `/invest/customer/:customer/fund/:fund/amount/:amount` - create or update an investment in a fund for a given customer

### Code

I have followed idiomatic Go patterns with the use of an `internal` directory
for core code, and a `cmd` directory containing binaries. I have 
created a `db` struct for communicating with the database which
wrappers auto generated `sqlc` structs for actual querying. The
reason I have added this layer is that I can easily use an interface
for the `db` struct in all the http handlers - which makes unit testing
with mocks much easier.