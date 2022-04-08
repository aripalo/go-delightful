<div align="center">
	<br/>
	<br/>
	<img width="200" src="assets/go-delightful.svg" alt="Got" />
  <h1>
  <code>go-delightful</code>
  <br/>
  <span>output with colors and emojis</span>
  </h1>
  <div>

  Go library providing delightful & pretty command-line output with colors and emojis.

  </div>
  <hr/>
  <br/>
</div>

**This tool is NOT meant to be used as structured logger for web applications and such. Instead it only focuses on enabling nice command-line user experience: Think _informational output_ not _log messages_.**

It's inteded for Command-Line applications that want to interact with humans and provide delightful user experience with coloured output, emojis and such. An example could be [`vegas-credentials`](https://github.com/aripalo/vegas-credentials) which utilizes this library. This tool is also somewhat opinionated and does not follow conventional application loggers.


## Features
- By default prints to tty/stderr
- Coloured output by default
- Emoji prefixes (customizable)
- Message structure: `emoji` + `namespace` + `...args`
- Output "level" types
- Respectful of user environment and disables fancy output if one of below set:
  - `NO_COLOR`
  - `<APP-NAME>_NO_COLOR`
  - `TERM=dumb`



## Install

```sh
go get github.com/aripalo/go-delightful
```

## Usage

```go
package main

import (
  "github.com/aripalo/go-delightful"
)

func main() {

  output := delightful.New("my-app")

  // print a banner showing your app name and other (optional) info
  output.Banner(output.BannerOptions{
    version: "development",
    url:     "example.com",
    extra:   "this is\nmy app",
  })

  // print title message
  output.Title("🔥", "foo", "hello")

  // passing empty string for emoji disables the emoji prefix
  output.Info("", "foo", "world")

  // does not actually read input, only shows the "question"
  output.Prompt("🖌", "foo", "please input maybe:")

  // print success message
  output.Success("💪", "foo", "nicely done")
}
```


## Actual Command Output

```go
output.Output()
```

```go
output.OutputJSON()
```
