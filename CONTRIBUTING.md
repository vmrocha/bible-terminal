# Contributing

## Prerequisites

- Go 1.26 or newer
- Make

## Local workflow

Run the complete local validation suite before submitting a change:

```console
make check
```

The individual commands are also available:

```console
make fmt
make lint
make test
make build
./bin/bible version
```

Keep reference parsing, Bible reading behavior, storage, and terminal rendering
in separate packages. New behavior should include tests, and commands should
write data to standard output and diagnostics to standard error.

## Bible text and licensing

Do not commit Bible text or generated translation databases without documented
provenance, redistribution terms, attribution requirements, source version, and
a source checksum. Translation licensing must be reviewed before import work is
merged.
