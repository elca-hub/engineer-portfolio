'use server'

import { isLogin } from '@/lib/auth/verifySession'
import { redirect } from 'next/navigation'
import UserRegisterPresentation from './presentation'

export default async function UserRegisterContainer() {
	const isAlreadyLogin = await isLogin()

	if (isAlreadyLogin) {
		redirect('/user')
	}

	return <UserRegisterPresentation></UserRegisterPresentation>
}
