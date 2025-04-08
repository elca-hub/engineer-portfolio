import { ExternalServiceUrlType } from '@/action/type/externalServiceUrl'
import { SkillType } from '@/action/type/skill'

export type UserType = {
	email: string
	user_id: string
	birthday: string
	name: string
	icon_name: string
	header_icon_name: string
	bio_path: string
	skills: SkillType[]
	external_service_url: ExternalServiceUrlType[]
}
