import ProfileContainer from '@/app/_containers/profile/container'

const DashboardPage = ({ params }: { params: { userId: string } }) => {
	const { userId } = params

	console.log(userId)

	return (
		<>
			<ProfileContainer></ProfileContainer>
		</>
	)
}

export default DashboardPage
