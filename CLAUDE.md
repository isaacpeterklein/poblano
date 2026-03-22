# Poblano

A lightweight, file-based static site generator with its own DSL. Write a `.pob` file, run `poblano build`, get a full static website.

## Concept

Users write a single `.pob` file that describes their entire site. Poblano parses it and compiles it into a `dist/` folder with static HTML/CSS/JS.

## CLI Commands

- `poblano new <name>` — scaffolds a new `<name>.pob` file with a starter template
- `poblano build` — compiles the `.pob` file into `dist/`

## .pob File Format

- Sections are separated by **blank lines**
- First line of a section is the **keyword** (page name, component name, or `config`/`header`)
- Subsequent lines are the **content** for that keyword

### Example

```
config
primary #6c3fc4
font Inter
dark-mode true
site-name My Site

header
about home projects contact

home
hero
Welcome to My Site
I build cool things

card
About Me
I am an engineer

about
card
Who I Am
I am an engineer based in NYC
```

### Rules

- `config` block must come first
- `header` defines the navbar — each word becomes a page with its own route
- Pages are defined by their name matching one of the header items
- Components inside a page section: first line after keyword = title/heading, remaining lines = body content
- Sections within a page are separated by blank lines

## Supported Config Keys

| Key | Example | Description |
|-----|---------|-------------|
| `primary` | `#6c3fc4` | Primary brand color |
| `font` | `Inter` | Google Font to use |
| `dark-mode` | `true` | Enable dark mode default |
| `site-name` | `My Site` | Site title in browser tab |

## Supported Components

| Component | Description |
|-----------|-------------|
| `header` | Navbar, auto-generated from page names |
| `hero` | Large banner section with title and subtitle |
| `card` | Card with title and body text |
| `button` | Clickable button |
| `image` | Image element |
| `text` | Plain text / paragraph block |
| `grid` | Arranges child cards in a grid layout |
| `footer` | Page footer |
| `link` | Hyperlink |

## Output

- `dist/index.html` — home page
- `dist/<page>/index.html` — one file per page
- `dist/styles.css` — generated stylesheet based on config
- All pages share the same navbar and config-driven theme

## Tech Stack

- **Language:** Go
- **Output:** Static HTML/CSS/JS (no server required)
- **Design:** Good defaults, configurable via `config` block

## Project Goals

- Simple enough that a non-developer can write a `.pob` file
- Fast builds (single Go binary)
- No dependencies for the end user — just install `poblano` and go
- KISS principle throughout
