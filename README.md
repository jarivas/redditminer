# Reddit Miner

## How to use

### Clone this repository
```bash
git clone https://github.com/jarivas/redditminer
```

### Create required files
```bash
cp .env.example .env
cp test-read.json subreddits.json
```

## Edit the files
**.env**
```bash
REDDIT_USERNAME=reddit_bot
REDDIT_PASSWORD=snoo
REDDIT_CLIENT_ID=p-jcoLKBynTLew
REDDIT_APP_SECRET=gko_LXELoV07ZBNUXrvWZfzE3aI
MONGO_HOST=mongo
MONGO_PORT=27017
MONGO_PORT_UI=8081
MONGO_DB_NAME=reddit
MONGO_ROOT_USERNAME=root
MONGO_ROOT_PASSWORD=example
```
**.subreddits.json**
```json
{
    "subreddits": [
        "AITAH",
        "AmItheAsshole",
        "AmITheDevil",
        "MaliciousCompliance"
    ]
}
```

## Compile and run
```bash
docker run --rm -v "$PWD":/usr/src/myapp -w /usr/src/myapp golang:1.23 go build -v 
```

## Optional run it with a network where there a mongo service

```bash
docker run --network=redditminer --rm -v .:/usr/src/myapp -w /usr/src/myapp golang:1.23 go build -v
```