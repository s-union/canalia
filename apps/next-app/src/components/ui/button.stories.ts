import type { Meta, StoryObj } from '@storybook/react';
import { fn } from '@storybook/test';

import { Button } from './button';

const meta: Meta<typeof Button> = {
	title: 'Components/UI/Button',
	component: Button,
	tags: ['autodocs'],
	parameters: {
		layout: 'centered',
	},
	argTypes: {
		onClick: { action: 'clicked' },
		variant: {
			control: 'select',
			description: 'The variant of the button',
			options: ['default', 'outline', 'destructive', 'secondary', 'ghost'],
		},
		size: {
			control: 'select',
			description: 'The size of the button',
			options: ['sm', 'default', 'lg', 'icon'],
		},
		disabled: {
			control: 'boolean',
			description: 'If the button is disabled',
		},
		children: {
			control: 'text',
			description: 'The content of the button',
		},
		className: {
			control: 'text',
			description: 'Custom tailwind CSS classes to apply to the button',
		},
	},
};

export default meta;

type Story = StoryObj<typeof meta>;

export const Default: Story = {
	args: {
		variant: 'default',
		size: 'default',
		disabled: false,
		onClick: fn(),
		children: 'Default Button',
	},
};
