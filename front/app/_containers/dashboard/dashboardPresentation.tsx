'use client'

import DPButton from '@/components/ui/button/button'
import TextWithIcon from '@/components/ui/text/textWithIcon'
import { signOut } from 'next-auth/react'
import { RiGoogleFill, RiLockLine } from 'react-icons/ri'

/**
 * @package
 */
export default function DashboardPresentation() {
	return (
		<div className="flex h-screen flex-col items-center justify-center">
			<header className="mb-6">
				<TextWithIcon icon={<RiLockLine />} size="text-4xl">
					<h1 className="text-4xl font-bold tracking-widest text-foreground">ダッシュボード</h1>
				</TextWithIcon>
			</header>

			<main className="flex w-1/3 flex-col gap-4">
				<DPButton className="flex justify-center" colormode="primary" onClick={() => signOut()}>
					<TextWithIcon icon={<RiGoogleFill />}>サインアウト</TextWithIcon>
				</DPButton>
			</main>
		</div>
	)
}
