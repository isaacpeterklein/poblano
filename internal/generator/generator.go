package generator

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"poblano/internal/components"
	"poblano/internal/parser"
)

func Build(site *parser.Site, outDir string) error {
	if err := os.MkdirAll(outDir, 0755); err != nil {
		return err
	}

	css := generateCSS(site.Config)
	if err := os.WriteFile(filepath.Join(outDir, "styles.css"), []byte(css), 0644); err != nil {
		return err
	}

	for _, page := range site.Pages {
		html := generatePage(site, page)
		var dir string
		if strings.ToLower(page.Name) == "home" {
			dir = outDir
		} else {
			dir = filepath.Join(outDir, strings.ToLower(page.Name))
			if err := os.MkdirAll(dir, 0755); err != nil {
				return err
			}
		}
		if err := os.WriteFile(filepath.Join(dir, "index.html"), []byte(html), 0644); err != nil {
			return err
		}
	}

	fmt.Printf("Built %d pages to %s/\n", len(site.Pages), outDir)
	return nil
}

func generatePage(site *parser.Site, page parser.Page) string {
	nav := generateNav(site.NavItems, page.Name)
	body := generateComponents(page.Components)

	darkClass := ""
	if site.Config.DarkMode {
		darkClass = " dark"
	}

	return fmt.Sprintf(`<!DOCTYPE html>
<html lang="en" class="%s">
<head>
  <meta charset="UTF-8" />
  <meta name="viewport" content="width=device-width, initial-scale=1.0" />
  <title>%s - %s</title>
  <link rel="preconnect" href="https://fonts.googleapis.com" />
  <link href="https://fonts.googleapis.com/css2?family=%s:wght@400;600;700&display=swap" rel="stylesheet" />
  <link rel="stylesheet" href="%s" />
</head>
<body>
%s
<main class="main">
%s
</main>
</body>
</html>`,
		darkClass,
		capitalize(page.Name),
		site.Config.SiteName,
		strings.ReplaceAll(site.Config.Font, " ", "+"),
		cssPath(page.Name),
		nav,
		body,
	)
}

func cssPath(pageName string) string {
	if strings.ToLower(pageName) == "home" {
		return "styles.css"
	}
	return "../styles.css"
}

func generateNav(items []string, activePage string) string {
	var links []string
	for _, item := range items {
		href := "/"
		if strings.ToLower(item) != "home" {
			href = "/" + strings.ToLower(item) + "/"
		}
		activeClass := ""
		if strings.ToLower(item) == strings.ToLower(activePage) {
			activeClass = " active"
		}
		links = append(links, fmt.Sprintf(`    <a href="%s" class="nav-link%s">%s</a>`, href, activeClass, capitalize(item)))
	}
	return fmt.Sprintf("<nav class=\"navbar\">\n  <div class=\"nav-inner\">\n%s\n  </div>\n</nav>", strings.Join(links, "\n"))
}

func generateComponents(comps []parser.Component) string {
	var parts []string
	for _, comp := range comps {
		parts = append(parts, components.Render(comp.Type, comp.Title, comp.Body))
	}
	return strings.Join(parts, "\n")
}

func generateCSS(config parser.Config) string {
	bgColor := "#ffffff"
	textColor := "#111827"
	navBg := "#ffffff"
	cardBg := "#f9fafb"
	borderColor := "#e5e7eb"

	if config.DarkMode {
		bgColor = "#0f172a"
		textColor = "#f1f5f9"
		navBg = "#1e293b"
		cardBg = "#1e293b"
		borderColor = "#334155"
	}

	return fmt.Sprintf(`*, *::before, *::after { box-sizing: border-box; margin: 0; padding: 0; }

:root {
  --primary: %s;
  --bg: %s;
  --text: %s;
  --nav-bg: %s;
  --card-bg: %s;
  --border: %s;
  --font: '%s', sans-serif;
}

body {
  font-family: var(--font);
  background: var(--bg);
  color: var(--text);
  min-height: 100vh;
}

.navbar {
  background: var(--nav-bg);
  border-bottom: 1px solid var(--border);
  padding: 0 2rem;
  position: sticky;
  top: 0;
  z-index: 100;
}

.nav-inner {
  max-width: 1100px;
  margin: 0 auto;
  display: flex;
  gap: 2rem;
  align-items: center;
  height: 60px;
}

.nav-link {
  text-decoration: none;
  color: var(--text);
  font-weight: 500;
  font-size: 0.95rem;
  opacity: 0.7;
  transition: opacity 0.15s;
}

.nav-link:hover, .nav-link.active {
  opacity: 1;
  color: var(--primary);
}

.main {
  max-width: 1100px;
  margin: 0 auto;
  padding: 3rem 2rem;
}

.hero {
  padding: 5rem 0 3rem;
  text-align: center;
}

.hero h1 {
  font-size: 3rem;
  font-weight: 700;
  margin-bottom: 1rem;
  color: var(--primary);
}

.hero p {
  font-size: 1.25rem;
  opacity: 0.75;
}

.card {
  background: var(--card-bg);
  border: 1px solid var(--border);
  border-radius: 12px;
  padding: 1.75rem;
  margin-bottom: 1.5rem;
}

.card h2 {
  font-size: 1.25rem;
  font-weight: 600;
  margin-bottom: 0.5rem;
}

.card p {
  opacity: 0.8;
  line-height: 1.6;
}

.grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(280px, 1fr));
  gap: 1.5rem;
  margin-bottom: 1.5rem;
}

.text-block {
  margin-bottom: 1.5rem;
  line-height: 1.7;
  opacity: 0.85;
}

.btn {
  display: inline-block;
  background: var(--primary);
  color: #fff;
  padding: 0.65rem 1.5rem;
  border-radius: 8px;
  text-decoration: none;
  font-weight: 600;
  border: none;
  cursor: pointer;
  font-size: 0.95rem;
  margin-bottom: 1.5rem;
  transition: opacity 0.15s;
}

.btn:hover { opacity: 0.85; }

.link {
  color: var(--primary);
  text-decoration: underline;
  margin-bottom: 1rem;
  display: inline-block;
}

.site-image {
  max-width: 100%%;
  border-radius: 10px;
  margin-bottom: 1.5rem;
}

footer {
  border-top: 1px solid var(--border);
  text-align: center;
  padding: 2rem;
  margin-top: 4rem;
  opacity: 0.6;
  font-size: 0.9rem;
}
`, config.Primary, bgColor, textColor, navBg, cardBg, borderColor, config.Font)
}

func capitalize(s string) string {
	if s == "" {
		return s
	}
	return strings.ToUpper(s[:1]) + s[1:]
}
