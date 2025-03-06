'use server'

import { JWTPayload, SignJWT, jwtVerify } from 'jose'

const secretKey = new TextEncoder().encode(process.env.JWT_SECRET_KEY)

async function makeJwt<T>(payload: T): Promise<string> {
	const pay: JWTPayload = {
		...payload,
		exp: Math.floor(Date.now() / 1000) + 60 * 60 * 24, // 1 day
	}
	const token = await new SignJWT(pay).setProtectedHeader({ alg: 'HS256' }).sign(secretKey)

	return token
}

async function verifyJwt<T>(token: string): Promise<T | null> {
	try {
		const { payload } = await jwtVerify(token, secretKey)

		return payload as T
		// eslint-disable-next-line @typescript-eslint/no-unused-vars
	} catch (e) {
		return null
	}
}

export { makeJwt, verifyJwt }
