'use client'

import { UserType } from '@/action/type/user'
import createWork from '@/action/usecase/works/createWork'
import { CalloutContext } from '@/app/state'
import { getSessionToken } from '@/lib/access'
import { useRouter } from 'next/navigation'
import { useContext, useEffect } from 'react'

type Props = {
	user: UserType
}

/**
 * @package
 */
export default function AddWorkPresentation({ user }: Props) {
	const { callout, setCallout } = useContext(CalloutContext)
	const router = useRouter()

	useEffect(() => {
		if (user) {
			const createWorkLogic = async () => {
				const token = await getSessionToken()
				if (!token) {
					setCallout([...callout, { content: 'ログインしてください', type: 'error' }])
					return
				}

				const res = await createWork(token)
				if (res.data?.work) {
					router.push(`/account/setting/works/${res.data.work.id}/edit`)
				} else {
					setCallout([...callout, { content: '作品の作成に失敗しました。再度ログインしてください', type: 'error' }])
				}
			}

			createWorkLogic()
		}
	}, [user])

	return <></>
}
