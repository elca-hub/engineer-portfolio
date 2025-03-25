'use server'

import { isNewUser } from '@/app/_containers/dashboard/action'
import NewUserPresentation from '@/app/_containers/dashboard/newUserPresentation'
import { authOptions } from '@/app/api/auth/[...nextauth]/route'
import { getServerSession } from 'next-auth'

export default async function DashboardContainer() {
	const session = await getServerSession(authOptions)

	if (!session) {
		return null
	}

	const sessionUser = session.user

	if (!sessionUser) {
		return null
	}

	const isNew = await isNewUser(sessionUser.email ?? '')

	console.log(isNew)

	if (isNew) {
		return <NewUserPresentation></NewUserPresentation>
	} else {
		return null
	}
}
