'use server'

import RegisterPresentation from '@/app/_containers/register/presentation'
import { handleAuthRedirect } from '@/lib/access'
import { redirect } from 'next/navigation'

export default async function RegisterContainer() {
	const authRedirectPath = await handleAuthRedirect('register')
	if (authRedirectPath) redirect(authRedirectPath)

	return (
		<>
			<RegisterPresentation></RegisterPresentation>
		</>
	)
}
