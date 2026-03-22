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
	// Clean output dir before each build
	if err := os.RemoveAll(outDir); err != nil {
		return err
	}
	if err := os.MkdirAll(outDir, 0755); err != nil {
		return err
	}

	css := generateCSS(site.Config)
	if err := os.WriteFile(filepath.Join(outDir, "styles.css"), []byte(css), 0644); err != nil {
		return err
	}

	for _, page := range site.Pages {
		html := generatePage(site, page)
		var outFile string
		if strings.ToLower(page.Name) == "home" {
			outFile = filepath.Join(outDir, "index.html")
		} else {
			outFile = filepath.Join(outDir, strings.ToLower(page.Name)+".html")
		}
		if err := os.WriteFile(outFile, []byte(html), 0644); err != nil {
			return err
		}
	}

	fmt.Printf("Built %d pages to %s/\n", len(site.Pages), outDir)
	return nil
}

func generatePage(site *parser.Site, page parser.Page) string {
	isHome := strings.ToLower(page.Name) == "home"
	nav := generateNav(site.NavItems, page.Name, isHome, site.Config)
	body := generateComponents(page.Components)

	darkClass := ""
	if site.Config.DarkMode {
		darkClass = " dark"
	}

	favicon := ""
	if site.Config.Favicon != "" {
		favicon = fmt.Sprintf("\n  <link rel=\"icon\" href=\"%s\" />", site.Config.Favicon)
	}

	return fmt.Sprintf(`<!DOCTYPE html>
<html lang="en" class="%s">
<head>
  <meta charset="UTF-8" />
  <meta name="viewport" content="width=device-width, initial-scale=1.0" />
  <title>%s - %s</title>
  <link rel="preconnect" href="https://fonts.googleapis.com" />
  <link href="https://fonts.googleapis.com/css2?family=%s:wght@400;600;700&display=swap" rel="stylesheet" />
  <link rel="stylesheet" href="%s" />%s
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
		favicon,
		nav,
		body,
	)
}

func cssPath(pageName string) string {
	return "styles.css"
}

func generateNav(items []string, activePage string, isHome bool, config parser.Config) string {
	var links []string
	for _, item := range items {
		var href string
		if strings.ToLower(item) == "home" {
			href = "index.html"
		} else {
			href = strings.ToLower(item) + ".html"
		}
		activeClass := ""
		if strings.ToLower(item) == strings.ToLower(activePage) {
			activeClass = " active"
		}
		links = append(links, fmt.Sprintf(`    <a href="%s" class="nav-link%s">%s</a>`, href, activeClass, capitalize(item)))
	}

	brand := fmt.Sprintf(`    <span class="nav-brand">%s</span>`, config.SiteName)
	if config.Logo != "" {
		brand = fmt.Sprintf(`    <a href="index.html" class="nav-logo"><img src="%s" alt="%s" /></a>`, config.Logo, config.SiteName)
	}

	return fmt.Sprintf("<nav class=\"navbar\">\n  <div class=\"nav-inner\">\n%s\n%s\n  </div>\n</nav>", brand, strings.Join(links, "\n"))
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
  --accent: %s;
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
  color: var(--accent);
}

.nav-brand {
  font-weight: 700;
  font-size: 1.1rem;
  margin-right: auto;
  color: var(--primary);
}

.nav-logo {
  margin-right: auto;
  display: flex;
  align-items: center;
}

.nav-logo img {
  height: 32px;
  width: auto;
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

.card-btn {
  margin-top: 1rem;
  margin-bottom: 0;
  font-size: 0.85rem;
  padding: 0.5rem 1.1rem;
}

.gallery {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(220px, 1fr));
  gap: 1rem;
  margin-bottom: 1.5rem;
}

.gallery-img {
  width: 100%%;
  height: 220px;
  object-fit: cover;
  border-radius: 10px;
  display: block;
}

.divider {
  border: none;
  border-top: 1px solid var(--border);
  margin: 2.5rem 0;
}

.code-block {
  background: var(--card-bg);
  border: 1px solid var(--border);
  border-left: 3px solid var(--primary);
  border-radius: 8px;
  padding: 1.25rem 1.5rem;
  margin-bottom: 1.5rem;
  overflow-x: auto;
  font-family: 'Fira Code', 'Courier New', monospace;
  font-size: 0.875rem;
  line-height: 1.7;
  white-space: pre;
}

.example-box {
  border: 1px dashed var(--border);
  border-radius: 10px;
  padding: 1.5rem;
  margin-bottom: 1.5rem;
}

.example-box .hero {
  padding: 2rem 0 1rem;
}

.comp-label {
  font-size: 0.75rem;
  font-weight: 600;
  letter-spacing: 0.08em;
  text-transform: uppercase;
  opacity: 0.45;
  margin-bottom: 0.5rem;
}

.heading {
  margin-bottom: 1rem;
}

.heading-title {
  font-size: 1.75rem;
  font-weight: 700;
  color: var(--primary);
}

.heading-sub {
  font-size: 0.8rem;
  font-weight: 600;
  letter-spacing: 0.08em;
  text-transform: uppercase;
  opacity: 0.45;
  margin-bottom: 0.25rem;
}
`, config.Primary, config.Accent, bgColor, textColor, navBg, cardBg, borderColor, config.Font)
}

func capitalize(s string) string {
	if s == "" {
		return s
	}
	return strings.ToUpper(s[:1]) + s[1:]
}
