# toolbox - Iterate over slices, maps, channels...

[![tag](https://img.shields.io/github/tag/wangliliwang/toolbox.svg)](https://github.com/wangliliwang/toolbox/releases)
![Go Version](https://img.shields.io/badge/Go-%3E%3D%201.20-%23007d9c)
![Build Status](https://github.com/wangliliwang/toolbox/actions/workflows/tests.yml/badge.svg)
[![Go report](https://goreportcard.com/badge/github.com/wangliliwang/toolbox)](https://goreportcard.com/report/github.com/wangliliwang/toolbox)
[![Coverage](https://img.shields.io/codecov/c/github/wangliliwang/toolbox)](https://codecov.io/gh/wangliliwang/toolbox)
[![Contributors](https://img.shields.io/github/contributors/wangliliwang/toolbox)](https://github.com/wangliliwang/toolbox/graphs/contributors)
[![License](https://img.shields.io/github/license/wangliliwang/toolbox)](./LICENSE)

✨ **`wangliliwang/toolbox` is a Lodash-style Go library based on Go 1.18+ Generics and is inspired by [samber/lo](https://github.com/samber/lo) and [lodash](https://lodash.com/).**

## 🚀 Install

```sh
go get github.com/wangliliwang/toolbox
```

## 💡 Usage

You can import `toolbox` using:

```go
import (
    "github.com/wangliliwang/toolbox"
)
```

Then use one of the helpers below:

```go
ch := toolbox.SliceToChannel[string]([]string{"Samuel", "John", "Samuel"})
```