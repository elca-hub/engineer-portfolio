'use server'

import { isExists } from '@/app/_containers/register/action'
import RegisterPresentation from '@/app/_containers/register/presentation'
import { authOptions } from '@/app/api/auth/[...nextauth]/route'
import { getServerSession } from 'next-auth'
import { redirect } from 'next/navigation'

export default async function RegisterContainer() {
	const session = await getServerSession(authOptions)

	if (!session) {
		redirect('/login')
	}

	const sessionUser = session.user

	if (!sessionUser) {
		redirect('/login')
	}

	const isAlreadyRegistration = await isExists(sessionUser.email ?? '')

	if (isAlreadyRegistration) {
		redirect('/dashboard')
	}

	return (
		<>
			<RegisterPresentation></RegisterPresentation>
		</>
	)
}
