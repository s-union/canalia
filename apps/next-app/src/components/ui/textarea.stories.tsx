import type { Meta, StoryObj } from '@storybook/react';
import { Textarea } from './textarea';

const meta: Meta<typeof Textarea> = {
	title: 'Components/UI/Textarea',
	component: Textarea,
	tags: ['autodocs'],
	parameters: {
		layout: 'centered',
	},
	argTypes: {
		placeholder: {
			control: 'text',
			description: 'プレースホルダー',
		},
		disabled: {
			control: 'boolean',
			description: 'disabled要素',
		},
		className: {
			control: 'text',
			description: 'カスタムクラス名',
		},
	},
};

export default meta;

type Story = StoryObj<typeof meta>;

export const DefaultTextarea: Story = {
	args: {
		placeholder: 'メッセージを入力してください',
		className: 'w-96',
	},
};

export const DisabledTextarea: Story = {
	args: {
		placeholder: 'メッセージを入力してください',
		className: 'w-96',
		disabled: true,
	},
};
