import ProfileIntroductionSettingContainer from '@/app/_containers/account-setting/profile/introduction/container'
import HeadContent from '@/components/layout/headContent'

const ProfileIntroductionSettingPage = async () => {
	return (
		<>
			<HeadContent
				title={'プロフィール編集'}
				des="DevPortは全てのエンジニアのためのポートフォリオサイトです。学生から社会人まで、幅広い層の方にご利用いただけます。"
			/>
			<ProfileIntroductionSettingContainer></ProfileIntroductionSettingContainer>
		</>
	)
}

export default ProfileIntroductionSettingPage
