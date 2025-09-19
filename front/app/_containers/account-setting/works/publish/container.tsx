'use server'

import { handleAuthRedirect, getAuthUser } from '@/lib/access'
import { redirect } from 'next/navigation'
import PublishWorkPresentation from './presentation'
import findWork from '@/action/usecase/works/findWork'

type PublishWorkContainerProps = {
	workId: string
}

export default async function PublishWorkContainer({ workId }: PublishWorkContainerProps) {
	const authRedirectPath = await handleAuthRedirect('auth')
	if (authRedirectPath) redirect(authRedirectPath)

	const authUser = await getAuthUser()

	if (authUser.errors || !authUser.data) {
		redirect('/')
	}

	const userData = authUser.data.user

	const work = await findWork(userData.user_id, workId)

	if (work.errors || !work.data) {
		redirect('/')
	}

	const workData = work.data.work

	return (
		<>
			<PublishWorkPresentation work={workData} user={userData} />
		</>
	)
}
