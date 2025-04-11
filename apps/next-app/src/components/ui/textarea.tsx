import type * as React from 'react';

import { cn } from '@/lib/utils';

function Textarea({ className, ...props }: React.ComponentProps<'textarea'>) {
	return (
		<textarea
			data-slot="textarea"
			className={cn(
				'field-sizing-content flex min-h-16 w-full rounded-md border border-gray bg-white px-3 py-2 text-base shadow-xs outline-none transition-[color,box-shadow] placeholder:text-light-gray focus-visible:border-primary-500 disabled:cursor-not-allowed disabled:opacity-50 aria-invalid:border-error aria-invalid:ring-error/20 md:text-sm dark:bg-input/30 dark:aria-invalid:ring-destructive/40',
				className,
			)}
			{...props}
		/>
	);
}

export { Textarea };
