import type { Meta, StoryObj } from '@storybook/react';

import { Switch } from './switch';

const meta: Meta<typeof Switch> = {
	title: 'Components/UI/Switch',
	component: Switch,
	tags: ['autodocs'],
	parameters: {
		layout: 'centered',
	},
};

export default meta;

type Story = StoryObj<typeof meta>;

export const SwitchSample: Story = {
	render: () => (
		<div className="flex items-center gap-4">
			<Switch />
			<span>機内モード</span>
		</div>
	),
};
