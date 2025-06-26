'use client';

import * as React from 'react';
import * as SliderPrimitive from '@radix-ui/react-slider';
import { cva, type VariantProps } from 'class-variance-authority';

import { cn } from '@/lib/utils';

const sliderTrackVariants = cva(
	'relative grow overflow-hidden rounded-full data-[orientation=horizontal]:h-1.5 data-[orientation=vertical]:h-full data-[orientation=horizontal]:w-full data-[orientation=vertical]:w-1.5',
	{
		variants: {
			variant: {
				primary: 'bg-input',
				secondary: 'bg-input',
			},
		},
		defaultVariants: {
			variant: 'primary',
		},
	},
);

const sliderRangeVariants = cva(
	'absolute data-[orientation=horizontal]:h-full data-[orientation=vertical]:w-full',
	{
		variants: {
			variant: {
				primary: 'bg-primary',
				secondary: 'bg-secondary',
			},
		},
		defaultVariants: {
			variant: 'primary',
		},
	},
);

const sliderThumbVariants = cva(
	'block size-4 shrink-0 rounded-full border shadow-sm transition-[color,box-shadow] focus-visible:outline-hidden disabled:pointer-events-none disabled:opacity-50',
	{
		variants: {
			variant: {
				primary: 'border-primary bg-primary',
				secondary: 'border-secondary bg-secondary',
			},
		},
		defaultVariants: {
			variant: 'primary',
		},
	},
);

interface SliderProps
	extends React.ComponentPropsWithoutRef<typeof SliderPrimitive.Root>,
		VariantProps<typeof sliderTrackVariants> {}

function Slider({
	className,
	defaultValue,
	value,
	min = 0,
	max = 100,
	variant,
	...props
}: SliderProps) {
	const _values = React.useMemo(
		() =>
			Array.isArray(value)
				? value
				: Array.isArray(defaultValue)
					? defaultValue
					: [min, max],
		[value, defaultValue, min, max],
	);

	return (
		<SliderPrimitive.Root
			data-slot="slider"
			defaultValue={defaultValue}
			value={value}
			min={min}
			max={max}
			className={cn(
				'relative flex w-full touch-none select-none items-center data-[orientation=vertical]:h-full data-[orientation=vertical]:min-h-44 data-[orientation=vertical]:w-auto data-[orientation=vertical]:flex-col data-[disabled]:opacity-50',
				className,
			)}
			{...props}
		>
			<SliderPrimitive.Track
				data-slot="slider-track"
				className={cn(sliderTrackVariants({ variant }))}
			>
				<SliderPrimitive.Range
					data-slot="slider-range"
					className={cn(sliderRangeVariants({ variant }))}
				/>
			</SliderPrimitive.Track>
			{Array.from({ length: _values.length }, (_, index) => (
				<SliderPrimitive.Thumb
					data-slot="slider-thumb"
					key={_values[index]}
					className={cn(sliderThumbVariants({ variant }))}
				/>
			))}
		</SliderPrimitive.Root>
	);
}

export { Slider };
