<div align="center">
	<h1><img alt="Crypto logo" src="https://github.com/ggoodwin/discord-crypto-bot/blob/master/cryptocurrency.png" height="300" /><br />
		Discord Crypto Bot
	</h1>

[![Go Reference](https://pkg.go.dev/badge/ggoodwin/discord-crypto-bot.svg)](https://pkg.go.dev/github.com/ggoodwin/discord-crypto-bot) [![Go Version](https://img.shields.io/github/go-mod/go-version/ggoodwin/discord-crypto-bot)](https://go.dev/) [![GoReportCard](https://goreportcard.com/badge/github.com/ggoodwin/discord-crypto-bot)](https://goreportcard.com/report/github.com/ggoodwin/discord-crypto-bot) [![CodeFactor](https://www.codefactor.io/repository/github/ggoodwin/discord-crypto-bot/badge)](https://www.codefactor.io/repository/github/ggoodwin/discord-crypto-bot) [![GitHub Workflow Status](https://img.shields.io/github/actions/workflow/status/ggoodwin/discord-crypto-bot/.github/workflows/go.yml)](https://github.com/ggoodwin/discord-crypto-bot/blob/master/.github/workflows/go.yml) ![Size](https://img.shields.io/github/languages/code-size/ggoodwin/discord-crypto-bot) [![Last Commit](https://img.shields.io/github/last-commit/ggoodwin/discord-crypto-bot)](https://github.com/ggoodwin/discord-crypto-bot/commits/master) [![License](https://img.shields.io/github/license/ggoodwin/discord-crypto-bot)](https://github.com/ggoodwin/discord-crypto-bot/blob/master/LICENSE.md)

</div>
<hr/>

## 🌟 How it works

I use the [`Yahoo!`](https://finance.yahoo.com/crypto/) API to gather real-time cryptocurrency data.

## 📦 Installation and Usage

### Config File

Rename the `config.json.example` file to `config.json` and add your `Discord Token` and `Crypto Ticker` to the file.

```json
{
	"DiscordToken": "Your Token Key",
	"CryptoTicker": "BTC-USD"
}
```

[List of Crypto Tickers](https://finance.yahoo.com/crypto/)

### Go

Make sure you have `Go` installed on your machine.

You can check by running the following command in the `console`

```plain
go version
```

If you don't have `Go` installed, you can download it from [here](https://go.dev/dl/).

### Project setup

Run the following command in the `console`, in the `project directory`

Run the following command to install the dependencies

```plain
go get
```

Then run the following command to tidy the dependencies

```plain
go mod tidy
```

Then run the following command to build the bot

```plain
go build
```

### Project run

Run the following command to run the bot

```plain
./discord-crypto-bot
```

## 💻 Dependencies

- [bwmarrin's DiscordGo](https://github.com/bwmarrin/discordgo)
- [My go-nyse-stocks Package](https://github.com/ggoodwin/go-nyse-stocks)
- [gorilla's WebSocket](github.com/gorilla/websocket)

## 🙇‍♂️ Issues and Contributing

If you find an issue with this library, please report the issue using our [GITHUB-ISSUES] or check out the [SECURITY] details if it is security related.

If you'd like, I welcome any contributions. Please read the [CONTRIBUTING] document then [FORK] this library and submit a [PULL-REQUEST]. Make sure to click `compare across forks` to see your fork.

## ⚖️ License

This project is under the MIT License. See the [LICENSE] file for the full license text.

## 📜 Changes

Check out our [CHANGELOG]

## 👍🏻 Code of Conduct

Please read my [CODE-OF-CONDUCT] before contributing or engaging in discussions.

<!-- Links -->
[LICENSE]: https://github.com/ggoodwin/discord-crypto-bot/blob/master/LICENSE.md
[CHANGELOG]: https://github.com/ggoodwin/discord-crypto-bot/blob/master/CHANGELOG.md
[SECURITY]: https://github.com/ggoodwin/discord-crypto-bot/blob/master/SECURITY.md
[FORK]: https://github.com/ggoodwin/discord-crypto-bot/fork
[PULL-REQUEST]: https://github.com/ggoodwin/discord-crypto-bot/compare
[CODE-OF-CONDUCT]: https://github.com/ggoodwin/discord-crypto-bot/blob/master/CODE_OF_CONDUCT.md
[CONTRIBUTING]: https://github.com/ggoodwin/discord-crypto-bot/blob/master/CONTRIBUTING.md
[GITHUB-ISSUES]: https://github.com/ggoodwin/discord-crypto-bot/issues
