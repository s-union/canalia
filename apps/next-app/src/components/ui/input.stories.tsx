import type { Meta, StoryObj } from '@storybook/nextjs';
import { Input } from './input';

const meta: Meta<typeof Input> = {
	title: 'Components/UI/Input',
	component: Input,
	tags: ['autodocs'],
	parameters: {
		layout: 'centered',
	},
	argTypes: {
		type: {
			control: 'select',
			description: '入力フィールドのタイプ',
			options: [
				'text',
				'email',
				'password',
				'number',
				'search',
				'tel',
				'url',
				'date',
			],
		},
		placeholder: {
			control: 'text',
			description: 'プレースホルダー',
		},
		disabled: {
			control: 'boolean',
			description: 'disabled属性',
		},
		className: {
			control: 'text',
			description: 'カスタムクラス',
		},
	},
};

export default meta;

type Story = StoryObj<typeof meta>;

export const DefaultInput: Story = {
	args: {
		type: 'text',
		placeholder: 'input',
	},
};

export const DisabledInput: Story = {
	args: {
		type: 'text',
		placeholder: 'input',
		disabled: true,
	},
};
