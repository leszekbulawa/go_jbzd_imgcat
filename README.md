# JBZD imgcat
This small script prints random image from jbzd.pl main page to the terminal.
Rewritten from python for educational purposes: https://github.com/leszekbulawa/jbzd_imgcat

Currently it only supports iTerm2.

## Source
https://github.com/leszekbulawa/go_jbzd_imgcat

## Installation
1. Make sure that go `/bin` is added to the `PATH`:

`export PATH=$PATH:$(go env GOPATH)/bin`

2. Install package

`go install github.com/leszekbulawa/go_jbzd_imgcat@latest`

## Usage
`go_jbzd_imgcat`

## TODO
- improve range and randomness of images
- support more terminal emulators
- exception handling and code quality :)
- write a test
- pin goquery version
