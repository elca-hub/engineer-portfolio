import ProfileLinkSettingContainer from '@/app/_containers/account-setting/profile/link/container'
import HeadContent from '@/components/layout/headContent'

const ProfileLinkSettingPage = async () => {
	return (
		<>
			<HeadContent
				title={'プロフィール編集'}
				des="DevPortは全てのエンジニアのためのポートフォリオサイトです。学生から社会人まで、幅広い層の方にご利用いただけます。"
			/>
			<ProfileLinkSettingContainer></ProfileLinkSettingContainer>
		</>
	)
}

export default ProfileLinkSettingPage
