import type { Meta, StoryObj } from '@storybook/react';

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
			description: '最小値',
		},
		max: {
			control: 'number',
			description: '最大値',
		},
		defaultValue: {
			description: 'デフォルト値',
		},
		value: {
			description: '値',
		},
		className: {
			control: 'text',
			description: 'カスタムクラス',
		},
	},
};

export default meta;

type Story = StoryObj<typeof meta>;

export const DefaultSlider: Story = {
	render: () => <Slider defaultValue={[50]} max={100} step={1} />,
};
