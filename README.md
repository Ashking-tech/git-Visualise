# git-visualize

A CLI tool to visualize your Git contribution graph in the terminal.

![Contribution Graph](img.png)

## Installation

### Quick Install (Linux)

```bash
curl -sSL https://raw.githubusercontent.com/Ashking-tech/git-Visualise/main/install.sh | bash
```

### From Source

```bash
go install
```

## Usage

### Add repositories to track

```bash
git-visualize -add /path/to/your/projects
```

### View your contribution stats

```bash
git-visualize -email "your@email.com"
```

Find your git email with:
```bash
git config --global user.email
```

## Features

- Scans directories recursively for Git repositories
- Tracks commits by email over the last 6 months
- Displays contribution graph similar to GitHub
- Color-coded commit intensity
