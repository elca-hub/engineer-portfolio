'use server'

import AuthCookiePresentation from '@/app/_containers/auth-cookie/presentation'
import { authOptions } from '@/app/api/auth/[...nextauth]/route'
import { handleAuthRedirect } from '@/lib/access'
import { getServerSession } from 'next-auth'
import { redirect } from 'next/navigation'

export default async function AuthCookieContainer() {
	const res = await handleAuthRedirect('/dashboard')

	if (res.isRedirect) redirect(res.redirectPath)

	const session = await getServerSession(authOptions)

	if (!session) redirect('/login')

	const email = session.user?.email || ''

	return (
		<>
			<AuthCookiePresentation email={email} />
		</>
	)
}
