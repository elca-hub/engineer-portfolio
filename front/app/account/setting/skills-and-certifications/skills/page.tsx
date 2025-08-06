import SkillsSettingContainer from '@/app/_containers/account-setting/skills-and-certifications/skills/container'
import HeadContent from '@/components/layout/headContent'

const SkillsSettingPage = async () => {
	return (
		<>
			<HeadContent
				title={'プロフィール編集'}
				des="DevPortは全てのエンジニアのためのポートフォリオサイトです。学生から社会人まで、幅広い層の方にご利用いただけます。"
			/>
			<SkillsSettingContainer></SkillsSettingContainer>
		</>
	)
}

export default SkillsSettingPage
