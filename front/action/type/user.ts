import { CertificationType } from "./certification"
import { SkillType } from "./skill"

export type UserType = {
	email: string
	user_id: string
	birthday: string
	name: string
	icon_name: string
	header_icon_name: string
	bio: string
	organization_name: string
	occupation_name: string
	place: string
	skills: SkillType[]
	certifications: CertificationType[]
}
