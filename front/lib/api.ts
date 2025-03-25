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
