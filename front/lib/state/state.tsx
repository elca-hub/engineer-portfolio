'use server'

import { cookies } from 'next/headers'

export type BPState = 'inConfirmEmailButCookieExpired'

const stateKey = 'bp_state'

export async function setBPState(state: BPState) {
	const cookie = await cookies()
	cookie.set(stateKey, state, {
		maxAge: 10,
	})
}

export async function getBPState(): Promise<BPState | null> {
	const cookie = await cookies()
	if (!cookie.get('bp_state')) return null
	const bpState = cookie.get(stateKey)?.value
	return bpState ? (bpState as BPState) : null
}

export async function clearBPState() {
	const cookie = await cookies()
	cookie.delete(stateKey)
}
