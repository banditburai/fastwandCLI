from fasthtml.common import *
from components.layout import navbar, hero, features, footer
from components.icons import WAND_ICON
from components.errors import not_found_page

theme_toggle_js = Script(src="/static/js/themeToggle.js")
tailwind_css = Link(rel="stylesheet", href="/static/css/output.css", type="text/css")

app, rt = fast_app(
    pico=False,
    surreal=False,
    live=True,
    hdrs=(theme_toggle_js, tailwind_css),
    htmlkw=dict(cls="bg-surface-light dark:bg-surface-dark")    
)

@rt("/")
def get():
    return Main(
        navbar(),
        hero(),
        features(),
        footer(),
        cls=f"min-h-screen"
    )
    
@rt("/{path:path}")
def not_found(path: str):
    return not_found_page()

if __name__ == "__main__":
    serve(reload_includes=["static/*"])