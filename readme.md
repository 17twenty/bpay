# BPAY Generator

Golang implementation of a [BPAY Generator written in PHP](https://github.com/fontis/bpay-ref-generator)

## Getting Started

```bash
go get -u github.com/17twenty/bpay
```

Then simply use it as:

```golang

import (
    "github.com/17twenty/bpay"
    ...
)

...
validBPAY := bpay.GenerateMOD10V1("yourInputAccount")
validBPAYv5 := bpay.GenerateMOD10V5("yourInputAccount")

```

Feature requests/PRs welcome.

## See Also

Originally based on a [StackOverflow answer](https://github.com/fontis/bpay-ref-generator)