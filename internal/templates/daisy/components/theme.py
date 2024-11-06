from fasthtml.common import *
from fasthtml.svg import *
from .icons import SUN_ICON, MOON_ICON

# Theme configuration
LIGHT_THEME = "retro"
DARK_THEME = "dim"

def theme_toggle():
    theme_checkbox = Input(
        type='checkbox',
        value='night',
        cls='theme-controller',
        id='theme-toggle'      
    )  

    return Label(
        theme_checkbox,       
        Span(SUN_ICON, cls="swap-off fill-current"),
        Span(MOON_ICON, cls="swap-on fill-current"),           
        cls='btn btn-ghost swap'
    ) 