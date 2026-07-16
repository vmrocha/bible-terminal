# Installing Bible Terminal

## Release archives

Tagged releases provide single-binary archives for:

- macOS on Intel (`darwin_amd64`)
- macOS on Apple silicon (`darwin_arm64`)
- Linux on x86-64 (`linux_amd64`)
- Linux on ARM64 (`linux_arm64`)

Download the archive for your platform and `checksums.txt` from the matching
[GitHub release](https://github.com/vmrocha/bible-terminal/releases). For
example, for version `0.1.0` on an Apple silicon Mac:

The repository is currently private, so the downloading GitHub account must
have access and the GitHub CLI must be authenticated with `gh auth login`.

```console
gh release download v0.1.0 \
  --repo vmrocha/bible-terminal \
  --pattern bible-terminal_0.1.0_darwin_arm64.tar.gz \
  --pattern checksums.txt
grep 'bible-terminal_0.1.0_darwin_arm64.tar.gz$' checksums.txt | shasum -a 256 -c -
tar -xzf bible-terminal_0.1.0_darwin_arm64.tar.gz
install -d ~/.local/bin
install -m 0755 bible ~/.local/bin/bible
```

On Linux, use `sha256sum --check -` in place of the `shasum` verification
command.

Ensure `~/.local/bin` is on your `PATH`, then verify the embedded release
metadata and read a verse:

```console
bible version
bible read "John 3:16"
```

Release archives contain the complete offline WEBP database. No separate text
download, account, API key, or network connection is needed after installation.

## Shell completion

Bible Terminal generates completion scripts without requiring additional files.

### Bash

```console
mkdir -p ~/.local/share/bash-completion/completions
bible completion bash > ~/.local/share/bash-completion/completions/bible
```

### Zsh

```console
mkdir -p ~/.zfunc
bible completion zsh > ~/.zfunc/_bible
```

Add `fpath=(~/.zfunc $fpath)` to `~/.zshrc` before `compinit` is called, then
start a new shell.

### Fish

```console
mkdir -p ~/.config/fish/completions
bible completion fish > ~/.config/fish/completions/bible.fish
```

PowerShell completion is also available through `bible completion powershell`.

## Build from source

Building from source requires Go 1.26 or newer and Make:

```console
gh repo clone vmrocha/bible-terminal
cd bible-terminal
make check
install -d ~/.local/bin
install -m 0755 bin/bible ~/.local/bin/bible
```
