export type DPResponseData<T> = {
	data?: T
	errors?: string[]
}

export async function NewDPResponse<T>(res: Response): Promise<DPResponseData<T>> {
	const resJson = await res.json()
	if (res.ok) {
		return { data: resJson as T }
	}

	return { errors: resJson.errors }
}

export function validateUUID(id: string): boolean {
	const uuidRegex = /^[0-9a-f]{8}-[0-9a-f]{4}-[1-5][0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$/i

	return uuidRegex.test(id)
}

export function validateUserId(userId: string): boolean {
	const maxUserIdLength = 50
	const excludeStrings = [
		' ',
		'@',
		'#',
		'$',
		'%',
		'&',
		'',
		'(',
		')',
		'+',
		'=',
		'{',
		'}',
		'[',
		']',
		'|',
		'\\',
		':',
		';',
		'"',
		"'",
		'<',
		'>',
		',',
		'.',
		'? ',
	]

	for (const str of excludeStrings) {
		if (userId.includes(str)) {
			return false
		}
	}

	if (userId.length < 0 || userId.length > maxUserIdLength) {
		return false
	}

	return true
}
