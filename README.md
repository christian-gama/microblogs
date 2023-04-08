# MicroBlogs
---

A simple microservices project that uses an in-memory database and a custom event bus for simplicity and learning purposes.

## Project Structure

The project is organized into the following directories:

`./posts`: Contains the code for the Posts service.

`./comments`: Contains the code for the Comments service.

`./shared`: Contains shared types and utility functions used across different services.

`./query`: Contains the code for the Query service.

`./moderation`: Contains the code for the CommentModerator service.

`./eventbus`: Contains the code for the EventBus service.

## Dependencies

The project uses Go 1.20 and the GoFiber framework (github.com/gofiber/fiber/v2) for HTTP server implementation.

## Build and Run

A Makefile is provided to build and run each service. The available targets are:

`post`: Builds and runs the Posts service.

`comment`: Builds and runs the Comments service.

`event`: Builds and runs the EventBus service.

`moderation`: Builds and runs the CommentModerator service.

`query`: Builds and runs the Query service.

To build and run a specific service, use the `make` command followed by the target name. For example, to run the Posts service:

```bash
make post
```

To run the other services, replace `post` with the appropriate target name.