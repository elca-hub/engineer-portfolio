import DPHeader from '@/components/layout/header'

type Props = {
	userIconName: string
	isLogin: boolean
}

export default function HeaderPresentation({ userIconName, isLogin }: Props) {
	return <DPHeader userIconName={userIconName} isLogin={isLogin}></DPHeader>
}
