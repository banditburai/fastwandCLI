from dataclasses import dataclass
from typing import Optional
from fasthtml.common import *
from .theme import theme_toggle
from .icons import DISCORD_ICON, WAND_ICON

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
    return Li(A(item.text, 
                href=item.href,
                **({"target": "_blank"} if item.external else {})))

def navbar() -> Div:
    """DaisyUI version of the navbar"""
    return Div(
        Div(
            # Navbar start - Logo
            Div(
                A("Fast", "WAND",
                  WAND_ICON(size="h-5 w-5", color="currentColor"),
                  cls="btn btn-ghost text-xl"),
                cls="navbar-start"
            ),
            # Navbar center - Navigation
            Div(
                Ul(*map(nav_link, NAV_ITEMS),
                   cls="menu menu-horizontal px-1"),
                cls="navbar-center hidden lg:flex"
            ),
            # Navbar end - Theme toggle and external links
            Div(
                theme_toggle(),
                A(DISCORD_ICON(),
                  href="https://discord.com/channels/689892369998676007/1247700012952191049",
                  cls="btn btn-ghost",
                  target="_blank"),
                A("GitHub ↗",
                  href="https://github.com/AnswerDotAI/fasthtml",
                  cls="btn btn-ghost",
                  target="_blank"),
                cls="navbar-end gap-2"
            ),
            cls="navbar bg-base-100 max-w-7xl mx-auto px-4"
        ),
        cls="border-b border-base-200"
    )

def hero() -> Div:
    """DaisyUI version of the hero section"""
    return Div(
        Div(
            # Logo section
            Div(
                Img(src="/static/logo.svg",
                    alt="FastHTML Logo",
                    cls="w-80 md:w-[28rem] lg:w-[32rem] mb-8 lg:mb-0"),
                cls="lg:order-2 flex justify-center"
            ),
            # Content section
            Div(
                H1("Build Modern Web Apps", cls="text-5xl lg:text-6xl font-bold"),
                P("FastHTML combines the simplicity of server-side rendering with the power of modern web features.",
                  cls="py-6 text-xl text-base-content/80"),
                Div(
                    A("Get Started →", href="/docs", cls="btn btn-primary btn-lg"),
                    A("View Examples", href="/examples", cls="btn btn-ghost btn-lg"),
                    cls="flex gap-4"
                ),
                cls="max-w-xl lg:order-1"
            ),
            cls="hero-content flex-col lg:flex-row justify-between gap-8 py-8 px-6 max-w-7xl mx-auto"
        ),
        cls="hero min-h-[85vh] bg-base-200"
    )

@dataclass
class Feature:
    title: str
    description: str
    icon: Optional[Div] = None

def feature_card(feature: Feature) -> Div:
    """DaisyUI version of feature cards"""
    return Div(
        Div(
            *([] if not feature.icon else [Div(feature.icon, cls="text-primary w-12 h-12")]),
            H3(feature.title, cls="card-title"),
            P(feature.description, cls="text-base-content/70"),
            cls="card-body"
        ),
        cls="card bg-base-100 shadow-xl"
    )

def features() -> Div:
    """Features section using DaisyUI cards"""
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
            H2("Why FastWAND?", cls="text-3xl font-bold text-center mb-8"),
            Div(*map(feature_card, FEATURES),
                cls="grid grid-cols-1 md:grid-cols-3 gap-8 max-w-7xl mx-auto px-4"),
            cls="pb-16"
        ),
        cls="bg-base-200"
    )

def footer() -> Footer:
    """DaisyUI version of the footer"""
    return Footer(
        Div(
            P("© 2024 FastWAND. Built with Python and magic. ✨",
              cls="text-center text-base-content/70 py-4 border-t border-base-300"),
            cls="max-w-7xl mx-auto px-4"
        ),
        cls="bg-base-100"
    )
        
