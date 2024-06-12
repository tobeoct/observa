# Observa

Welcome to the Observa Solution README! This document provides information about building, running, and using all applications within this solution.

## Table of Contents

- [Introduction](#introduction)
- [Folder Structure](#folder-structure)
- [Dependencies](#dependencies)
- [Building Applications](#building-the-application)
- [Running Applications](#running-the-application)
- [API Endpoints](#api-endpoints)
- [Configuration](#configuration)
- [Contributing](#contributing)
- [License](#license)

## Introduction

The Observa Application is a GoLang agent that provides observability capabilities cutting across endpoint-monitoring, log analysis and metric generation, alerting etc

## Folder Structure
The solution structure is this
- `projects/<project-name>`: Every application has its own folder For example `endpoint-monitor`
- - Every application follows the following folder structure:

- - - `cmd/`: Contains the main application entry point (`main.go`).
- - - - `app/main.go`: The entry point into all applications
- - - - `config.sh`: Metadata about the project
- - - `internal/`: Contains internal application packages.
- - - - `pkg`: Reusable packages shared used only internally

- - `internal/`: Contains internal solution packages.
- - - `pkg/`: Contains reusable packages that can be imported by other applications.
- `scripts/`: Contains scripts for building, running, and testing the application.
- `deployment/`: Contains deployment-related files such as IIS web config, Dockerfile and Kubernetes configuration.
- `.env`: Environment variables file (if applicable).
- `README.md`: This file.
- `go.mod` and `go.sum`: Go module files for managing dependencies.

Make sure to install these dependencies before building and running the application.

## Building the Application

To build the application, you can use the provided build script located in the `scripts/` directory. Follow these steps:

1. Open your terminal.
2. Navigate to the root directory of the project.
3. Run the build script:
    ```bash
    ./scripts/build.sh <application-folder-name>
    ```

The build script will compile the GoLang application and create an executable file in the `build/` directory.

## Running the Application

Once the application is built, you can run it using the following command:
```bash
./build/<executable-name>.exe
```
The executable name can be different from the folder name for example `endpoint-monitor` builds as `observa.endpoint-monitor`. The details are in the config.sh 



You can also run the application by using the go command

```Go
go run projects/<project-folder-name>/cmd/app/main.go
```
