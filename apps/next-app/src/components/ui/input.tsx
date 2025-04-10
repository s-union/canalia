import type * as React from 'react';

import { cn } from '@/lib/utils';

type InputType = Exclude<React.HTMLInputTypeAttribute, 'file'>;

/**
 * `type="file"` は利用できません。FileInputコンポーネントを使用してください。
 */
function Input({
	className,
	type,
	...props
}: Omit<React.ComponentProps<'input'>, 'type'> & { type?: InputType }) {
	return (
		<input
			type={type}
			data-slot="input"
			className={cn(
				'flex h-9 w-full min-w-0 rounded-md border border-gray bg-transparent px-3 py-1 text-base shadow-xs outline-none transition-[color,box-shadow] selection:bg-primary selection:text-black placeholder:text-light-gray disabled:pointer-events-none disabled:cursor-not-allowed disabled:opacity-50 md:text-sm dark:bg-input/30',
				'focus-visible:border-primary-500',
				'aria-invalid:border-error aria-invalid:ring-error/20 dark:aria-invalid:ring-destructive/40',
				className,
			)}
			{...props}
		/>
	);
}

export { Input };
