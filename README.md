# Slim VTT

A [gh cli](https://cli.github.com/) extension to remove cruft from vtt files (like Zoom transcripts)

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

This extension uses the base CLI token so you should not need any special auth.

```
gh extension install .
make build
gh slim-vtt --help
```

Made with 💖 (and a little bit of 😩) by @maxbeizer
