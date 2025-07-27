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
