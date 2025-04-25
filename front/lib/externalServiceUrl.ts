/**
const (
	ExternalServiceOther = iota
	ExternalServiceGithub
	ExternalServiceX
	ExternalServiceQiita
	ExternalServiceZenn
	ExternalServiceNote
)
 */

export const MAX_SERVICE_TYPE_LEN = 5

export function ServiceTypeToString(serviceType: number) {
	switch (serviceType) {
		case 0:
			return 'other'
		case 1:
			return 'github'
		case 2:
			return 'x'
		case 3:
			return 'qiita'
		case 4:
			return 'zenn'
		case 5:
			return 'note'
		default:
			return 'other'
	}
}

export function ServiceTypeToServiceName(serviceType: number) {
	switch (serviceType) {
		case 1:
			return 'GitHub'
		case 2:
			return 'X'
		case 3:
			return 'Qiita'
		case 4:
			return 'Zenn'
		case 5:
			return 'note'
		default:
			return 'other'
	}
}

export function StringToServiceType(serviceType: string) {
	switch (serviceType) {
		case 'github':
			return 1
		case 'x':
			return 2
		case 'qiita':
			return 3
		case 'zenn':
			return 4
		case 'note':
			return 5
		default:
			return 0
	}
}

export function ImageNameByServiceType(serviceType: number) {
	switch (serviceType) {
		case 1:
			return 'github.png'
		case 2:
			return 'x.png'
		case 3:
			return 'qiita.png'
		case 4:
			return 'zenn.svg'
		case 5:
			return 'note.png'
		default:
			return 'other.svg'
	}
}

export function ConvertToExternalServiceUrl(serviceType: number, name: string): string {
	if (name === '') return ''
	switch (serviceType) {
		case 1:
			return `https://github.com/${name}`
		case 2:
			return `https://x.com/${name}`
		case 3:
			return `https://qiita.com/${name}`
		case 4:
			return `https://zenn.dev/${name}`
		case 5:
			return `https://note.com/${name}`
		default:
			return ''
	}
}

export function ConvertUrlToServiceUserName(url: string): string {
	if (url.includes('github.com')) {
		return url.split('/')[3]
	} else if (url.includes('x.com')) {
		return url.split('/')[3]
	} else if (url.includes('qiita.com')) {
		return url.split('/')[3]
	} else if (url.includes('zenn.dev')) {
		return url.split('/')[3]
	} else if (url.includes('note.com')) {
		return url.split('/')[3]
	}
	return url
}
