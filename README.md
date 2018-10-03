# KGB Super Secret Offender list 😎

## Offender criteria

Offenders are chosen by seeking out their overly enthusiastic view of the McKaig Chevrolet Buick dealership.

Certain key phrases will bring you to the top of the offender list very quickly:

```
love
perfect
care
honest
awesome
quick
efficient
friend
best
smile
helpful
pleasant
superior
```

Each offender is ranked based on their use of these words.

> Note: Keywords are of no particular weight

## Using the app

_Download the CLI_

Go to the releases page to download the latest build for your machine

[Releases](https://github.com/tylerwray/red-scare/releases)

_Run it yourself_

> Requires go [1.10](https://golang.org/dl/)

First get the single dependancy

```bash
go get github.com/PuerkitoBio/goquery
```

List offenders:

```bash
go run main.go
```

Run the test suite:

```bash
go test -cover ./...
```

_In Docker?_

List offenders:

```bash
make run
```

> Note: Running the application in docker _may_ not work if you are on a mac due to certificate issues with the request server

Run the test suite:

```bash
make test
```
