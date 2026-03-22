package components

import "fmt"

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
	case "footer":
		return renderFooter(title, body)
	case "link":
		return renderLink(title, body)
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
	alt := title
	if alt == "" {
		alt = "image"
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

func renderGrid(title, body string) string {
	// Grid wraps its children — for MVP just open a grid div with title as comment
	_ = body
	return fmt.Sprintf(`<div class="grid">
  <!-- grid: %s -->
</div>`, title)
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
