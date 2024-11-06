from dataclasses import dataclass
from typing import Optional
from fasthtml.common import *
from .theme import theme_toggle
from .icons import DISCORD_ICON, WAND_ICON

NAV_LINK = "px-4 py-2 text-content-light-secondary dark:text-content-dark-secondary hover:text-content-light-primary dark:hover:text-content-dark-primary transition-colors"
CONTAINER = "max-w-7xl mx-auto px-4"

@dataclass
class NavItem:
    text: str
    href: str
    external: bool = False

NAV_ITEMS = [
    NavItem("Docs", "/docs"),
    NavItem("Components", "/components"),
    NavItem("Examples", "/examples"),
]

def nav_link(item: NavItem) -> Li:
    """Generate a navigation link with consistent styling"""
    return Li(
        A(item.text, 
          href=item.href, 
          cls=NAV_LINK,
          **({"target": "_blank"} if item.external else {}))
    )

def navbar() -> Div:
    """Main navigation bar component"""
    return Div(
        Div(
            # Brand
            Div(
                A("FastWAND", WAND_ICON(size="h-5 w-5", color="currentColor"),
                  cls="inline-flex items-center px-4 py-2 text-xl font-semibold text-content-light-primary dark:text-content-dark-primary hover:bg-surface-light dark:hover:bg-surface-dark rounded-lg transition-colors justify-normal"),
                cls="flex items-center"
            ),
            # Nav links
            Div(
                Ul(*map(nav_link, NAV_ITEMS), 
                   cls="hidden lg:flex items-center space-x-2"),
                cls="flex-1 justify-center"
            ),
            # Actions
            Div(
                theme_toggle(),
                A(DISCORD_ICON(), 
                  href="https://discord.com/channels/689892369998676007/1247700012952191049",
                  cls="inline-flex items-center p-2 text-content-light-secondary dark:text-content-dark-secondary hover:bg-surface-light dark:hover:bg-surface-dark rounded-lg transition-colors",
                  target="_blank"),
                A("GitHub ↗",
                  href="https://github.com/AnswerDotAI/fasthtml",
                  cls="inline-flex items-center px-4 py-2 text-content-light-secondary dark:text-content-dark-secondary hover:bg-surface-light dark:hover:bg-surface-dark rounded-lg transition-colors",
                  target="_blank"),
                cls="flex items-center gap-2"
            ),
            cls=f"{CONTAINER} h-16 flex items-center justify-between"
        ),
        cls="border-b border-border-light dark:border-border-dark bg-surface-light dark:bg-surface-dark"
    )

def hero() -> Div:
    """Hero section with main messaging"""
    return Div(
        Div(
            # Logo section
            Div(
                Img(src="/static/logo.svg", 
                    alt="FastHTML Logo",
                    cls="w-80 md:w-[28rem] lg:w-[32rem] mb-8 lg:mb-0 transition-transform hover:scale-105"),
                cls="lg:order-2 flex justify-center"
            ),
            # Content section
            Div(
                H1("Build Modern Web Apps",
                   cls="text-4xl sm:text-5xl lg:text-6xl font-bold text-content-light-primary dark:text-content-dark-primary leading-tight"),
                P("FastHTML combines the simplicity of server-side rendering with the power of modern web features.", 
                  cls="py-6 text-xl text-content-light-secondary dark:text-content-dark-secondary"),
                Div(
                    A("Get Started →", 
                      href="/docs", 
                      cls="inline-flex items-center px-6 py-3 text-lg rounded-lg bg-blue-600 text-white hover:bg-blue-700 transition-colors"),
                    A("View Examples", 
                      href="/examples", 
                      cls="inline-flex items-center px-6 py-3 text-lg rounded-lg border border-border-light dark:border-border-dark text-content-light-secondary dark:text-content-dark-secondary hover:bg-surface-light dark:hover:bg-surface-dark transition-colors"),
                    cls="flex gap-4"
                ),
                cls="max-w-xl lg:order-1 mx-auto lg:mx-0"
            ),
            cls=f"{CONTAINER} py-16 lg:py-24 flex flex-col lg:flex-row items-center justify-between gap-12"
        ),
        cls="min-h-[70vh] bg-base-light dark:bg-base-dark flex items-center"
    )

@dataclass
class Feature:
    title: str
    description: str
    icon: Optional[Div] = None

def feature_card(feature: Feature) -> Div:
    """Feature card component"""
    return Div(
        Div(
            *([] if not feature.icon else [Div(feature.icon, cls="text-blue-600 w-12 h-12")]),
            H3(feature.title, 
               cls="text-xl font-semibold mb-2 text-content-light-primary dark:text-content-dark-primary"),
            P(feature.description, 
              cls="text-content-light-secondary dark:text-content-dark-secondary"),
            cls="p-6"
        ),
        cls="bg-surface-light dark:bg-surface-dark rounded-lg shadow-lg hover:shadow-xl transition-shadow"
    )

def features() -> Div:
    """Features section"""
    FEATURES = [
        Feature("Pure Python Magic", 
               "Build any web app with just Python. Import any Python or JS library you need."),
        Feature("Unlimited Potential",
               "Full access to HTTP, HTML, JS, and CSS fundamentals. No abstractions in your way - build anything you can imagine."),
        Feature("Production Ready",
               "Fast, scalable, and easy to deploy. Run anywhere Python runs, with built-in dev features like hot reload."),
    ]
    
    return Div(
        Div(
            H2("Why FastWAND?", 
               cls="text-3xl font-bold text-center mb-8 text-content-light-primary dark:text-content-dark-primary"),
            Div(*map(feature_card, FEATURES),
                cls="grid grid-cols-1 md:grid-cols-3 gap-8 max-w-7xl mx-auto px-4"),
            cls="py-16"
        ),
        cls="bg-base-light dark:bg-base-dark"
    )

def footer() -> Div:
    """Footer component"""
    return Footer(        
        Div(
            P("© 2024 FastWAND. Built with Python and magic. ✨",
              cls="text-content-light-secondary dark:text-content-dark-secondary"),
            cls="text-center py-6"
        ),
        cls="bg-surface-light dark:bg-surface-dark border-t border-border-light dark:border-border-dark"
    )