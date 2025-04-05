import ProfileContainer from '@/app/_containers/profile/container'
import HeadContent from '@/components/layout/headContent'

const ProfilePage = async (props: { params: Promise<{ userId: string }> }) => {
	const { userId } = await props.params

	return (
		<>
			<HeadContent
				title={`${userId}のプロフィール`}
				des="DevPortは全てのエンジニアのためのポートフォリオサイトです。学生から社会人まで、幅広い層の方にご利用いただけます。"
			/>
			<ProfileContainer userId={userId}></ProfileContainer>
		</>
	)
}

export default ProfilePage
