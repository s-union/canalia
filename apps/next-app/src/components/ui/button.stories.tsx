import type { Meta, StoryObj } from '@storybook/nextjs';
import { fn } from 'storybook/test';

import { Upload } from 'lucide-react';

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
			options: ['primary', 'secondary', 'error', 'success', 'outline', 'ghost'],
		},
		size: {
			control: 'select',
			description: 'The size of the button',
			options: ['default', 'lg', 'icon', 'icon-sm'],
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

export const DefaultButton: Story = {
	args: {
		variant: 'primary',
		size: 'default',
		disabled: false,
		onClick: fn(),
		children: 'Upload',
	},
};

export const SecondaryButton: Story = {
	args: {
		variant: 'secondary',
		size: 'default',
		disabled: false,
		onClick: fn(),
		children: 'Upload',
	},
};

export const ErrorButton: Story = {
	args: {
		variant: 'error',
		size: 'default',
		disabled: false,
		onClick: fn(),
		children: 'Delete',
	},
};

export const SuccessButton: Story = {
	args: {
		variant: 'success',
		size: 'default',
		disabled: false,
		onClick: fn(),
		children: 'Save',
	},
};

export const OutlineButton: Story = {
	args: {
		variant: 'outline',
		size: 'default',
		disabled: false,
		onClick: fn(),
		children: 'Upload',
	},
};

export const GhostButton: Story = {
	args: {
		variant: 'ghost',
		size: 'default',
		disabled: false,
		onClick: fn(),
		children: 'Upload',
	},
};

export const DisabledButton: Story = {
	args: {
		variant: 'primary',
		size: 'default',
		disabled: true,
		onClick: fn(),
		children: 'Disabled',
	},
};

export const LargeButton: Story = {
	args: {
		variant: 'primary',
		size: 'lg',
		disabled: false,
		onClick: fn(),
		children: 'Upload',
	},
};

export const WithIconButton: Story = {
	args: {
		variant: 'primary',
		size: 'default',
		disabled: false,
		onClick: fn(),
		children: (
			<>
				<Upload /> Upload
			</>
		),
	},
};

export const IconButton: Story = {
	args: {
		variant: 'primary',
		size: 'icon',
		disabled: false,
		onClick: fn(),
		children: <Upload />,
	},
};

export const SmallIconButton: Story = {
	args: {
		variant: 'primary',
		size: 'icon-sm',
		disabled: false,
		onClick: fn(),
		children: <Upload />,
	},
};
