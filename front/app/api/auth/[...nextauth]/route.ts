import { loginFlow } from '@/lib/access'
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
			const res = await loginFlow(user.email || '')

			console.log('res', res)

			if (res.errors) {
				return false
			}

			if (res.data) {
				return `/${res.data.id}/profile`
			} else {
				return false
			}
		},
	},
}

const handler = NextAuth(authOptions)
export { handler as GET, handler as POST }
