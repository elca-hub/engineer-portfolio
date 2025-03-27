'use client'

import DPHeader from '@/components/layout/header'
import DPButton from '@/components/ui/button/button'
import TextWithIcon from '@/components/ui/text/textWithIcon'
import { signOut } from 'next-auth/react'
import { RiLockLine } from 'react-icons/ri'

type Props = {
	userIconPath: string
}

export default function HeaderPresentation({ userIconPath }: Props) {
	return (
		<DPHeader userIconPath={userIconPath}>
			<DPButton colormode="primary" onClick={() => signOut()}>
				<TextWithIcon icon={<RiLockLine />}>ログアウト</TextWithIcon>
			</DPButton>
		</DPHeader>
	)
}
