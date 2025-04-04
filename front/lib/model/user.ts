import { ExternalServiceUrlType } from '@/lib/model/externalServiceUrl'
import { SkillType } from '@/lib/model/skill'

export type UserType = {
	email: string
	user_id: string
	name: string
	icon_name: string
	header_path: string
	bio_path: string
	skills: SkillType[]
	external_service_url: ExternalServiceUrlType[]
}
