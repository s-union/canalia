import type { Meta, StoryObj } from '@storybook/nextjs';

import { Slider } from './slider';

const meta: Meta<typeof Slider> = {
	title: 'Components/UI/Slider',
	component: Slider,
	tags: ['autodocs'],
	parameters: {
		layout: 'centered',
	},
	argTypes: {
		min: {
			control: 'number',
			description: 'The minimum value of the slider.',
		},
		max: {
			control: 'number',
			description: 'The maximum value of the slider.',
		},
		variant: {
			control: 'select',
			options: ['primary', 'secondary'],
			description: 'The color variant of the slider.',
		},
	},
};

export default meta;

type Story = StoryObj<typeof meta>;

export const Primary: Story = {
	args: {
		className: 'w-60',
		min: 0,
		max: 100,
		variant: 'primary',
	},
};

export const Secondary: Story = {
	args: {
		className: 'w-60',
		min: 0,
		max: 100,
		variant: 'secondary',
	},
};
