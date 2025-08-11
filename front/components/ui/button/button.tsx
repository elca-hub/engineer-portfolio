import { Button, ButtonProps } from 'react-aria-components'

interface DBButtonProps extends ButtonProps, React.RefAttributes<HTMLButtonElement> {
	colormode: 'primary' | 'secondary' | 'mono'
	buttonSize?: 'small'
}

const DPButton = ({ ...props }: DBButtonProps) => {
	const convertColorMode = (colormode: 'primary' | 'secondary' | 'mono') => {
		switch (colormode) {
			case 'primary':
				return 'bg-primary text-foreground'
			case 'secondary':
				return 'bg-secondary text-white'
			case 'mono':
				return 'bg-white text-foreground border border-foreground'
		}
	}

	return (
		<Button
			{...props}
			className={`
				${props.className}
        ${convertColorMode(props.colormode)}
        rounded px-4
        py-2 ${props.buttonSize === 'small' ? 'text-sm' : 'text-lg'}
        font-bold
        transition-all duration-300
        hover:scale-[0.98] hover:opacity-80
        focus:outline-none focus:ring-2 focus:ring-sky-500
				disabled:opacity-50
      `}
		>
			{props.children}
		</Button>
	)
}

export default DPButton
