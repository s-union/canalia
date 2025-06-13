import type { Meta, StoryObj } from '@storybook/nextjs';

import { Badge } from './badge';

const meta: Meta<typeof Badge> = {
	title: 'Components/UI/Badge',
	component: Badge,
	parameters: {
		layout: 'centered',
	},
	tags: ['autodocs'],
	argTypes: {
		variant: {
			control: 'select',
			options: [
				'info',
				'primary',
				'secondary',
				'error',
				'success',
				'warning',
				'outline',
			],
			description: 'バッジのスタイルバリアント',
		},
		children: {
			control: 'text',
			description: 'バッジのコンテンツ',
		},
		className: {
			control: 'text',
			description: 'バッジに適用するカスタムTailwind CSSクラス',
		},
		asChild: {
			control: 'boolean',
			description: 'コンポーネントをスロットとして使用するかどうか',
		},
	},
};

export default meta;

type Story = StoryObj<typeof meta>;

export const PrimaryBadge: Story = {
	args: {
		children: 'プライマリ',
		variant: 'primary',
	},
};

export const SecondaryBadge: Story = {
	args: {
		children: 'セカンダリ',
		variant: 'secondary',
	},
};

export const InfoBadge: Story = {
	args: {
		children: '受付終了',
		variant: 'info',
	},
};

export const ErrorBadge: Story = {
	args: {
		children: '重要',
		variant: 'error',
	},
};

export const SuccessBadge: Story = {
	args: {
		children: '提出済み',
		variant: 'success',
	},
};

export const WarningBadge: Story = {
	args: {
		children: '未提出',
		variant: 'warning',
	},
};

export const OutlineBadge: Story = {
	args: {
		children: 'アウトライン',
		variant: 'outline',
	},
};

// 複数のバッジを表示する例
export const AllVariants: Story = {
	render: () => (
		<div className="flex flex-wrap gap-2">
			<Badge variant="info">受付終了</Badge>
			<Badge variant="primary">プライマリ</Badge>
			<Badge variant="secondary">セカンダリ</Badge>
			<Badge variant="error">重要</Badge>
			<Badge variant="success">提出済み</Badge>
			<Badge variant="warning">未提出</Badge>
			<Badge variant="outline">アウトライン</Badge>
		</div>
	),
};
