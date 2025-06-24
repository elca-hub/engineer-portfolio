import { ButtonHTMLAttributes, DetailedHTMLProps } from 'react'
import { RiCloseLine } from 'react-icons/ri'

export default function CloseButton({ className, ...props }: DetailedHTMLProps<ButtonHTMLAttributes<HTMLButtonElement>, HTMLButtonElement>) {
	return (
		<button className={`bp-2 rounded shadow hover:bg-gray-300 transition-all duration-300 border ${className || ''}`} {...props}>
			<RiCloseLine className="size-8" />
		</button>
	)
}
