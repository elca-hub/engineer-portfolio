'use client'

import { ExternalServiceUrlType } from '@/action/type/externalServiceUrl'
import { UserType } from '@/action/type/user'
import ProfileBasicSettingPresentation from '@/app/_containers/account-setting/profile/basic/basicPresentation'
import ProfileIntroductionSettingPresentation from '@/app/_containers/account-setting/profile/introduction/introductionPresentation'
import ProfileLinkSettingPresentation from '@/app/_containers/account-setting/profile/link/linkPresentation'
import SettingPageMainTitle from '@/components/ui/text/setting-page/mainTitle'
import SettingPageSubTitle from '@/components/ui/text/setting-page/subTitle'
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
			<SettingPageMainTitle icon={<RiUserLine />}>プロフィール設定</SettingPageMainTitle>

			<SettingPageSubTitle icon={<RiUserLine />}>基本設定</SettingPageSubTitle>
			<ProfileBasicSettingPresentation user={user}></ProfileBasicSettingPresentation>

			<SettingPageSubTitle icon={<RiPencilLine />}>自己紹介設定</SettingPageSubTitle>
			<ProfileIntroductionSettingPresentation user={user}></ProfileIntroductionSettingPresentation>

			<SettingPageSubTitle icon={<RiLink />}>リンク設定</SettingPageSubTitle>
			<ProfileLinkSettingPresentation externalServiceUrls={externalServiceUrls}></ProfileLinkSettingPresentation>
		</>
	)
}
