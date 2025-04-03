import ProfileContainer from '@/app/_containers/profile/container'

const DashboardPage = async (props: { params: Promise<{ userId: string }> }) => {
	const { userId } = await props.params

	return (
		<>
			<ProfileContainer userId={userId}></ProfileContainer>
		</>
	)
}

export default DashboardPage
