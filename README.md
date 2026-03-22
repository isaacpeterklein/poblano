# Poblano 🌶️

A lightweight, file-based static site generator with its own simple syntax. Write a single `.pob` file describing your entire site, run one command, and get a fully built static website.

No config files. No dependencies. No framework knowledge required. Just write and build.

Named after Poblano the dog — a puggle, a beagle-pug mix. He is great.

<img src="https://i.ibb.co/08T9FLr/IMG-4905.jpg" width="280"/>

---

## Install

**One-liner:**

```bash
curl -fsSL https://raw.githubusercontent.com/isaacpeterklein/poblano/master/install.sh | sh
```

**Or build from source:**

```bash
git clone https://github.com/isaacpeterklein/poblano
cd poblano
go build -o poblano .
mv poblano /usr/local/bin/poblano
```

---

## Quick Start

```bash
poblano new mysite     # create a starter mysite.pob file
poblano build          # compile to dist/
poblano serve          # build, serve at localhost:3000, and watch for changes
```

---

## CLI Commands

| Command | Description |
|---|---|
| `poblano new <name>` | Scaffold a new `<name>.pob` starter file |
| `poblano build` | Build the `.pob` file in the current directory |
| `poblano build <file>` | Build a specific `.pob` file |
| `poblano serve` | Build, serve at localhost:3000, and watch for changes |
| `poblano serve <file>` | Serve a specific `.pob` file |

Output is written to a `dist/` folder. The folder is wiped and rebuilt on every run.

---

## The .pob File Format

Everything lives in one `.pob` file. Sections are separated by **blank lines**. The first line of each section is the keyword. Lines after that are the content.

```
config
primary #6d28d9
accent #a855f7
font Inter
dark-mode false
site-name My Site

header
home about projects contact

home
hero
Welcome to My Site
I build cool things.

card
About Me
I am an engineer based in NYC.

about
card
Who I Am
More detail about me here.

contact
card
Get In Touch
you@example.com

footer
My Site — Built with Poblano
```

### Rules

- `config` should come first
- `header` defines the navbar — each word becomes a page and a nav link
- Page sections match their name to a header item
- Components inside a page are separated by blank lines
- The first line after a component keyword is the **title**
- Remaining lines are the **body**
- Poblano warns at build time if a nav item has no matching page section, or if an unknown component or config key is used

---

## Config Options

| Key | Default | Description |
|---|---|---|
| `primary` | `#4f46e5` | Primary brand color (hex) |
| `accent` | `#7c3aed` | Accent color for nav hover/active states |
| `font` | `Inter` | Google Font family name |
| `dark-mode` | `false` | Set to `true` for a dark theme |
| `site-name` | `My Site` | Site title shown in the browser tab |
| `logo` | — | Path or URL to a logo image shown in the navbar |
| `favicon` | — | Path or URL to a favicon |

```
config
primary #e11d48
accent #fb7185
font Poppins
dark-mode false
site-name My Portfolio
logo /logo.png
favicon /favicon.ico
```

---

## Components

### `hero`
A large banner section. First line is the heading, second is the subtitle.

```
hero
Hello, I'm Isaac
I build things for the web.
```

---

### `card`
A content card with a title and body.

```
card
My Project
A short description of what this project does.
```

---

### `grid`
A responsive grid of cards. Each line is one card: `Title | Body | Button Label > URL`. The button is optional.

```
grid
Project One | A cool project I built. | View Demo > https://example.com
Project Two | Another thing I made.
Project Three | No button on this one.
```

---

### `gallery`
A responsive image grid. Each line is: `image-src | alt text`.

```
gallery
/images/photo1.jpg | Mountains
/images/photo2.jpg | Forest
https://example.com/image.jpg | Sunset
```

---

### `heading`
A large section heading with an optional subtitle below it.

```
heading
My Section Title
An optional subtitle or description here.
```

---

### `text`
A plain paragraph of text.

```
text
This is a paragraph that will appear on the page.
```

---

### `code`
A styled code or syntax block. Everything after the keyword is displayed verbatim. Note: blank lines inside a code block will end the block — use indentation instead.

```
code
poblano new mysite
poblano build
poblano serve
```

---

### `button`
A styled button. First line is the label, second is the URL.

```
button
View My Work
https://github.com/yourusername
```

---

### `link`
An inline hyperlink. First line is the label, second is the URL.

```
link
Visit my GitHub
https://github.com/yourusername
```

---

### `image`
An image. First line is the src, second is the alt text.

```
image
/images/photo.jpg
A photo of me
```

---

### `divider`
A horizontal rule to separate sections. Takes no content.

```
divider
```

---

### `footer`
A site footer.

```
footer
© 2026 My Site — Built with Poblano
```

---

## Output Structure

```
dist/
├── index.html
├── about.html
├── projects.html
├── contact.html
└── styles.css
```

Drop the `dist/` folder onto any static host.

---

## Hosting

- [Netlify](https://netlify.com) — drag and drop `dist/`
- [Vercel](https://vercel.com)
- [GitHub Pages](https://pages.github.com)
- [Surge](https://surge.sh) — `surge dist/`
- Any web server that serves static files

---

## Full Example

```
config
primary #6d28d9
accent #a855f7
font Inter
dark-mode false
site-name Jane Doe

header
home about projects contact

home
hero
Hi, I'm Jane
A software engineer who loves building for the web.

card
What I Do
Full-stack web development, clean UIs, fast backends.

button
See My Work
projects.html

about
heading
My Background
A little about me.

text
I've been building software for 8 years, working with startups and large companies alike.

card
Skills
JavaScript, Go, Python, React, PostgreSQL, Docker

projects
hero
Projects
Things I've built.

grid
Portfolio Site | The site you're looking at — built with Poblano. | View Source > https://github.com/janedoe
Open Source CLI | A developer tool for automating repetitive tasks. | View Source > https://github.com/janedoe

contact
card
Say Hello
jane@example.com

link
GitHub
https://github.com/janedoe

footer
© 2026 Jane Doe — Built with Poblano
```

---

## Why Poblano?

Most site generators require you to learn a framework, install dependencies, and write config files before you see anything. Poblano skips all of that.

Write a text file. Run one command. Get a website.
