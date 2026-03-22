# Poblano 🌶️

A lightweight, file-based static site generator with its own simple syntax. Write a single `.pob` file describing your entire site, run one command, and get a fully built static website.

No config files. No dependencies. No framework knowledge required. Just write and build.

---

## Install

```bash
git clone https://github.com/yourusername/poblano
cd poblano
go build -o poblano .
```

Move the binary somewhere on your PATH:

```bash
mv poblano /usr/local/bin/poblano
```

---

## Quick Start

```bash
poblano new mysite
# edit mysite.pob
poblano build
# open dist/index.html
```

---

## CLI Commands

| Command | Description |
|---|---|
| `poblano new <name>` | Scaffold a new `<name>.pob` starter file |
| `poblano build` | Build the site from the `.pob` file in the current directory |

Output is written to a `dist/` folder.

---

## The .pob File Format

Everything lives in one `.pob` file. Sections are separated by **blank lines**. The first line of each section is the keyword (a component name, page name, `config`, or `header`). Lines after that are the content.

### Basic structure

```
config
primary #4f46e5
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

- The `config` block should come first
- The `header` block defines the navbar — each word becomes a page and a nav link
- Page sections are identified by their name matching a nav item
- Components inside a page are separated by blank lines
- The first line after a component keyword is the **title**
- Remaining lines are the **body**

---

## Config Options

Defined in the `config` block at the top of your `.pob` file.

| Key | Default | Description |
|---|---|---|
| `primary` | `#4f46e5` | Primary brand color (hex) |
| `font` | `Inter` | Google Font family name |
| `dark-mode` | `false` | Set to `true` for a dark theme |
| `site-name` | `My Site` | Site title shown in the browser tab |

Example:

```
config
primary #e11d48
font Poppins
dark-mode true
site-name Isaac's Portfolio
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

### `button`
A styled button. First line is the label, second is the URL.

```
button
View My Work
https://github.com/yourusername
```

---

### `image`
An image. First line is the src path or URL, second is the alt text.

```
image
/images/photo.jpg
A photo of me
```

---

### `text`
A plain paragraph block.

```
text
This is a paragraph of text that will appear on the page.
```

---

### `grid`
A responsive grid container for laying out cards side by side.

```
grid
My Projects
```

---

### `link`
A hyperlink. First line is the label, second is the URL.

```
link
Visit my GitHub
https://github.com/yourusername
```

---

### `footer`
A footer at the bottom of the page.

```
footer
© 2026 Isaac — Built with Poblano
```

---

## Output Structure

Running `poblano build` generates a `dist/` folder:

```
dist/
├── index.html        # home page
├── styles.css        # generated stylesheet
├── about/
│   └── index.html
├── projects/
│   └── index.html
└── contact/
    └── index.html
```

Each page gets its own folder so URLs are clean (e.g. `/about/` instead of `/about.html`). Just drop the `dist/` folder onto any static host.

---

## Hosting

The `dist/` output is plain static HTML — host it anywhere:

- [Netlify](https://netlify.com) — drag and drop the `dist/` folder
- [Vercel](https://vercel.com)
- [GitHub Pages](https://pages.github.com)
- [Surge](https://surge.sh) — `surge dist/`
- Any web server that can serve static files

---

## Full Example

```
config
primary #6d28d9
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
I specialize in full-stack web development, with a focus on clean UIs and fast backends.

button
See My Work
/projects/

about
card
My Background
I've been building software for 8 years, working with startups and large companies alike.

card
Skills
JavaScript, Go, Python, React, PostgreSQL, Docker

projects
card
Portfolio Site
The site you're looking at — built with Poblano.

card
Open Source CLI
A developer tool for automating repetitive tasks.

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

Most site generators require you to learn a framework, set up a project structure, install dependencies, and write config files before you see anything. Poblano skips all of that.

Write a text file. Run one command. Get a website.

It's named after my dog, Poblano or Pobby for short. Sometimes we call him Pob. He is a puggle, a beagle-pug mix. He is great.

<img src="https://i.ibb.co/08T9FLr/IMG-4905.jpg" width="300"/>
