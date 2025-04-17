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

export function ConvertToExternalServiceUrl(serviceType: number): string {
	switch (serviceType) {
		case 1:
			return `https://github.com/${serviceType}`
		case 2:
			return `https://x.com/${serviceType}`
		case 3:
			return `https://qiita.com/${serviceType}`
		case 4:
			return `https://zenn.dev/${serviceType}`
		case 5:
			return `https://note.com/${serviceType}`
		default:
			return ''
	}
}
