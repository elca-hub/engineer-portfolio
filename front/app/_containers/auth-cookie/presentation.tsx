'use client'

import { CalloutContext } from '@/app/state'
import { loginFlow } from '@/lib/access'
import { useRouter } from 'next/navigation'
import { useContext, useEffect } from 'react'

type Props = {
	email: string
}

/**
 * @package
 */
export default function AuthCookiePresentation({ email }: Props) {
	const { callout, setCallout } = useContext(CalloutContext)
	const router = useRouter()

	useEffect(() => {
		if (email) {
			const loginLogic = async () => {
				const res = await loginFlow(email)
				if (res.data?.isSuccess) {
					router.back()
				} else {
					setCallout([...callout, { content: 'ログイン情報の取得に失敗しました。再度ログインしてください', type: 'error' }])
				}
			}

			loginLogic()
		}
	}, [email])

	return <></>
}
