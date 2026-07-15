# Bible Terminal

An offline-first command-line application for reading and searching the Bible
without leaving the terminal.

## Status

The Go CLI is under active development and includes an offline WEBP reader
backed by an embedded SQLite database. See [docs/PLAN.md](docs/PLAN.md) for the
product scope, architecture, milestones, and acceptance criteria.

## Product direction

The initial interface will be command-oriented and friendly to both people and
shell pipelines:

```console
bible read John 3
bible read "John 3:16"
bible read "John 3:16-21" --plain
bible search "for God so loved"
bible books
bible random
```

The currently implemented reading forms are:

```console
bible read John 3
bible read "John 3:16"
bible read "John 3:16-21"
bible read Jn 3:16 --plain
```

The first release should:

- work completely offline;
- start quickly and ship as a single executable;
- understand common Bible reference formats and book aliases;
- produce readable terminal output and clean redirected output;
- use Bible text that is legally redistributable; and
- leave room for additional languages and translations.

## Proposed stack

- Go
- Cobra for CLI command parsing
- SQLite with full-text search
- Lip Gloss for terminal presentation
- Bubble Tea later, if a full-screen reader proves useful

## Development

The project requires Go 1.26 or newer and Make. Run the complete local check:

```console
make check
./bin/bible version
```

See [CONTRIBUTING.md](CONTRIBUTING.md) for the development workflow and Bible
text licensing requirements.

## License

No project license has been selected yet. Bible translations have their own
copyright and redistribution terms; no translation should be bundled until its
license and required attribution have been verified and documented.
