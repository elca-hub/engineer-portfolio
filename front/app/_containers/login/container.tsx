'use server'

import { handleAuthRedirect } from '@/lib/access'
import { redirect } from 'next/navigation'
import UserLoginPresentation from './presentation'

export default async function UserLoginContainer() {
	const res = await handleAuthRedirect('/login')
	if (res.isRedirect) redirect(res.redirectPath)

	return <UserLoginPresentation></UserLoginPresentation>
}
