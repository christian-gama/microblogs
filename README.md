# MicroBlogs
---

This is a port of the [Stephen Grider's MicroBlogs project](https://www.udemy.com/course/microservices-with-node-js-and-react/) written in NodeJS to Go.

A simple microservices project that uses an in-memory database and a custom event bus for simplicity and learning purposes. It runs on a Docker/Kubernetes cluster and uses a React web client.

## Project Structure

The project is organized into the following directories:

`./posts`: Contains the code for the Posts service.

`./comments`: Contains the code for the Comments service.

`./shared`: Contains shared types and utility functions used across different services.

`./query`: Contains the code for the Query service.

`./moderation`: Contains the code for the CommentModerator service.

`./eventbus`: Contains the code for the EventBus service.

`./client`: Contains the code for the Web Client service, which is built in React.

## Dependencies

- [Docker](https://docs.docker.com/get-docker/)
- [Kubernetes](https://kubernetes.io/docs/tasks/tools/install-kubectl/)
- [Skaffold](https://skaffold.dev/docs/install/)
- [Go](https://golang.org/doc/install)
- [React](https://react.dev/learn)

## Build and Run

A Makefile is provided to build and run each service. The available targets are:

`post`: Builds and runs the Posts service.

`comment`: Builds and runs the Comments service.

`event`: Builds and runs the EventBus service.

`moderation`: Builds and runs the CommentModerator service.

`query`: Builds and runs the Query service.

`client`: Builds and runs the Web Client service in React.

To build and run a specific service, use the `make` command followed by the target name. For example, to run the Posts service:

```bash
make dev
```

Alternatively, you can run the `skaffold dev` command to build and run all services. This will also watch for changes and automatically rebuild and redeploy the services as needed.
