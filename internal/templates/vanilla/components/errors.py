from fasthtml.common import *
from .icons import WAND_ICON

CAT_REGULAR = r"""
      /\___/\    *
     (  o o  ) ⚡
     (  =^=  )━━━╋
      (--m--) ⚡ *
       |===|
"""   

CAT_REVERSED = r"""
  *    /\___/\
    ⚡ (  o o  )
  ╋━━━(  =^=  )
   * ⚡ (--m--) 
        |===|
"""

def create_cat_pattern(num_cats: int = 8) -> Div:
    """Creates a repeating pattern of cats"""
    return Div(
        *[Pre(
            CAT_REGULAR,            
            cls="font-mono text-xl text-error-accent-light dark:text-error-accent-dark opacity-70 whitespace-pre animate-sparkle",
            style=f"animation-delay: {i * 0.2}s"
        ) for i in range(num_cats)],
        cls="hidden lg:block overflow-hidden"
    )

def create_error_header() -> Div:
    return Div(        
        WAND_ICON(size="h-[175px] w-[170px]", 
                 cls="text-error-accent-light dark:text-error-accent-dark"),
        H1("404", 
           cls="text-[8rem] font-bold text-error-accent-light dark:text-error-accent-dark mx-8"),
        cls="flex items-center justify-center mb-2"
    )

def create_error_message() -> Div:
    return Div(
        Script(popover_script),
        Style(popover_style),
        H2("Oops! Page not found", 
           cls="text-4xl font-bold mb-8 text-content-light-primary dark:text-content-dark-primary"),
        Div(
            P("Looks like our wand misfired!", 
              cls="text-xl text-content-light-secondary dark:text-content-dark-secondary"),
            P(
                "The page you're looking for might have apparated ",
                Button(
                    "elsewhere",
                    cls="text-xl text-content-light-secondary dark:text-content-dark-secondary underline cursor-pointer hover:text-error-accent-light dark:hover:text-error-accent-dark transition-colors",
                    popovertarget="cat-wizard",
                    id="elsewhere-btn"
                ),
                ".",
                cls="text-xl text-content-light-secondary dark:text-content-dark-secondary mb-8"
            ),
            Div(
                Pre(CAT_REVERSED, 
                    cls="font-mono text-sm text-error-accent-light dark:text-error-accent-dark whitespace-pre cat-desktop animate-wiggle"),
                Pre(CAT_REGULAR, 
                    cls="font-mono text-sm text-error-accent-light dark:text-error-accent-dark whitespace-pre cat-mobile animate-wiggle"),
                id="cat-wizard",
                popover="auto",
                cls="bg-transparent border-none shadow-none"
            ),
            cls="space-y-2"
        ),
        cls="mb-8 text-left max-w-2xl mx-auto px-8"
    )

def create_back_button() -> Div:
    return Div(
        Button(
            "← Go Back",
            onclick="history.back()",
            cls="px-6 py-3 text-lg font-semibold rounded-lg bg-error-button-light dark:bg-error-button-dark hover:bg-error-button-hover-light dark:hover:bg-error-button-hover-dark text-white transition-colors"
        ),
        cls="flex justify-center gap-4 mb-12"
    )

def not_found_page() -> Main:
    return Main(
        Style("""
            html, body {
                overscroll-behavior: none;
                overflow: hidden;
                touch-action: pan-y;
                height: 100%;
                width: 100%;
            }
        """),
        Div(
            Div(
                create_cat_pattern(),
                Div(
                    create_error_header(),
                    create_error_message(),
                    create_back_button(),
                    cls="flex-1 flex flex-col justify-center"
                ),
                cls="flex w-full h-screen"
            ),            
            cls="h-screen bg-error-base-light dark:bg-error-base-dark"
        ),        
        cls=""
    )

popover_script = """
document.addEventListener('DOMContentLoaded', () => {
    const openPopovers = new WeakMap();
    
    function updatePosition(popover, invoker) {
        const buttonRect = invoker.getBoundingClientRect();
        const popoverRect = popover.getBoundingClientRect();
        const isMobile = window.innerWidth < 768;
        
        const x = buttonRect.left + (buttonRect.width / 2) - (popoverRect.width / 2);
        const y = isMobile ? buttonRect.bottom + 20 : buttonRect.top - popoverRect.height - 20;
        
        Object.assign(popover.style, {
            position: 'absolute',
            margin: '0',
            left: `${x}px`,
            top: `${y}px`
        });
        
        popover.classList.add('is-positioned');
    }
    
    function positionPopover(e) {
        const popover = e.target;
        const invoker = document.querySelector(
            `[popovertarget="${popover.getAttribute('id')}"]`
        );
        
        if (e.newState === 'open') {
            const intervalId = setInterval(() => updatePosition(popover, invoker), 100);
            openPopovers.set(popover, intervalId);
            invoker.classList.add('is-popover-open');
        } else {
            const intervalId = openPopovers.get(popover);
            if (intervalId) {
                clearInterval(intervalId);
                openPopovers.delete(popover);
                popover.classList.remove('is-positioned');
            }
            invoker.classList.remove('is-popover-open');
        }
    }
    
    document.querySelectorAll('[popover]')
        .forEach(p => p.addEventListener('toggle', positionPopover));
    
    // Auto-open popover
    const popover = document.getElementById('cat-wizard');
    popover?.showPopover();
});
"""

popover_style = """
[popover] {
    width: max-content;
    margin: 0;
    padding: 0;
    position: absolute;
    top: 0;
    left: 0;
    visibility: hidden;
}

[popover].is-positioned {
    visibility: visible;
    animation: wiggle 2s ease-in-out infinite;
}

.cat-desktop { display: none; }
.cat-mobile { display: block; }

@media (min-width: 768px) {
    .cat-desktop { display: block; }
    .cat-mobile { display: none; }
}
"""