import HeadContent from "@/components/layout/headContent";
import ConfirmEmailContainer from "../_containers/confirm-email/container";

const ConfirmEmailPage = () => {
  return (
    <>
      <HeadContent title="メール認証" des="DevPortは全てのエンジニアのためのポートフォリオサイトです。学生から社会人まで、幅広い層の方にご利用いただけます。" />
      <ConfirmEmailContainer></ConfirmEmailContainer>
    </>
  )
}

export default ConfirmEmailPage;
