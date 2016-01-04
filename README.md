# Go Test Applications

A collection of applications used for testing the Go buildpack

## Applications

| Name | Description
| ---- | -----------
| `simple-go-application` | A simple Go application
| `cf-go-test-app` | Cloned from https://github.com/combor/cf-test-app

### Output Content

All applications support the following REST operations:

| URI | Description
| --- | -----------
| `GET /` | The health of the application

## Building

Before building the project, the following tools must be installed:

* [Go](https://golang.org/dl/)

To build the project, navigate to the app directory and run the following:

```
$ go build app.go
$ ./app
```

## Deploying to Cloud Foundry

Each test application contains a `manifest.yml` file which allows the built application to be deployed to Cloud Foundry by simply issuing:

```
cf push
```
