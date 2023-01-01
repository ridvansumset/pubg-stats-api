# pubg-stats-api

A stats API for the PUBG player, The4nswer. You need to install and run [pubg-stats-ui](https://github.com/ridvansumset/pubg-stats-ui) in order to use this project.

First, you need to have `go` installed to your computer. Visit [this page](https://go.dev/dl/) and choose appropriate Go release for your system.

Then, you'll need to have a PUBG API key. Visit [this page](https://developer.pubg.com/) and create yours if you don't have one.

## Next Steps

Clone the repo:

```
git clone git@github.com:ridvansumset/pubg-stats-api.git
```

Go to app directory:

```
cd pubg-stats-api/
```

Create a .env file and add your PUBG API key in it:

```
touch .env && echo "API_KEY=YOUR_API_KEY" >> .env
```

Download dependencies:

```
go mod download
```

Finally, run the server:

```
make pubg && ./pubg
```
