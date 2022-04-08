# go-delightful

Delightful CLI output with colors and emojis. Somewhat opinionated as well, since I've built this for my own CLI tools such as [`vegas-credentials`](https://github.com/aripalo/vegas-credentials).

**This tool is NOT meant as structured logger for web applications and such. Instead it only focuses on enabling nice command-line user experience.** Think more of _output_ not _log messages_.

Main features:
- Coloured output by default
- Emoji prefixes (customizable)
- Message structure: `emoji` + `namespace` + `...args`
- Output "level" types
- Respectful of user environment and disables fancy output if one of below set:
  - `NO_COLOR`
  - `<APP-NAME>_NO_COLOR`
  - `TERM=dumb`
- Defaults to text output but supports JSON as well (useful when pairing with CLIs that support `--output=text|json`)



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
