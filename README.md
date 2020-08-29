# wadl-dumper

Dump all available paths and/ endpoints on WADL file.

## Installation

- Download a prebuilt binary from [releases page](https://github.com/dwisiswant0/wadl-dumper/releases/latest), unpack and run! or
- If you have go1.13+ compiler installed: `go get github.com/dwisiswant0/wadl-dumper`

## Usage

It's very simple!

```bash
▶ wadl-dumper -i https://domain.tld/application.wadl
```

### Flags

```
Usage:
  wadl-dumper -i http://domain.tld/application.wadl [options...]
  wadl-dumper -i /path/to/wadl.xml --show-base -r "-alert(1)-"

Options:
  -i, --input <URL/FILE>         URL/path to WADL file
  -b, --show-base                Add base URL to paths
  -r, --replace <string>         Replace all placeholder with given value
  -h, --help                     Show its help text
```

### Show Base URL

By default, `wadl-dumper` doesn't add a base URL to each paths, use  `-b/--show-base` flag to add it.

```bash
▶ wadl-dumper -i https://beacon.netflix.com/application.wadl -b
http://nmtracking.netflix.com/users/customerevents
http://nmtracking.netflix.com/nm
http://nmtracking.netflix.com/tracking
http://nmtracking.netflix.com/itracking
http://nmtracking.netflix.com/track
http://nmtracking.netflix.com/{subResources:.*}
http://nmtracking.netflix.com/users
http://nmtracking.netflix.com/presentationtracking
```

### Replace Placeholder

You can also replace all **{placeholder}** values in each paths with `-r/--replace` flag.

```bash
▶ wadl-dumper -i https://tw.bid.yahoo.com/api/application.wadl -b -r "-alert(1)-"
http://auc-papi.bid.yahoo.com:4080/api/api/-alert(1)-
http://auc-papi.bid.yahoo.com:4080/api/campaigns/-alert(1)-
http://auc-papi.bid.yahoo.com:4080/api/escrows/-alert(1)-
http://auc-papi.bid.yahoo.com:4080/api/p2ppListing/-alert(1)-
http://auc-papi.bid.yahoo.com:4080/api/channels/-alert(1)-
http://auc-papi.bid.yahoo.com:4080/api/channels/-alert(1)-/readInfo
http://auc-papi.bid.yahoo.com:4080/api/channels/-alert(1)-/members
http://auc-papi.bid.yahoo.com:4080/api/users/-alert(1)-/qnas
http://auc-papi.bid.yahoo.com:4080/api/users/-alert(1)-
http://auc-papi.bid.yahoo.com:4080/api/users/-alert(1)-/ranking
http://auc-papi.bid.yahoo.com:4080/api/users/-alert(1)-/rating
...
```

## Supporting Materials/References

- [Leveraging Exposed WADL XML in Burp Suite](https://www.nopsec.com/leveraging-exposed-wadl-xml-in-burp-suite/)

## License

`wadl-dumper` is released under MIT. See `LICENSE`.