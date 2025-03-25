'use client'

import DPButton from '@/components/ui/button/button'
import TextWithIcon from '@/components/ui/text/textWithIcon'
import { signIn } from 'next-auth/react'
import { useSearchParams } from 'next/navigation'
import { RiGoogleFill, RiLockLine } from 'react-icons/ri'

/**
 * @package
 */
export default function DashboardPresentation() {
	const callbackUrl = useSearchParams().get('callbackUrl') || '/dashboard'

	return (
		<div className="flex h-screen flex-col items-center justify-center">
			<header className="mb-6">
				<TextWithIcon icon={<RiLockLine />} size="text-4xl">
					<h1 className="text-4xl font-bold tracking-widest text-foreground">ダッシュボード</h1>
				</TextWithIcon>
			</header>

			<main className="flex w-1/3 flex-col gap-4">
				<DPButton className="flex justify-center" colormode="primary" onClick={() => signIn('google', { callbackUrl })}>
					<TextWithIcon icon={<RiGoogleFill />}>メールアドレスでログイン</TextWithIcon>
				</DPButton>
			</main>
		</div>
	)
}
