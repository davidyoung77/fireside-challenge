# Connect4
Connect4 implementation in Golang

The purpose of this api is to support a connect4 game.
The api has a single endpoint that recieves an array of numbers that represent tokens placed in order in the columns of the connect 4 game.

## Build
```docker build -t connect4.```

## Serve
```docker run -p 8080:8080 -it connect4```

## Deployment instructions

This provides the initial serverless configuration for deploying the docker image to your chosen cloud platform. The cloud and account specific configuration values still need to be set.

> **Requirements**: Docker. In order to build images locally and push them to ECR, you need to have Docker installed on your local machine. Please refer to [official documentation](https://docs.docker.com/get-docker/).

In order to deploy your service, run the following command

```
sls deploy
```

## endpoints
- /connect4
Method: "POST"
Accepts an array of numbers representing columns in connect4
example curl:
```
curl --location --request POST 'domain/connect4' \
--header 'Content-Type: application/json' \
--data-raw '[0, 1, 1, 2, 3, 2, 2, 3, 3, 4, 3]'
```
possible response values:
- Game not over, add another token
- WINNER: Player 1
- WINNER: Player 2
- DRAW
- Too many tokens