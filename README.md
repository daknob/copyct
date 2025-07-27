# copyct
For when you need a roots.pem and you trust your friends ;)

When
[operating](https://ipng.ch/s/articles/2025/07/26/certificate-transparency-part-1/)
a [CT Log](https://certificate.transparency.dev/), you need a list of Root CAs
that you will accept certificate chains from. As there's currently no single
source of truth, and there are [many](https://github.com/daknob/root-programs)
Root Programs and Trust Stores to choose from, operators are usually creating a
union of Apple, Mozilla, and Google, plus a few extra CAs used for monitoring
the uptime of the service.

As more and more operators choose to run these logs, eventually we'll get
duplicated work, so this tool helps you copy the `roots.pem` file from an
operator you trust. Point this binary to the hostname of a CT Log and it will
create a file containing all of the certificates that are trusted.

## Usage

You can run `copyct` with `-h` to get the help:

```
Usage of copyct:
  -c string
        Validate Root CAs: (no|warn|fail|remove) (default "warn")
  -l string
        CT Log URL (default "tuscolo2026h2.sunlight.geomys.org")
  -o string
        Output PEM file (default "roots.pem")
```

The CT Log URL and the output file name are self-explanatory. For the `-c`
option `copyct` can attempt to parse each certificate and ensure it is valid.
Here are the various modes of operation:

* `no` will skip validation and won't even check if each entry is a certificate
* `warn` will warn if there's an error, but will still add it to the output
* `fail` will not export a file if any of these certificates is invalid
* `remove` will export only valid certificates and skip invalid ones

Currently the default is `warn` as there's a single Root CA with a negative
serial number, which is invalid according to Go's `x509` package.
