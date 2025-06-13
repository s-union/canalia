// biome-ignore lint/correctness/noUnusedImports: <explanation>
import React from 'react';

import type { Preview } from '@storybook/nextjs';
import { Noto_Sans_JP } from 'next/font/google';

import '../src/app/globals.css';

const notoSansJP = Noto_Sans_JP({
	variable: '--font-noto-sans-jp',
	display: 'swap',
	preload: false,
	weight: ['400', '700'],
});

const preview: Preview = {
	parameters: {
		controls: {
			matchers: {
				color: /(background|color)$/i,
				date: /Date$/i,
			},
		},
	},
	decorators: [
		(Story) => (
			<div className={`${notoSansJP.variable} font-sans`}>
				<Story />
			</div>
		),
	],
};

export default preview;
