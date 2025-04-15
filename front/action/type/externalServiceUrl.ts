export type ExternalServiceUrlType = {
	service_type: 'qiita' | 'zenn' | 'note' | 'x' | 'github' | 'other'
	service_id: number
	url: string
}
