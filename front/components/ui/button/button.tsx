import { ButtonHTMLAttributes, DetailedHTMLProps } from 'react'

interface DBButtonProps extends DetailedHTMLProps<ButtonHTMLAttributes<HTMLButtonElement>, HTMLButtonElement> {
	colormode: 'primary' | 'secondary'
}

const DPButton = ({ ...props }: DBButtonProps) => {
	const convertColorMode = (colormode: 'primary' | 'secondary') => {
		return colormode === 'primary' ? 'bg-primary text-foreground' : 'bg-secondary text-white'
	}

	return (
		<button
			{...props}
			className={`
        ${convertColorMode(props.colormode)}
        rounded px-4
        py-2 text-lg
        font-bold
        transition-all duration-300
        hover:scale-[0.98] hover:opacity-80
        focus:outline-none focus:ring-2 focus:ring-sky-500
      `}
		>
			{props.children}
		</button>
	)
}

export default DPButton
