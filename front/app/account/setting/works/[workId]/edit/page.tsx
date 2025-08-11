import EditWorkContainer from '@/app/_containers/account-setting/works/edit/container'

const EditWorkPage = async (props: { params: Promise<{ workId: string }> }) => {
	const { workId } = await props.params

	return (
		<>
			<EditWorkContainer workId={workId} />
		</>
	)
}

export default EditWorkPage
