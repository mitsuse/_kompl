# Kompl

[![License](https://img.shields.io/badge/license-MIT-yellowgreen.svg?style=flat-square)][license]

[license]: http://opensource.org/licenses/MIT

(deprecated) A server for K-best word completion based on N-gram frequency.

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

## Installation

The executable binaries are downloadable from the [release page][release page].

[release page]: https://github.com/mitsuse/kompl/releases

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
