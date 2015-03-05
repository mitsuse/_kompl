# Kompl

[![License](https://img.shields.io/badge/license-MIT-yellowgreen.svg?style=flat-square)][license]
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)][godoc]
[![Wercker](http://img.shields.io/wercker/ci/54e76fead9b14636630d99c6.svg?style=flat-square)][wercker]
[![Coverage](https://img.shields.io/coveralls/mitsuse/kompl.svg?style=flat-square)][coverage]

[license]: http://opensource.org/licenses/MIT
[godoc]: http://godoc.org/github.com/mitsuse/kompl
[wercker]: https://app.wercker.com/project/bykey/1651e84f4992dc9cde16eb1433f9e648
[coverage]: https://coveralls.io/r/mitsuse/kompl

A server for K-best word completion based on N-gram frequency.

## Usage

### Build a predictor

Kompl requires a *predictor* file for word completion.
The file is built from a raw text file like [this](test/wiki.txt) by counting N-grams.

To build a *predictor* file, execute the following command:

```bash
$ kompl build -p <the path to the output predictor file> -n <the order of N-gram> -c <the path to a raw text file>
```

### Run a server for completion

Kompl is a server for K-best word completion.
To run the server, execute the following command:

```bash
$ kompl run -n <the port number> -p <the path to the predictor file>
```

## TODO

- Write test codes more.
- Implement tokenizer for Engish.
- Use a space and time efficient implementation of trie.
- Fall back into lower-order N-grams appropriately.

## License

The MIT License (MIT)

Copyright (c) 2014 Tomoya Kose.

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
