'use server'

import { UserType } from '@/action/type/user'
import { apiPrefix } from '@/constants/constant'
import { NewDPResponse } from '@/lib/api'

export default async function userUpdate(
	token: string,
	{ userData, icon, headerIcon }: { userData?: Partial<UserType>; icon?: File; headerIcon?: File },
): Promise<UserType | null> {
	const formData = new FormData()

	// ユーザーデータをJSONとして追加
	if (userData) {
		formData.append('user_data', JSON.stringify(userData))
	}

	// 画像ファイルがある場合は追加
	if (icon) {
		formData.append('icon', icon)
	}
	if (headerIcon) {
		formData.append('header_icon', headerIcon)
	}

	const updateUser = await fetch(`${apiPrefix}/auth/user/`, {
		method: 'PUT',
		body: formData,
		headers: {
			Authorization: `Bearer ${token}`,
		},
	})

	const updateUserRes = await NewDPResponse<{ user: UserType }>(updateUser)

	if (updateUserRes.errors) {
		return null
	}

	return updateUserRes.data?.user ?? null
}
