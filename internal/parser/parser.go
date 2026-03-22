package parser

import (
	"strings"
)

type Config struct {
	Primary  string
	Font     string
	DarkMode bool
	SiteName string
}

type Component struct {
	Type    string
	Title   string
	Body    string
}

type Page struct {
	Name       string
	Components []Component
}

type Site struct {
	Config  Config
	NavItems []string
	Pages   []Page
}

func Parse(content string) (*Site, error) {
	site := &Site{}
	site.Config = Config{
		Primary:  "#4f46e5",
		Font:     "Inter",
		DarkMode: false,
		SiteName: "My Site",
	}

	// Split into blocks by blank lines
	blocks := splitBlocks(content)

	var currentPage *Page
	var i int

	for i < len(blocks) {
		block := blocks[i]
		lines := splitLines(block)
		if len(lines) == 0 {
			i++
			continue
		}

		keyword := strings.ToLower(strings.TrimSpace(lines[0]))

		switch keyword {
		case "config":
			parseConfig(lines[1:], &site.Config)
		case "header":
			if len(lines) > 1 {
				site.NavItems = strings.Fields(lines[1])
			}
		default:
			// Check if this is a page name
			if isPageName(keyword, site.NavItems) {
				page := Page{Name: keyword}
				currentPage = &page
				site.Pages = append(site.Pages, page)
				// update reference
				i++
				// Gather components until next page or end
				for i < len(blocks) {
					b := blocks[i]
					blines := splitLines(b)
					if len(blines) == 0 {
						i++
						continue
					}
					bkeyword := strings.ToLower(strings.TrimSpace(blines[0]))
					if isPageName(bkeyword, site.NavItems) {
						break
					}
					comp := parseComponent(blines)
					if comp != nil {
						site.Pages[len(site.Pages)-1].Components = append(site.Pages[len(site.Pages)-1].Components, *comp)
					}
					i++
				}
				_ = currentPage
				continue
			}
		}
		i++
	}

	return site, nil
}

func parseConfig(lines []string, config *Config) {
	for _, line := range lines {
		parts := strings.SplitN(strings.TrimSpace(line), " ", 2)
		if len(parts) < 2 {
			continue
		}
		key := strings.ToLower(parts[0])
		val := strings.TrimSpace(parts[1])
		switch key {
		case "primary":
			config.Primary = val
		case "font":
			config.Font = val
		case "dark-mode":
			config.DarkMode = val == "true"
		case "site-name":
			config.SiteName = val
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

func splitBlocks(content string) []string {
	// Normalize line endings
	content = strings.ReplaceAll(content, "\r\n", "\n")
	raw := strings.Split(content, "\n\n")
	var blocks []string
	for _, b := range raw {
		b = strings.TrimSpace(b)
		if b != "" {
			blocks = append(blocks, b)
		}
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
