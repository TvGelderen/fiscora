let darkMode = $state(false);

export function createDarkMode() {
    return {
        get darkMode() {
            return darkMode;
        },
        set: (value: boolean) => (darkMode = value),
        toggle: () => (darkMode = !darkMode),
    };
}
