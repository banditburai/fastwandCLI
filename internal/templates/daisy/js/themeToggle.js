// Prevent FOUC (Flash of Unstyled Content)
(function() {
    const lightTheme = '{lightTheme}';
    const darkTheme = '{darkTheme}';
    const savedTheme = localStorage.getItem('theme');
    const prefersDark = window.matchMedia('(prefers-color-scheme: dark)').matches;
    const theme = savedTheme || (prefersDark ? darkTheme : lightTheme);
    document.documentElement.setAttribute('data-theme', theme);
})();

// Theme toggle functionality
function updateTheme(isDark) {
    const lightTheme = '{lightTheme}';
    const darkTheme = '{darkTheme}';
    const html = document.documentElement;
    html.setAttribute('data-theme', isDark ? darkTheme : lightTheme);
    localStorage.setItem('theme', isDark ? darkTheme : lightTheme);
    
    // Only try to update checkbox if it exists
    const themeToggle = document.getElementById('theme-toggle');
    if (themeToggle) {
        themeToggle.checked = isDark;
    }
}

// Initialize theme and set up listeners when DOM is ready
document.addEventListener('DOMContentLoaded', function() {
    const lightTheme = '{lightTheme}';
    const darkTheme = '{darkTheme}';
    const themeToggle = document.getElementById('theme-toggle');
    
    // Initialize theme based on localStorage or system preference
    const savedTheme = localStorage.getItem('theme');
    const prefersDark = window.matchMedia('(prefers-color-scheme: dark)').matches;
    const theme = savedTheme || (prefersDark ? darkTheme : lightTheme);
    updateTheme(theme === darkTheme);
    
    // Listen for toggle changes
    if (themeToggle) {
        themeToggle.addEventListener('change', function() {
            updateTheme(this.checked);
        });
    }
});
