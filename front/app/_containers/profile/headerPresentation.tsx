import DPHeader from '@/components/layout/header'

type Props = {
	userIconPath: string
	isLogin: boolean
}

export default function HeaderPresentation({ userIconPath, isLogin }: Props) {
	return <DPHeader userIconPath={userIconPath} isLogin={isLogin}></DPHeader>
}
