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

- _Log Events_ for tracing/debugging the program execution<br/>→ Use logging frameworks such as [zap](https://github.com/uber-go/zap), [zerolog](https://github.com/rs/zerolog) or [logrus](https://github.com/sirupsen/logrus) for more advanced application logging. You still should ensure they write to `stderr` (or to a rotated log file or something).

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

~~For example usage, see [`vegas-credentials`](https://github.com/aripalo/vegas-credentials) which utilizes this library.~~ (Not yet implemented!)

<br/>

## Install

```sh
go get github.com/aripalo/go-delightful
```

<br/>

## Usage

### Basics

Here are some of the basic methods:
```go
package main

import (
	"github.com/aripalo/go-delightful"
)

func main() {

	// Initialize
	message := delightful.New("demo")

	// Purple message line
	message.Titleln("👷", "Just showing things around here")

	// Gray message line
	message.Infoln("ℹ️", "Good to know")

  // Cyan message (without newline)
  message.Prompt("📝", "Provide input")

  // Green message line
  message.Successln("✅", "Great Success")

  // Red message line
  message.Failureln("❌", "Error")
}
```

### Real world'ish example

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
	flag "github.com/spf13/pflag"
)

func main() {

	// Initialize
	message := delightful.New("greet")

	// Some setup with spf13/pflag, not really important for this example
	var silentMode bool
	var verboseMode bool
	var noEmoji bool
	var noColor bool
	flag.BoolVarP(&silentMode, "silent", "s", false, "silent mode (hides everything except prompt/failure messages)")
	flag.BoolVar(&verboseMode, "verbose", false, "verbose output (show everything, overrides silent mode)")
	flag.BoolVar(&noEmoji, "no-emoji", false, "disable emojis")
	flag.BoolVar(&noColor, "no-color", false, "disable colors and emojis")
	flag.Parse()

	// Configure how messaging works based on above CLI flags
	message.SetSilentMode(silentMode)
	message.SetVerboseMode(verboseMode)
	message.SetEmojiMode(!noEmoji)
	message.SetColorMode(!noColor)

	// Print a "banner" showing your app name and other (optional) info.
	// Banner optional info only printed if in verbose mode.
	message.Banner(delightful.BannerOptions{
		Version: "v0.0.1",
		Website: "http://example.com",
		Command: "hello",
		Extra:   "Some extra info, let's keep it short!",
	})

	// Print "title" message line in purple.
	message.Titleln("🔥", "This is going to be lit!")

	// Print "debug" message line in dark gray.
	// Only printed if in verbose mode.
	message.Debugln("", "This is only visible if in verbose mode") // passing empty string for emoji disables it

	// Print "info" message line in gray.
	message.Infoln("ℹ️", "Just something for your information.")

	// Print "prompt" message in cyan.
	// Does not actually read input, only shows the "question".
	message.Prompt("🙋", "Tell us Your name: ")

	// Actually query the name via stdin
	reader := bufio.NewReader(os.Stdin)
	value, err := reader.ReadString('\n')
	if err != nil {
		// Print "failure" message line in red.
		message.Failure("❌", "galaxies exploded")
		log.Fatal(err)
	}
	name := strings.TrimSpace(value)

	if strings.ContainsRune(name, '💩') {

		// Unfortunately many environments print gendered/toned emojis incorrectly
		// so you might want to use github.com/enescakir/emoji to assign "neutral" emoji
		facepalm := emoji.Emoji(emoji.PersonFacepalming.Tone(emoji.Default))

		// Print "warning" message line in yellow.
		message.Warningln(facepalm, "Really? Your name has a poop emoji? You're being silly...")
	} else {
		// Print "success" message line in green.
		message.Successln("✅", "Name received!")
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

#### Example Output

Results of running above code with `go run main.go` and various flags:

##### No flags
![default](/assets/screenshots/default.png)

##### `--no-emoji`
![emoji disabled](/assets/screenshots/no-emoji.png)

... or with `NO_EMOJI` or `<APP_NAME>_NO_EMOJI` environment variables set (to something other than `false` or `0`).

##### `--no-color`
![color disabled](/assets/screenshots/no-color.png)

... or with `NO_COLOR` or `<APP_NAME>_NO_COLOR` environment variables set (to something other than `false` or `0`).

##### `--silent`
![silent mode](/assets/screenshots/silent.png)

##### `--verbose`
![verbose mode](/assets/screenshots/verbose.png)

... or with `VERBOSE` or `<APP_NAME>_VERBOSE` environment variables set (to something other than `false` or `0`).

<br/>

## Emojis

**Unfortunately** not all command-line environments are capable of rendering all emojis as they should be rendered. To ensure best possible compatibility with different systems and fonts:
- Use non-gendered emojis
- Use default skintone emojis
- Avoid complex group emojis such as 👨‍👩‍👧‍👧

Sometimes using [`enescakir/emoji`](https://github.com/enescakir/emoji) can make getting the exactly right emoji easier:
```go
emoji.Emoji(emoji.PersonFacepalming.Tone(emoji.Default))
```

Of course you may use any emoji you wish but note that they may not render as expected in different systems and environments.

<br/>

## Configuration

**TODO!**

- VERBOSE
- SILENT
- Set* methods
