from fasthtml.common import *
from fasthtml.svg import *
from .icons import SUN_ICON, MOON_ICON

def theme_toggle():
    theme_checkbox = Input(
        type='checkbox',
        id='theme-toggle',
        cls='hidden'
    )  

    return Label(
        theme_checkbox,       
        Span(
            SUN_ICON(), 
            cls="absolute fill-gray-600 dark:fill-gray-300 p-2 inset-0 flex items-center justify-center transition-all duration-300 transform opacity-100 dark:opacity-0 dark:rotate-90 dark:scale-0"            
        ),
        Span(
            MOON_ICON(), 
            cls="absolute fill-gray-600 dark:fill-gray-300 p-2 inset-0 flex items-center justify-center transition-all duration-300 transform opacity-0 rotate-90 scale-0 dark:opacity-100 dark:rotate-0 dark:scale-100"            
        ),           
        cls="relative inline-flex h-10 w-10 items-center justify-center rounded-lg hover:bg-gray-100 dark:hover:bg-gray-700 cursor-pointer transition-colors"
    )