package parser

import (
	"fmt"
	"strings"
)

var knownComponents = map[string]bool{
	"hero": true, "card": true, "button": true, "image": true,
	"text": true, "grid": true, "gallery": true, "divider": true,
	"footer": true, "link": true, "code": true, "example": true, "heading": true,
}

var knownConfigKeys = map[string]bool{
	"primary": true, "accent": true, "font": true, "dark-mode": true,
	"site-name": true, "logo": true, "favicon": true,
}

type Config struct {
	Primary  string
	Accent   string
	Font     string
	DarkMode bool
	SiteName string
	Logo     string
	Favicon  string
}

type Component struct {
	Type  string
	Title string
	Body  string
}

type Page struct {
	Name       string
	Components []Component
}

type Site struct {
	Config   Config
	NavItems []string
	Pages    []Page
	Warnings []string
}

func Parse(content string) (*Site, error) {
	site := &Site{}
	site.Config = Config{
		Primary:  "#4f46e5",
		Accent:   "#7c3aed",
		Font:     "Inter",
		DarkMode: false,
		SiteName: "My Site",
		Logo:     "",
		Favicon:  "",
	}

	blocks := splitBlocksWithLines(content)

	if len(blocks) == 0 {
		return nil, fmt.Errorf("file is empty")
	}

	// Validate config comes first
	if len(blocks) > 0 {
		firstLines := splitLines(blocks[0].text)
		if len(firstLines) > 0 && strings.ToLower(firstLines[0]) != "config" {
			site.Warnings = append(site.Warnings, fmt.Sprintf("line %d: expected 'config' block at the top of the file", blocks[0].line))
		}
	}

	var i int
	for i < len(blocks) {
		b := blocks[i]
		lines := splitLines(b.text)
		if len(lines) == 0 {
			i++
			continue
		}

		keyword := strings.ToLower(strings.TrimSpace(lines[0]))

		switch keyword {
		case "config":
			parseConfig(lines[1:], &site.Config, b.line, site)
		case "header":
			if len(lines) < 2 {
				site.Warnings = append(site.Warnings, fmt.Sprintf("line %d: 'header' block has no nav items", b.line))
			} else {
				site.NavItems = strings.Fields(lines[1])
			}
		default:
			if isPageName(keyword, site.NavItems) {
				site.Pages = append(site.Pages, Page{Name: keyword})
				i++
				for i < len(blocks) {
					cb := blocks[i]
					clines := splitLines(cb.text)
					if len(clines) == 0 {
						i++
						continue
					}
					ckeyword := strings.ToLower(strings.TrimSpace(clines[0]))
					if isPageName(ckeyword, site.NavItems) {
						break
					}
					comp := parseComponent(splitLines(cb.text))
					if comp != nil {
						if !knownComponents[comp.Type] {
							site.Warnings = append(site.Warnings, fmt.Sprintf("line %d: unknown component '%s'", cb.line, comp.Type))
						}
						site.Pages[len(site.Pages)-1].Components = append(site.Pages[len(site.Pages)-1].Components, *comp)
					}
					i++
				}
				continue
			} else {
				// Unknown top-level keyword
				if !knownComponents[keyword] && keyword != "config" && keyword != "header" {
					site.Warnings = append(site.Warnings, fmt.Sprintf("line %d: '%s' is not a page defined in the header and not a known keyword — skipping", b.line, keyword))
				}
			}
		}
		i++
	}

	// Warn about nav items with no page section
	for _, item := range site.NavItems {
		found := false
		for _, page := range site.Pages {
			if strings.ToLower(page.Name) == strings.ToLower(item) {
				found = true
				break
			}
		}
		if !found {
			site.Warnings = append(site.Warnings, fmt.Sprintf("warning: nav item '%s' has no matching page section in this file", item))
		}
	}

	return site, nil
}

func parseConfig(lines []string, config *Config, startLine int, site *Site) {
	for _, line := range lines {
		parts := strings.SplitN(strings.TrimSpace(line), " ", 2)
		if len(parts) < 2 {
			if strings.TrimSpace(parts[0]) != "" && !knownConfigKeys[strings.ToLower(parts[0])] {
				site.Warnings = append(site.Warnings, fmt.Sprintf("config: unknown key '%s' (missing value?)", parts[0]))
			}
			continue
		}
		key := strings.ToLower(parts[0])
		val := strings.TrimSpace(parts[1])
		switch key {
		case "primary":
			config.Primary = val
		case "accent":
			config.Accent = val
		case "font":
			config.Font = val
		case "dark-mode":
			config.DarkMode = val == "true"
		case "site-name":
			config.SiteName = val
		case "logo":
			config.Logo = val
		case "favicon":
			config.Favicon = val
		default:
			site.Warnings = append(site.Warnings, fmt.Sprintf("config: unknown key '%s' — did you mean one of: primary, accent, font, dark-mode, site-name, logo, favicon?", key))
		}
	}
}

func parseComponent(lines []string) *Component {
	if len(lines) == 0 {
		return nil
	}
	comp := &Component{
		Type: strings.ToLower(strings.TrimSpace(lines[0])),
	}
	if len(lines) > 1 {
		comp.Title = strings.TrimSpace(lines[1])
	}
	if len(lines) > 2 {
		comp.Body = strings.Join(lines[2:], "\n")
	}
	return comp
}

func isPageName(name string, navItems []string) bool {
	for _, item := range navItems {
		if strings.ToLower(item) == name {
			return true
		}
	}
	return false
}

type block struct {
	text string
	line int
}

func splitBlocksWithLines(content string) []block {
	content = strings.ReplaceAll(content, "\r\n", "\n")
	raw := strings.Split(content, "\n\n")
	var blocks []block
	lineNum := 1
	for _, b := range raw {
		trimmed := strings.TrimSpace(b)
		if trimmed != "" {
			blocks = append(blocks, block{text: trimmed, line: lineNum})
		}
		lineNum += strings.Count(b, "\n") + 2
	}
	return blocks
}

func splitLines(block string) []string {
	lines := strings.Split(block, "\n")
	var result []string
	for _, l := range lines {
		trimmed := strings.TrimSpace(l)
		if trimmed != "" {
			result = append(result, trimmed)
		}
	}
	return result
}
