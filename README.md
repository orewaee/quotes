## TL;DR

The project was originally created as a kind of “Hello World” for quick deployment to a random node from docker and subsequent tests.
That is, it is a small production ready REST API built in a Docker image. Both Docker Compose and Swarm mode are available for deployment.

You can also say that this project is a tutorial project, as I tested Swarm mode for the first time.


## How do I add quotes?

Quotes are stored in the `/quotes` directory as json files. It is enough just to add a file with a new quote.


## Endpoints

It's trivial. There are only two routes:

- `/random` to get a random quote
- `/ping` to check if the api is ready to work
