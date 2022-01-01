# gottago

This repo contains my personal development environment for golang.


## Preparation

1. Run docker app.
2. Change the current working directory and execute `docker-compose up -d` command on the terminal.


## Dive into the container 🤿

```bash
docker exec -it workspace /bin/sh
```

### If you use VSCode

you can access to the container from VSCode via [Remote-Containers](https://code.visualstudio.com/docs/remote/containers) extension.

The dev container settings are described in `.devcontainer/devcontainer.json` file.


## In closing

That's all. I gotta *Go*. 👋
