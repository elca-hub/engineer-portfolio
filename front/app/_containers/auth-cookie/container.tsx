'use server'

import AuthCookiePresentation from '@/app/_containers/auth-cookie/presentation'
import { authOptions } from '@/app/api/auth/[...nextauth]/route'
import { handleAuthRedirect } from '@/lib/access'
import { getServerSession } from 'next-auth'
import { redirect } from 'next/navigation'

export default async function AuthCookieContainer() {
	const authRedirectPath = await handleAuthRedirect('expired-cookie')
	if (authRedirectPath) redirect(authRedirectPath)

	const session = await getServerSession(authOptions)

	if (!session) redirect('/')

	const email = session.user?.email || ''

	return (
		<>
			<AuthCookiePresentation email={email} />
		</>
	)
}
