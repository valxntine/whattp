# whattp - A Simple, Offline, Easy to Use HTTP status code explorer.

![whattp](/doc/static/whattpbanner.png)

___


`whattp` is a CLI developed using Go which allows you to quickly get information on a specific HTTP status code with a single command.

## Installation

`whattp` is distributed using Brew, so is only available to those with access to homebrew.

To install `whattp` first tap the repo:  
`brew tap valxntine/valxntine`

Then install `whattp`:  
`brew install whattp`

Alternatively, if you have Go installed you can install `whattp` using `go install`

`go install github.com/valxntine/whattp@latest`

## Usage

Using `whattp` is a breeze, just use the HTTP status code you're interested in as the first and only argument to the command, for example:
`whattp 500`

And you'll be greeted with an output explaining what that status code means.

![whattpexample](/doc/static/whattpexample.png)
