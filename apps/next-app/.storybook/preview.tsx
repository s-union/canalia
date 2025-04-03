// biome-ignore lint/correctness/noUnusedImports: <explanation>
import React from 'react';

import type { Preview } from '@storybook/react';

import { notoSansJP } from '../src/app/layout';
import '../src/app/globals.css';

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
			<div className={`${notoSansJP.variable} font-default`}>
				<Story />
			</div>
		),
	],
};

export default preview;
