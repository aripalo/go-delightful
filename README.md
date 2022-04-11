🚧 🚧 🚧  **Work-in-Progress**

This is still under heavy construction. **Do NOT use just yet!**

---

<br/>
<br/>


<div align="center">
	<br/>
	<br/>
	<img width="200" src="assets/go-delightful.svg" alt="Got" />
  <h1>
  <code>go-delightful</code>
  <br/>
  <span>command-line messaging with colors and emojis</span>
  </h1>
  <br/>
</div>

Go library designed for command-line applications interacting with humans, providing delightful end-user experience for _informational messaging_ via colourful, lightly structured messages with emojis.

<br/>

✅ &nbsp;**This tool is for _Informational messaging_ for the _end-user_ while the command is running:**

Messages are written into standard error stream (`stderr`) to avoid polluting the actual program output in standard output stream (`stdout`); Hence the messages are intended for end-user communication during the program runtime, not for the actual parseable result output:

> _Not everything on stderr is an error though. For example, you can use `curl` to download a file but the progress output is on stderr. This allows you to redirect the stdout while still seeing the progress._
>
> _In short: `stdout` is for **output**, `stderr` is for **messaging**._
>
> – [12 Factor CLI Apps](https://medium.com/@jdxcode/12-factor-cli-apps-dd3c227a0e46)

<br/>

❌ &nbsp;**This tool is _NOT_ for:**

- _Log Events_ for tracing/debugging the program execution<br/>→ Use logging frameworks such as [zap](https://github.com/uber-go/zap), [zerolog](https://github.com/rs/zerolog) or [logrus](https://github.com/sirupsen/logrus) for more advanced application logging. You still should ensure they write to `stderr`.

- _Command Output_ (i.e. the result) written into `stdout` (which could be redirected to a file for example)<br/>→ Use [`fmt.Print`](https://pkg.go.dev/fmt#Print) for that which defaults to writing into standard output.



<br/>

## Features
- Prints to Standard Error stream (`stderr`)
- Subprocess friendly: Tries to access the `tty` and print to its `stderr` (can be disabled)
- Message structure:
  1. `emoji` (optional, pass empty string to ignore it)
  2. `...args` (often just a single string)
- Coloured output by default with emojis
- Respectful of user environment and disables color output with emojis if one of below set:
  - `NO_COLOR` (to something other than `false` or `0`)
  - `<APP-NAME>_NO_COLOR` (to something other than `false` or `0`)
  - `TERM=dumb`
- Allows user to control verbosity by configuring environment variables:
  - `VERBOSITY` (to something other than `false` or `0`)
  - `<APP-NAME>_VERBOSITY` (to something other than `false` or `0`)

<br/>

For example usage, see [`vegas-credentials`](https://github.com/aripalo/vegas-credentials) which utilizes this library.

<br/>

## Install

```sh
go get github.com/aripalo/go-delightful
```

<br/>

## Usage

An example command-line application using `aripalo/go-delightful`:
```go
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/aripalo/go-delightful"
	"github.com/enescakir/emoji"
)

func main() {

	message := delightful.New("greet")

	// Print a "banner" showing your app name and other (optional) info.
	// Banner optional info only printed if in verbose mode.
	message.Banner(delightful.BannerOptions{
		Version: "v0.0.1",
		Website: "http://example.com",
		Command: "hello",
		Extra:   "Some extra info, let's keep it short!",
	})

	// Print "title" message in purple.
	message.Title("🔥", "This is going to be lit!")

	// Print "debug" message in dark gray.
	// Only printed if in verbose mode.
	message.Debug("⚙️", "This is only visible if in verbose mode")

	// Print "info" message in gray.
  // Passing empty string for emoji disables it.
	message.Info("", "FYI: Just something for for your information.")

	// Print "prompt" message in cyan.
	// Does not actually read input, only shows the "question".
	message.Prompt("🖋️", "Input your name:")

	// Actually query the name via stdin
	reader := bufio.NewReader(os.Stdin)
	value, err := reader.ReadString('\n')
	if err != nil {
		// Print "failure" message in red.
		message.Failure("❌", "galaxies exploded")
		log.Fatal(err)
	}
	name := strings.TrimSpace(value)

	if strings.ContainsRune(name, '💩') {

		// Unfortunately many environments print gendered/toned emojis incorrectly
		// so you might want to use github.com/enescakir/emoji to assign "neutral" emoji
		facepalm := emoji.Emoji(emoji.PersonFacepalming.String())

		// Print "warning" message in yellow.
		message.Warning(facepalm, "Really? Your name has a poop emoji? You're being a wise-ass...")
	} else {
		// Print "success" message in green.
		message.Success("✅", "Name received!")
	}

	// Print horizontal ruler.
	// Useful for visually separating the informational messages above from
	// actual command output.
	// Visible only on verbose mode.
	message.HorizontalRuler()

	// Finally you often should print some actual command output into standard
	// output stream.
	fmt.Printf("Hello %s!\n", name)
}
```

Now running `VERBOSE=true go run main.go` yields to following:
![Example Output](/assets/example-output.png)


## Emojis

Unfortunately not all command-line environments are capable of rendering all emojis as they should be rendered. To ensure best possible compatibility with different systems and fonts:
- Use non-gendered emojis
- Default skintone emojis
- Avoid complex group emojis such as 👨‍👩‍👧‍👧


## Silent

TODO --silent mode

only:
- prompt
- error
