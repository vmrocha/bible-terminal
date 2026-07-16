# Bible Terminal

An offline-first command-line application for reading and searching the Bible
without leaving the terminal.

## Status

The Go CLI is under active development and includes an offline WEBP reader
backed by an embedded SQLite database. Reading output is styled when used in a
terminal and automatically switches to stable plain text when redirected. See
[docs/PLAN.md](docs/PLAN.md) for the product scope, architecture, milestones,
and acceptance criteria.

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

The currently implemented commands include:

```console
bible read John 3
bible read "John 3:16"
bible read "John 3:16-21"
bible read Jn 3:16 --plain
bible read John 3 --next
bible read Matthew 1 --previous
bible read "1 Cor 13"
bible books
bible search "living water"
bible random
bible translations
```

`bible books` lists all canonical names, source codes, and accepted aliases.
Chapter navigation crosses book boundaries, so moving forward from John 21 reads
Acts 1 and moving backward from Matthew 1 reads Malachi 4.

Output adapts to its destination:

```console
bible read "Psalm 23"          # styled terminal output
bible read "Psalm 23" --plain  # stable tab-separated output
bible read "Psalm 23" | less   # automatically plain, with no ANSI escapes
bible read "Psalm 23" --no-color # readable layout without terminal colors
```

`--plain` and `--no-color` are global flags and may appear before or after the
subcommand.

Search works entirely offline and returns verses containing every query token,
ranked by relevance with canonical Scripture order as a stable tie-breaker.
Matching words are emphasized in interactive colored output:

```console
bible search "living water"
bible search "faith hope love" --limit 10
bible search "kingdom of God" --plain
```

The default result limit is 20. Use `--limit` (or `-n`) to request between 1 and
100 results. Punctuation and case do not affect matching. A plain search with no
matches writes no output, making it safe to use in shell pipelines.
Match highlighting never changes `--no-color`, `--plain`, or redirected verse
text.

Discovery commands also work offline:

```console
bible random
bible random --plain
bible translations
```

`bible random` selects uniformly from all 31,103 bundled WEBP verses.
`bible translations` reports the bundled text edition, language, canon, source,
public-domain rights notice, trademark notice, and publisher text policy.

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

## Installation

Tagged releases ship checksummed single-binary archives for macOS and Linux on
Intel/AMD64 and ARM64. See the [installation guide](docs/INSTALL.md) for archive
verification, installation, source builds, and Bash, Zsh, Fish, and PowerShell
completion setup.

## License

No project license has been selected yet. Bible translations have their own
copyright and redistribution terms; no translation should be bundled until its
license and required attribution have been verified and documented.
