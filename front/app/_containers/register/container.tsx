'use server'

import RegisterPresentation from '@/app/_containers/register/presentation'
import { handleAuthRedirect } from '@/lib/access'
import { redirect } from 'next/navigation'

export default async function RegisterContainer() {
	const res = await handleAuthRedirect('/register')

	if (res.isRedirect) redirect(res.redirectPath)

	return (
		<>
			<RegisterPresentation></RegisterPresentation>
		</>
	)
}
