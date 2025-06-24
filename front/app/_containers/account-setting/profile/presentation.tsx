'use client'

import { ExternalServiceUrlType } from '@/action/type/externalServiceUrl'
import { UserType } from '@/action/type/user'
import ProfileBasicSettingPresentation from '@/app/_containers/account-setting/profile/basic/basicPresentation'
import ProfileIntroductionSettingPresentation from '@/app/_containers/account-setting/profile/introduction/introductionPresentation'
import ProfileLinkSettingPresentation from '@/app/_containers/account-setting/profile/link/linkPreesntation'
import TextWithIcon from '@/components/ui/text/textWithIcon'
import { RiLink, RiPencilLine, RiUserLine } from 'react-icons/ri'

type Props = {
	user: UserType
	externalServiceUrls: ExternalServiceUrlType[]
}

/**
 * @package
 */
export default function ProfileSettingPresentation({ user, externalServiceUrls }: Props) {
	return (
		<>
			<h1 className="text-3xl font-bold tracking-wide text-foreground">プロフィール設定</h1>
			<h2 className="text-xl font-medium text-gray-600">
				<TextWithIcon icon={<RiUserLine />}>基本設定</TextWithIcon>
			</h2>
			<ProfileBasicSettingPresentation user={user}></ProfileBasicSettingPresentation>
			<h2 className="text-xl font-medium text-gray-600">
				<TextWithIcon icon={<RiPencilLine></RiPencilLine>}>自己紹介設定</TextWithIcon>
			</h2>
			<ProfileIntroductionSettingPresentation user={user}></ProfileIntroductionSettingPresentation>
			<h2 className="text-xl font-medium text-gray-600">
				<TextWithIcon icon={<RiLink />}>リンク設定</TextWithIcon>
			</h2>
			<ProfileLinkSettingPresentation externalServiceUrls={externalServiceUrls}></ProfileLinkSettingPresentation>
		</>
	)
}
