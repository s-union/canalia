import type { Meta, StoryObj } from '@storybook/nextjs';

import { Ellipsis, X } from 'lucide-react';

import { Button } from './button';
import {
	Card,
	CardAction,
	CardContent,
	CardDescription,
	CardFooter,
	CardHeader,
	CardTitle,
} from './card';

const meta: Meta<typeof Card> = {
	title: 'Components/UI/Card',
	component: Card,
	tags: ['autodocs'],
	parameters: {
		layout: 'centered',
	},
};

export default meta;

type Story = StoryObj<typeof meta>;

export const DefaultCard: Story = {
	render: () => (
		<Card className="w-[350px]">
			<CardHeader>
				<CardTitle>カードタイトル</CardTitle>
				<CardDescription>カードの説明文がここに入ります。</CardDescription>
				<CardAction>
					<Button variant="ghost" size="icon-sm">
						<Ellipsis />
					</Button>
				</CardAction>
			</CardHeader>
			<CardContent>
				<p>カードのコンテンツ部分です。ここには主要な情報を表示します。</p>
				<p className="mt-4">複数の段落や他のコンポーネントも配置できます。</p>
			</CardContent>
			<CardFooter className="justify-between">
				<Button variant="outline">キャンセル</Button>
				<Button>保存</Button>
			</CardFooter>
		</Card>
	),
};

export const SimpleCard: Story = {
	render: () => (
		<Card className="w-[350px]">
			<CardHeader>
				<CardTitle>シンプルカード</CardTitle>
			</CardHeader>
			<CardContent>
				<p>必要最小限のコンポーネントだけを使った例です。</p>
			</CardContent>
		</Card>
	),
};

export const WithActionCard: Story = {
	render: () => (
		<Card className="w-[350px]">
			<CardHeader>
				<CardTitle>通知</CardTitle>
				<CardDescription>新しいメッセージが2件あります</CardDescription>
				<CardAction>
					<Button variant="ghost" size="icon-sm">
						<X />
					</Button>
				</CardAction>
			</CardHeader>
			<CardContent>
				<p>重要なお知らせがあります。確認してください。</p>
			</CardContent>
		</Card>
	),
};

export const CompleteCard: Story = {
	render: () => (
		<Card className="w-[450px]">
			<CardHeader className="border-b">
				<CardTitle>プロフィール</CardTitle>
				<CardDescription>ユーザー情報の詳細</CardDescription>
				<CardAction>
					<Button variant="outline" size="default">
						編集
					</Button>
				</CardAction>
			</CardHeader>
			<CardContent>
				<div className="space-y-4">
					<div>
						<div className="font-semibold text-sm">名前</div>
						<p>田中 太郎</p>
					</div>
					<div>
						<div className="font-semibold text-sm">メールアドレス</div>
						<p>tanaka.taro@example.com</p>
					</div>
					<div>
						<div className="font-semibold text-sm">役職</div>
						<p>シニアエンジニア</p>
					</div>
				</div>
			</CardContent>
			<CardFooter className="justify-end gap-2 border-t">
				<Button variant="outline">キャンセル</Button>
				<Button>保存</Button>
			</CardFooter>
		</Card>
	),
};
