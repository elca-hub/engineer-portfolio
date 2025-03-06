import UserLoginContainer from '@/app/_containers/login/container'
import HeadContent from '@/components/layout/headContent'
import UserLoginPresentation from '../_containers/login/presentation'

const LoginPage = () => {
	return (
		<>
			<HeadContent
				title="ログイン"
				des="DevPortは全てのエンジニアのためのポートフォリオサイトです。学生から社会人まで、幅広い層の方にご利用いただけます。"
			/>
			<UserLoginPresentation></UserLoginPresentation>
		</>
	)
}

export default LoginPage
