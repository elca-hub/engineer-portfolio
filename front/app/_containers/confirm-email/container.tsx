'use server'

import { cookieConfirmName } from '@/constants/cookieTokenName'
import { verifyJwt } from '@/lib/auth/jwt'
import { redirectIfAlreadyLogin } from '@/lib/auth/verifySession'
import { cookies } from 'next/headers'
import { redirect } from 'next/navigation'
import ConfirmEmailPresenter from './presentation'

export default async function ConfirmEmailContainer() {
	await redirectIfAlreadyLogin() // 既にログインしている場合はリダイレクト

	// 1日以上経過し、cookieに保存しているtokenがない場合は新規登録画面へリダイレクト
	const userCookie = await cookies()
	const emailToken = userCookie.get(cookieConfirmName)
	const verify = await verifyJwt<{ email: string }>(emailToken?.value || '')
	if (verify === null) {
		redirect('/register?userState=expiredEmail')
	}

	return <ConfirmEmailPresenter></ConfirmEmailPresenter>
}
