'use server'
import { apiPrefix } from '@/constants/api'
import { ReadonlyRequestCookies } from 'next/dist/server/web/spec-extension/adapters/request-cookies'

export default async function Fetcher(path: string, cookieStore: ReadonlyRequestCookies, options?: RequestInit) {
	const csrfToken = cookieStore.get('csrf_token')?.value ?? ''

	const res = await fetch(`${apiPrefix}/${path}`, {
		...options,
		headers: {
			'Content-Type': 'application/json',
			'X-CSRF-Token': csrfToken,
			...options?.headers,
		},
	})

	const resCookie = res.headers.get('set-cookie')
	console.log(resCookie)

	if (resCookie) {
		// parse cookie
		const resCookieArray = resCookie.split(';')
		const resCookieObject = resCookieArray.reduce<Record<string, string>>((acc, cur) => {
			const [key, value] = cur.split('=')
			acc[key] = value
			return acc
		}, {})
		console.log(resCookieObject)

		// 全てのcookieをcookieStoreに保存
		Object.entries(resCookieObject).forEach(([key, value]) => {
			cookieStore.set(key, value)
		})
	} else {
		console.log('no cookie')
	}

	return res
}
