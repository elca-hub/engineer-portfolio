'use server'

import { isNewUser } from '@/app/_containers/dashboard/action'
import DashboardPresentation from '@/app/_containers/dashboard/dashboardPresentation'
import NewUserPresentation from '@/app/_containers/dashboard/newUserPresentation'
import { authOptions } from '@/app/api/auth/[...nextauth]/route'
import { getServerSession } from 'next-auth'
import { redirect } from 'next/navigation'

export default async function DashboardContainer() {
	const session = await getServerSession(authOptions)

	if (!session) {
		redirect('/login')
	}

	const sessionUser = session.user

	if (!sessionUser) {
		redirect('/login')
	}

	const isNew = await isNewUser(sessionUser.email ?? '')

	console.log(isNew)

	if (isNew) {
		return <NewUserPresentation></NewUserPresentation>
	} else {
		return <DashboardPresentation></DashboardPresentation>
	}
}
