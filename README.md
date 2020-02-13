# fastbin cli

Command line client for [fastbin](https://github.com/jozsefsallai/fastbin).

[![asciicast](https://asciinema.org/a/301276.svg)](https://asciinema.org/a/301276)

## Usage

You can either upload the contents of a file or the output of another command. In both cases, the app will return the URL and the raw URL of the uploaded snippet and will copy the URL to the clipboard.

**To upload the contents of a file:**

```
fastbin-cli file.txt
```

**To upload the output of a command:**

```
ls -ag | fastbin-cli
```

## Installation

### Linux, Mac OS, FreeBSD, Windows

Download the latest [release](https://github.com/jozsefsallai/fastbin-cli/releases) for your operating system and architecture. For easy access, you should place the downloaded executable somewhere inside your PATH (and maybe even rename it too, if you want).

### Go

You can also download a binary using `go get`:

```
go get github.com/jozsefsallai/fastbin-cli
```

## Configuration

You can use the built-in `init` command to initialize your environment settings.

```
fastbin-cli init
```

It will prompt you for the URL of the fastbin server you want to use. The settings will be stored in `~/.fastbinrc.json`. You can also create this file yourself, based on the [example](https://github.com/jozsefsallai/fastbin-cli/blob/master/.fastbinrc.example.json).

## Contribution

Your contribution is most appreciated! If you've found an issue or you want to improve something, feel free to create an issue or submit a PR.

```sh
git clone git@github.com:jozsefsallai/fastbin-cli.git
cd fastbin-cli
go mod tidy
```

## License

MIT.
