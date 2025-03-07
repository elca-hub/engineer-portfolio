import { Metadata } from 'next'

export const setMetadata = (title: string, description: string): Metadata => {
	return {
		title: `${title} | DevPort`,
		description: description,
		icons: {
			icon: '/favicon.ico',
		},
	}
}
