from fasthtml.common import *
from components.layout import navbar, hero, features, footer
from components.icons import WAND_ICON
from components.errors import not_found_page
from components.theme import LIGHT_THEME, DARK_THEME

theme_toggle_js = ScriptX(
    "static/js/themeToggle.js", 
    lightTheme=LIGHT_THEME,
    darkTheme=DARK_THEME
)
tailwind_css = Link(rel="stylesheet", href="/static/css/output.css", type="text/css")

app, rt = fast_app(
    pico=False,
    surreal=False,
    live=True,
    hdrs=(theme_toggle_js, tailwind_css),
    htmlkw=dict(data_theme=LIGHT_THEME)
)

@rt("/")
def get():
    return Main(
        navbar(),
        hero(),
        features(),
        footer(),
        cls="min-h-screen bg-base-100"
    )
    
@rt("/{path:path}")  # Catch-all route that matches any path
def not_found(path: str):
    return not_found_page()


if __name__ == "__main__":
    serve(reload_includes=["static/*"])