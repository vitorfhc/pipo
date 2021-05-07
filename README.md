# Pipo
Pipo is composed of two main parts: an auth server and a [Hasura](https://hasura.io) server.

## Running the project

You can run it with a `docker-compose up`. All environment variables are already added, but you could change it. It is important to mention that the RSA keys present at `heimdall/jwt` are not for production use since they were published in this repository. Anyway, feel free to use them for testing and development.

## Seeding

To seed the database you may use `docker-compose exec -w /hasura hasura hasura-cli seeds apply`.

## Hasura UI

You may access Hasura through `https://localhost:8080`. The secret for authenticating is in your `docker-compose.yaml`.

## Authentication

If you need a JWT token for testing queries as an employer, you may use the command below to get one:

```
curl -i -X POST localhost:9000/auth --data '{"username":"<employer-cnpj>", "password":"<employer-plain-text-password>"}'
```

The token is in a `Set-Cookie` header using the key `pipo-token`.
