import { Slot } from '@radix-ui/react-slot';
import { type VariantProps, cva } from 'class-variance-authority';
// biome-ignore lint/style/useImportType: <explanation>
import * as React from 'react';

import { cn } from '@/lib/utils';

const buttonVariants = cva(
	"inline-flex shrink-0 items-center justify-center gap-2 whitespace-nowrap font-medium text-sm outline-none transition-all focus-visible:border-ring focus-visible:ring-[3px] focus-visible:ring-ring/50 disabled:pointer-events-none disabled:bg-info aria-invalid:border-error aria-invalid:ring-error/20 dark:aria-invalid:ring-error/40 [&_svg:not([class*='size-'])]:size-4 [&_svg]:pointer-events-none [&_svg]:shrink-0",
	{
		variants: {
			variant: {
				primary: 'bg-primary text-white shadow-xs hover:bg-primary/80',
				secondary: 'bg-secondary text-white shadow-xs hover:bg-secondary/80',
				error: 'bg-error text-white shadow-xs hover:bg-error/80',
				success: 'bg-success text-white shadow-xs hover:bg-success/80',
				outline:
					'border bg-background-primary shadow-xs hover:bg-accent hover:text-accent-foreground dark:border-input dark:bg-input/30 dark:hover:bg-input/50',
				ghost:
					'hover:bg-accent hover:text-accent-foreground dark:hover:bg-accent/50',
			},
			size: {
				default: 'h-9 rounded-xl px-4 py-2 has-[>svg]:px-3',
				lg: 'h-12 rounded-xl px-6 py-3 text-xl has-[>svg]:px-4',
				icon: "size-12 rounded-full [&_svg:not([class*='size-'])]:size-8",
			},
		},
		defaultVariants: {
			variant: 'primary',
			size: 'default',
		},
	},
);

function Button({
	className,
	variant,
	size,
	asChild = false,
	...props
}: React.ComponentProps<'button'> &
	VariantProps<typeof buttonVariants> & {
		asChild?: boolean;
	}) {
	const Comp = asChild ? Slot : 'button';

	return (
		<Comp
			data-slot="button"
			className={cn(buttonVariants({ variant, size, className }))}
			{...props}
		/>
	);
}

export { Button, buttonVariants };
