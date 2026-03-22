package components

import (
	"fmt"
	"strings"
)

func Render(compType, title, body string) string {
	switch compType {
	case "hero":
		return renderHero(title, body)
	case "card":
		return renderCard(title, body)
	case "button":
		return renderButton(title, body)
	case "image":
		return renderImage(title, body)
	case "text":
		return renderText(title, body)
	case "grid":
		return renderGrid(title, body)
	case "gallery":
		return renderGallery(title, body)
	case "divider":
		return renderDivider()
	case "footer":
		return renderFooter(title, body)
	case "link":
		return renderLink(title, body)
	case "code":
		return renderCode(title, body)
	case "example":
		return renderExample(title, body)
	case "heading":
		return renderHeading(title, body)
	default:
		return fmt.Sprintf("<div><!-- unknown component: %s --></div>", compType)
	}
}

func renderHero(title, body string) string {
	subtitle := ""
	if body != "" {
		subtitle = fmt.Sprintf("\n  <p>%s</p>", body)
	}
	return fmt.Sprintf(`<section class="hero">
  <h1>%s</h1>%s
</section>`, title, subtitle)
}

func renderCard(title, body string) string {
	bodyHTML := ""
	if body != "" {
		bodyHTML = fmt.Sprintf("\n  <p>%s</p>", body)
	}
	return fmt.Sprintf(`<div class="card">
  <h2>%s</h2>%s
</div>`, title, bodyHTML)
}

func renderButton(title, body string) string {
	href := "#"
	if body != "" {
		href = body
	}
	return fmt.Sprintf(`<a href="%s" class="btn">%s</a>`, href, title)
}

func renderImage(title, body string) string {
	alt := body
	if alt == "" {
		alt = title
	}
	return fmt.Sprintf(`<img src="%s" alt="%s" class="site-image" />`, title, alt)
}

func renderText(title, body string) string {
	content := title
	if body != "" {
		content = title + " " + body
	}
	return fmt.Sprintf(`<p class="text-block">%s</p>`, content)
}

// Grid syntax per line: Title | Body | Button Label > URL
// Button segment is optional.
func renderGrid(title, body string) string {
	var cards []string
	lines := strings.Split(body, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		parts := strings.SplitN(line, "|", 3)
		cardTitle := strings.TrimSpace(parts[0])

		cardBody := ""
		if len(parts) >= 2 {
			cardBody = strings.TrimSpace(parts[1])
		}

		buttonHTML := ""
		if len(parts) == 3 {
			btn := strings.TrimSpace(parts[2])
			btnParts := strings.SplitN(btn, ">", 2)
			btnLabel := strings.TrimSpace(btnParts[0])
			btnHref := "#"
			if len(btnParts) == 2 {
				btnHref = strings.TrimSpace(btnParts[1])
			}
			buttonHTML = fmt.Sprintf("\n  <a href=\"%s\" class=\"btn card-btn\">%s</a>", btnHref, btnLabel)
		}

		bodyHTML := ""
		if cardBody != "" {
			bodyHTML = fmt.Sprintf("\n  <p>%s</p>", cardBody)
		}

		cards = append(cards, fmt.Sprintf(`  <div class="card">
  <h2>%s</h2>%s%s
</div>`, cardTitle, bodyHTML, buttonHTML))
	}
	return fmt.Sprintf("<div class=\"grid\">\n%s\n</div>", strings.Join(cards, "\n"))
}

// Gallery syntax per line: image-src | alt text
func renderGallery(title, body string) string {
	var images []string
	lines := strings.Split(body, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		parts := strings.SplitN(line, "|", 2)
		src := strings.TrimSpace(parts[0])
		alt := src
		if len(parts) == 2 {
			alt = strings.TrimSpace(parts[1])
		}
		images = append(images, fmt.Sprintf(`  <img src="%s" alt="%s" class="gallery-img" />`, src, alt))
	}
	return fmt.Sprintf("<div class=\"gallery\">\n%s\n</div>", strings.Join(images, "\n"))
}

func renderDivider() string {
	return `<hr class="divider" />`
}

func renderFooter(title, body string) string {
	content := title
	if body != "" {
		content += " " + body
	}
	return fmt.Sprintf(`<footer>%s</footer>`, content)
}

func renderLink(title, body string) string {
	href := "#"
	if body != "" {
		href = body
	}
	return fmt.Sprintf(`<a href="%s" class="link">%s</a>`, href, title)
}

// example wraps its content lines as rendered HTML inside a dashed preview box.
// Each line is rendered as a raw HTML snippet — use it to show live component previews.
func renderExample(title, body string) string {
	inner := title
	if body != "" {
		inner = title + "\n" + body
	}
	return fmt.Sprintf(`<div class="example-box">%s</div>`, inner)
}

func renderHeading(title, body string) string {
	sub := ""
	if body != "" {
		sub = fmt.Sprintf(`<p class="heading-sub">%s</p>`, body)
	}
	return fmt.Sprintf(`<div class="heading"><h2 class="heading-title">%s</h2>%s</div>`, title, sub)
}

func renderCode(title, body string) string {
	content := title
	if body != "" {
		content = title + "\n" + body
	}
	escaped := strings.NewReplacer(
		"&", "&amp;",
		"<", "&lt;",
		">", "&gt;",
	).Replace(content)
	return fmt.Sprintf(`<pre class="code-block"><code>%s</code></pre>`, escaped)
}
