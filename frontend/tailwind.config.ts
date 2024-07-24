import { join } from 'path';
import type { Config } from 'tailwindcss';

import { skeleton } from '@skeletonlabs/tw-plugin';

const config = {
    darkMode: 'class',
    content: [
        './src/**/*.{html,js,svelte,ts}',
        join(require.resolve(
            '@skeletonlabs/skeleton'),
            '../**/*.{html,js,svelte,ts}'
        )
    ],
    theme: {
        extend: {},
    },
    plugins: [
        skeleton({
            themes: {
                preset: [
                    // Enable 'enhancements' per each registered theme:
                    { name: "vintage", enhancements: true }
                ]
            }
        })
    ]
} satisfies Config;

export default config;
