import { apiPrefix } from '@/constants/constant'
import { loginFlow } from '@/lib/access'
import { NewDPResponse } from '@/lib/api'
import NextAuth, { NextAuthOptions } from 'next-auth'
import GoogleProvider from 'next-auth/providers/google'

export const authOptions: NextAuthOptions = {
	providers: [
		GoogleProvider({
			clientId: process.env.GOOGLE_CLIENT_ID as string,
			clientSecret: process.env.GOOGLE_CLIENT_SECRET as string,
		}),
	],
	pages: {
		signIn: '/login',
	},
	session: {
		strategy: 'jwt',
	},
	callbacks: {
		jwt: ({ token, user }) => {
			if (user) {
				// eslint-disable-next-line @typescript-eslint/no-explicit-any
				const u = user as unknown as any
				return {
					...token,
					id: u.id,
				}
			}
			return token
		},
		session: ({ session, token }) => {
			return {
				...session,
				user: {
					...session.user,
					id: token.id,
				},
			}
		},
		signIn: async ({ user }) => {
			const rawRes = await fetch(`${apiPrefix}/is_exists?email=${user.email ?? ''}`, {
				method: 'GET',
			})

			const res = await NewDPResponse<{ is_exists: boolean; user_id: string }>(rawRes)

			if (res.data) {
				if (res.data.is_exists) {
					const loginRes = await loginFlow(user.email ?? '')
					if (loginRes.errors) {
						const errorJson = (await rawRes.json()) as { errors: string[] }
						errorJson.errors.forEach((value) => console.error(value))
						return false
					}
					return true
				} else {
					return '/register'
				}
			} else {
				if (res.errors) {
					res.errors.forEach((value) => console.error(value))
				}
				return false
			}
		},
	},
}

const handler = NextAuth(authOptions)
export { handler as GET, handler as POST }
