'use server'

// export default async function fetchBio(user: UserType): Promise<string | null> {
// 	console.log(user.bio_path)
// 	const res = await fetch(`http://minio:9000/devport/bio/${user.user_id}/${user.bio_path}.md`, {
// 		method: 'GET',
// 	})

// 	if (!res.ok) {
// 		console.error(res)
// 		return null
// 	}

// 	const bio = await res.text()

// 	return bio
// }
