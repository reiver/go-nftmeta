# go-nftmeta

Package **nftmeta** provides tool for working with **NFT metadata**, for the Go programming language.

Especially creating the NFT metadata JSON.

## Documention

Online documentation, which includes examples, can be found at: http://godoc.org/github.com/reiver/go-nftmeta

[![GoDoc](https://godoc.org/github.com/reiver/go-nftmeta?status.svg)](https://godoc.org/github.com/reiver/go-nftmeta)

## Example

Here is an example:

```golang
import (
	"encoding/json"

	"github.com/reiver/go-nftmeta"
)

// ...

var metadata nftmeta.MetaData

metadata.SetExternalLink("http://example.com/token/123")
metadata.SetName("peanut-butter-jelly-time")
metadata.SetYouTubeURL("https://youtu.be/eRBOgtp0Hac")

metadata.AppendAttribute(  nftmeta.AttributeString("Bread 1", "Peanut Butter")  )
metadata.AppendAttribute(  nftmeta.AttributeString("Bread 2", "Jelly")  )

// ...

// 'writer' (below) could be http.ResponseWriter
err := json.NewEncoder(writer).Encode(metadata)

```

## Import

To import package **nftmeta** use `import` code like the follownig:
```
import "github.com/reiver/go-nftmeta"
```

## Installation

To install package **nftmeta** do the following:
```
GOPROXY=direct go get https://github.com/reiver/go-nftmeta
```

## Author

Package **nftmeta** was written by [Charles Iliya Krempeaux](http://reiver.link)
