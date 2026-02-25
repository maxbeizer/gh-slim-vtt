# Slim VTT

A [gh cli](https://cli.github.com/) extension to remove cruft from vtt files (like Zoom and Microsoft Teams transcripts)

## Installation

Install the `gh` cli via [the instructions](https://github.com/cli/cli#installation)

Then install the extension:

```
gh extension install maxbeizer/gh-slim-vtt
```

## Usage

```
gh slim-vtt ~/Downloads/GMT... | pbcopy
```

## Development

```
make help          # list all available targets
make build         # compile the binary
make test          # run unit tests
make ci            # build + vet + test-race
make install-local # install extension from current checkout
```

Made with 💖 (and a little bit of 😩) by @maxbeizer
