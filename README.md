# Not Another Markup Language, or nano-YAML

[![Go reference](https://pkg.go.dev/badge/github.com/andreyvit/naml.svg)](https://pkg.go.dev/github.com/andreyvit/naml) ![Zero dependencies](https://img.shields.io/badge/deps-zero-brightgreen) ![&lt;100 LOC](https://img.shields.io/badge/size-%3C100%20LOC-green) ![100% coverage](https://img.shields.io/badge/coverage-100%25-green) [![Go Report Card](https://goreportcard.com/badge/github.com/andreyvit/naml)](https://goreportcard.com/report/github.com/andreyvit/naml)

NAML is a subset of YAML that only requires a JSON parser (and a tiny conversion func).

If you ever thought, _I want something like a simple YAML parser_ but _man those YAML parsers are monstrous_, this solution is for you.

The format is basically HTTP headers where all values are JSON:

```yaml
template: "home"
page_classes: ["page--jumbo"]
cta: {
    "primary": {
      "title": "Click me",
      "href": "http://example.com/"
    }
  }
```

Syntax:

1. Each non-empty line is a `key: value` pair.
2. Each value is a valid JSON.
3. You can continue a value on multiple lines by indenting all continuation lines (using spaces or tabs).
4. `#` starts a comment if it's the first non-whitespace character on the line.

That's it.

This is valid YAML, and can be trivially converted to an actual JSON, so that you can use all features of your favorite JSON parser.


## Usage

Copy [`naml.go`](naml.go) into your project. Not worth adding a dependency over it. This is a good example of a piece of code that's _done_ and not expected to ever change.

But, of course, if your project is already a dumpster fire, nobody is stopping you from doing `go get github.com/andreyvit/naml@latest`.

Then you just do `json.Unmarshal(naml.Convert(data), ...)`.

For bonus points, combine it with [andreyvit/jsonfix](https://github.com/andreyvit/jsonfix) to allow trailing commas in that inline JSON.


## Contributing

Umm. Well if you need this to run on older versions of Go and backport to not use `bytes.Lines`, please share.

Bug fixes are welcome too, of course.


## Dual License

Copyright (c) 2025, Andrey Tarantsov.

I release this code under a MIT license and a Zero-Clause BSD license.

Zero-Clause BSD license allows usage of this code without attribution and without preserving the copyright notice — basically just copy it into your project.
