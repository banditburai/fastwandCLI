package templates

import _ "embed" // for single file embedding

//go:embed shared/icons.py
var sharedIcons []byte

//go:embed shared/logo.svg
var sharedLogo []byte

//go:embed vanilla/css/input.css
var vanillaInputCSS []byte

//go:embed vanilla/main.py
var vanillaMainPY []byte

//go:embed vanilla/tailwind.config.js
var vanillaTailwindConfig []byte

//go:embed vanilla/components/layout.py
var vanillaLayout []byte

//go:embed vanilla/components/theme.py
var vanillaTheme []byte

//go:embed vanilla/components/errors.py
var vanillaErrors []byte

//go:embed vanilla/js/themeToggle.js
var vanillaThemeToggle []byte

//go:embed daisy/css/input.css
var daisyInputCSS []byte

//go:embed daisy/main.py
var daisyMainPY []byte

//go:embed daisy/tailwind.config.js
var daisyTailwindConfig []byte

//go:embed daisy/components/layout.py
var daisyLayout []byte

//go:embed daisy/components/theme.py
var daisyTheme []byte

//go:embed daisy/components/errors.py
var daisyErrors []byte

//go:embed daisy/js/themeToggle.js
var daisyThemeToggle []byte

var templateFiles = map[string][]TemplateFiles{
	"shared": {
		{Path: "components/icons.py", Content: sharedIcons},
		{Path: "static/logo.svg", Content: sharedLogo},
	},
	"vanilla": {
		{Path: "static/css/input.css", Content: vanillaInputCSS},
		{Path: "main.py", Content: vanillaMainPY},
		{Path: "tailwind.config.js", Content: vanillaTailwindConfig},
		{Path: "components/layout.py", Content: vanillaLayout},
		{Path: "components/theme.py", Content: vanillaTheme},
		{Path: "components/errors.py", Content: vanillaErrors},
		{Path: "static/js/themeToggle.js", Content: vanillaThemeToggle},
	},
	"daisy": {
		{Path: "static/css/input.css", Content: daisyInputCSS},
		{Path: "main.py", Content: daisyMainPY},
		{Path: "tailwind.config.js", Content: daisyTailwindConfig},
		{Path: "components/layout.py", Content: daisyLayout},
		{Path: "components/theme.py", Content: daisyTheme},
		{Path: "components/errors.py", Content: daisyErrors},
		{Path: "static/js/themeToggle.js", Content: daisyThemeToggle},
	},
}

func getFilesFromDir(dir string) ([]TemplateFiles, error) {
	return templateFiles[dir], nil
}
