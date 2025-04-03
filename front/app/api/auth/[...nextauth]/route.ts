import { handleAuthRedirect } from '@/lib/access'
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
		signIn: async ({}) => {
			const authRedirectHandle = await handleAuthRedirect('/login')

			if (authRedirectHandle.isRedirect) {
				return authRedirectHandle.redirectPath
			} else {
				return true
			}
		},
	},
}

const handler = NextAuth(authOptions)
export { handler as GET, handler as POST }
