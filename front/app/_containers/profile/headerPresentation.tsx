import { UserType } from '@/action/type/user'
import DPHeader from '@/components/layout/header'

type Props = {
	user: UserType | undefined
	isLogin: boolean
}

/**
 * @package
 */
export default function HeaderPresentation({ user, isLogin }: Props) {
	return <DPHeader user={user} isLogin={isLogin}></DPHeader>
}
