<div align="center">
    <h1>💾 Wolke API</h1>
    Wolke API is the API behind Wolke image storage and processing aswell as user management
    <br>
    <br>
</div>

# Deploying

To deploy Wolke Bot you'll need podman with rootless setup

To start the database server use:

```
# This is where database data will be stored
$ mkdir postgresql_data
# Run a postgres container in detached mode
$ podman run --name wolke_postgres -d -e POSTGRES_DB=wolke -e POSTGRES_PASSWORD=postgres -p 5432:5432 -v $(pwd)/postgresql_data:/var/lib/postgresql/data:z docker.io/postgres:13-alpine
```

To build Wolke API main image use:

```
# Build Wolke API with tag latest
$ podman build -t wolke_api:latest .
```

To run Wolke API use:

```
# Run Wolke API in detached mode
$ podman run --name wolke_api -d wolke_api:latest
```

# License
MIT