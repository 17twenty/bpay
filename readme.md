# BPAY Generator

Golang implementation of a [BPAY Generator written in PHP](https://github.com/fontis/bpay-ref-generator)

## Background

BPAY has become a really easy way to receive payments in Australia so this helps generate the BPAY values in a compliant form.

 1. Register as the BPAY Biller at http://www.bpay.com.au/ -> you'll get your BPAY Biller code then; When your customer is willing to make a BPAY transfer - your application should generate a Customer Reference Number (CRN) programmatically (see below), or you can even generate valid CRNs manually using CRN Generator Tool (win/mac desktop app) - anyway, each CRN generated for a User can be used by him further on, although some applications/services do prefer to generate a new CRN for each new Order a User makes - both approaches are quite fine

 1. You should provide this pair of values to such a Customer:
    * Your BPAY Biller code;
    * The CRN generated/assigned to him;
 1. They then make a BPAY payment using these credentials
 1. You accept it on your end, then you just need to check the CRN value mentioned inside the payment details -> identify your Customer -> process the corresponding Order.

### Customer Reference Number (CRN) generation

CRN is an identifier of a Customer within the BPAY Biller's customer base. This is a number of certain length (that is set up in advance by Biller when registered) that consists of the following two parts:

 1. CRN "seed" - certain numeric value that you choose, representing your Customer or an Order document inside your database. It can be your Customer or Order ID, but we recommend to avoid exposing your internal autoincremented IDs, instead - some randomized public IDs (of certain length, e.g. 12 digits) would better be used.
 1. A Check digit has to be calculated and appended to the value above. For BPAY CRNs specifically this Check digit gets calculated based on the value above - using the Luhn algorithm, mod 10 version 5 (MOD10V05):

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