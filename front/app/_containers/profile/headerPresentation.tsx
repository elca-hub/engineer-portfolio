import DPHeader from '@/components/layout/header'

type Props = {
	userIconPath: string
}

export default function HeaderPresentation({ userIconPath }: Props) {
	return <DPHeader userIconPath={userIconPath}></DPHeader>
}
