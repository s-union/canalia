import type { Meta, StoryObj } from '@storybook/react';
import { FileInput } from './fileinput';

import { FileUp } from 'lucide-react';

const meta: Meta<typeof FileInput> = {
	title: 'Components/UI/FileInput',
	component: FileInput,
	tags: ['autodocs'],
	parameters: {
		layout: 'centered',
	},
	argTypes: {
		children: {
			control: 'text',
			description: 'ボタンのテキスト',
		},
		inputProps: {
			control: 'object',
			description: 'input要素のプロパティ',
		},
		variant: {
			control: 'select',
			options: ['default', 'outline', 'ghost'],
			description: 'ボタンのスタイル',
		},
		size: {
			control: 'select',
			options: ['default', 'lg'],
			description: 'ボタンのサイズ',
		},
	},
};

export default meta;

type Story = StoryObj<typeof meta>;

export const DefaultFileInput: Story = {
	render: () => (
		<FileInput
			inputProps={{
				accept: 'image/*',
				multiple: true,
			}}
		>
			<FileUp />
			ファイルアップロード
		</FileInput>
	),
};
