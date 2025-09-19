import PublishWorkContainer from '@/app/_containers/account-setting/works/publish/container'

const PublishWorkPage = async (props: { params: Promise<{ workId: string }> }) => {
	const { workId } = await props.params

	return (
		<>
			<PublishWorkContainer workId={workId} />
		</>
	)
}

export default PublishWorkPage
