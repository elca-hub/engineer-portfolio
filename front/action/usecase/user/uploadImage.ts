'use server'

import { apiPrefix } from '@/constants/constant'
import { NewDPResponse } from '@/lib/api'

type ResponseType = {
	icon_name: string
	header_icon_name: string
}

export default async function uploadImage(token: string, { icon, headerIcon }: { icon?: File; headerIcon?: File }): Promise<ResponseType | null> {
	const formData = new FormData()

	// 画像ファイルがある場合は追加
	if (icon) {
		formData.append('icon', icon)
	}
	if (headerIcon) {
		formData.append('header_icon', headerIcon)
	}

	const uploadImage = await fetch(`${apiPrefix}/auth/user/image`, {
		method: 'PUT',
		body: formData,
		headers: {
			Authorization: `Bearer ${token}`,
		},
	})

	const uploadImageRes = await NewDPResponse<ResponseType>(uploadImage)

	if (uploadImageRes.errors) {
		return null
	}

	return uploadImageRes.data ?? null
}
