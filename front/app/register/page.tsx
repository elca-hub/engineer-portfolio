import HeadContent from '@/components/layout/headContent'
import UserRegisterContainer from '../_containers/register/container'

const LoginPage = () => {
	return (
		<>
			<HeadContent
				title="新規登録"
				des="DevPortは全てのエンジニアのためのポートフォリオサイトです。学生から社会人まで、幅広い層の方にご利用いただけます。"
			/>
			<UserRegisterContainer></UserRegisterContainer>
		</>
	)
}

export default LoginPage
