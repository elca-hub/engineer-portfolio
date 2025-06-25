import ProfileBasicSettingContainer from '@/app/_containers/account-setting/profile/basic/container'
import HeadContent from '@/components/layout/headContent'

const ProfileBasicSettingPage = async () => {
	return (
		<>
			<HeadContent
				title={'プロフィール編集'}
				des="DevPortは全てのエンジニアのためのポートフォリオサイトです。学生から社会人まで、幅広い層の方にご利用いただけます。"
			/>
			<ProfileBasicSettingContainer></ProfileBasicSettingContainer>
		</>
	)
}

export default ProfileBasicSettingPage
