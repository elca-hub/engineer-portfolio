type TextWithIconProps = {
	icon: React.ReactNode
	children: React.ReactNode
	size?: string
}

export default function TextWithIcon({ icon, children, size }: TextWithIconProps) {
	return (
		<div className="flex items-center gap-2">
			<span className={size}>{icon}</span>
			{children}
		</div>
	)
}
