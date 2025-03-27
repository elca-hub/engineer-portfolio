'use client'

type Props = {
	header: React.ReactNode
}

/**
 * @package
 */
export default function DashboardPresentation({ header }: Props) {
	return (
		<div className="flex flex-col h-screen">
			{header}

			<main>
				<p>Dashboard</p>
			</main>
		</div>
	)
}
