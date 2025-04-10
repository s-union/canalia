import type * as React from 'react';
import { useRef } from 'react';

import { Button } from './button';

type NativeInputProps = React.InputHTMLAttributes<HTMLInputElement>;
type NativeButtonProps = React.ComponentProps<typeof Button>;

type InputProps = Omit<
	NativeInputProps,
	'type' | 'style' | 'className' | 'ref'
>;
type ButtonProps = Omit<NativeButtonProps, 'onClick'>;

type ButtonWithInputFileProps = ButtonProps & { inputProps: InputProps };

export function FileInput({
	children,
	inputProps,
	variant = 'outline',
	size = 'default',
	...buttonProps
}: ButtonWithInputFileProps) {
	const inputFileRef = useRef<HTMLInputElement>(null);

	return (
		<>
			<Button
				variant={variant}
				size={size}
				onClick={() => inputFileRef.current?.click()}
				{...buttonProps}
			>
				{children}
			</Button>
			<input
				ref={inputFileRef}
				type="file"
				className="hidden"
				{...inputProps}
			/>
		</>
	);
}
