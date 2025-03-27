import RegisterContainer from '@/app/_containers/register/container'
import HeadContent from '@/components/layout/headContent'

const RegisterPage = () => {
	return (
		<>
			<HeadContent
				title="新規登録"
				des="DevPortは全てのエンジニアのためのポートフォリオサイトです。学生から社会人まで、幅広い層の方にご利用いただけます。"
			/>
			<RegisterContainer></RegisterContainer>
		</>
	)
}

export default RegisterPage
