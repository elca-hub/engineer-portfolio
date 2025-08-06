import CertificationsSettingContainer from '@/app/_containers/account-setting/skills-and-certifications/certifications/container'
import HeadContent from '@/components/layout/headContent'

const CertificationsSettingPage = async () => {
	return (
		<>
			<HeadContent
				title={'プロフィール編集'}
				des="DevPortは全てのエンジニアのためのポートフォリオサイトです。学生から社会人まで、幅広い層の方にご利用いただけます。"
			/>
			<CertificationsSettingContainer></CertificationsSettingContainer>
		</>
	)
}

export default CertificationsSettingPage
