<!--
  Attractive html formatting for rendering in github. sorry text editor
  readers! Besides the header and section links, everything should be clean and
  readable.
-->
<h1 align="center">framerr</h1>
<p align="center"><i>Extract frame information from <a href="https://golang.org">Go</a> error values</i></p>

<div align="center">
  <a href="https://godoc.org/github.com/jbowes/framerr"><img src="https://godoc.org/github.com/jbowes/framerr?status.svg" alt="GoDoc"></a>
  <img alt="Alpha Quality" src="https://img.shields.io/badge/status-ALPHA-orange.svg" >
  <a href="https://travis-ci.com/jbowes/framerr"><img alt="Build Status" src="https://travis-ci.com/jbowes/framerr.svg?branch=master"></a>
  <a href="https://github.com/jbowes/framerr/releases/latest"><img alt="GitHub tag" src="https://img.shields.io/github/tag/jbowes/framerr.svg"></a>
  <a href="./LICENSE"><img alt="BSD license" src="https://img.shields.io/badge/license-BSD-blue.svg"></a>
  <a href="https://codecov.io/gh/jbowes/framerr"><img alt="codecov" src="https://img.shields.io/codecov/c/github/jbowes/framerr.svg"></a>
  <a href="https://goreportcard.com/report/github.com/jbowes/framerr"><img alt="Go Report Card" src="https://goreportcard.com/badge/github.com/jbowes/framerr"></a>
</div><br /><br />


## Introduction
Introduction | [Usage] | [Contributing] <br /><br />

🚧 ___Disclaimer___: _`framerr` is alpha quality software. The API may change
without warning between revisions._ 🚧

`framerr` extracts stack frame information for Go2 error values


## Usage
[Introduction] | Usage | [Contributing] <br /><br />

As Go2 error value stack frame information is not programatically accessible,
you can use `framerr` to extract a slice of frames from an error chain.
```go
err1 := xerrors.New("an error")
err2 := cling.Wrap(err1, "wrapped")

// ...

frames := framerr.Extract(err)
```

Each `Frame` holds the error message, and a `Source` field that contains file,
package, function, and line information. Use these details to send stack info
to your favorite error tracking service!

For complete examples and usage, see the [GoDoc documentation](https://godoc.org/github.com/jbowes/framerr).

## Contributing
[Introduction] | [Usage] | Contributing <br /><br />

I would love your help!

`framerr` is still a work in progress. You can help by:

- Opening a pull request to resolve an [open issue][issues].
- Adding a feature or enhancement of your own! If it might be big, please
  [open an issue][enhancement] first so we can discuss it.
- Improving this `README` or adding other documentation to `framerr`.
- Letting [me] know if you're using `framerr`.


<!-- Section link definitions -->
[introduction]: #introduction
[examples]: #examples
[usage]: #usage
[contributing]: #contributing

<!-- Other links -->
[go]: https://golang.org

[issues]: ./issues
[bug]: ./issues/new?labels=bug
[enhancement]: ./issues/new?labels=enhancement

[me]: https://twitter.com/jrbowes
