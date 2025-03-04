import UserLoginContainer from "@/app/_containers/userLogin/container";
import HeadContent from "@/components/layout/headContent";

const LoginPage = () => {
  return (
    <>
      <HeadContent title="ログイン" des="DevPortは全てのエンジニアのためのポートフォリオサイトです。学生から社会人まで、幅広い層の方にご利用いただけます。" />
      <UserLoginContainer></UserLoginContainer>
    </>
  )
}

export default LoginPage;
