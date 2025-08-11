'use client'
import { UserType } from '@/action/type/user'
import DPHeader from '@/components/layout/header'
import DPButton from '@/components/ui/button/button'
import TextWithIcon from '@/components/ui/text/textWithIcon'
import { Link } from 'react-aria-components'
import { RiPencilLine } from 'react-icons/ri'

type Props = {
	user: UserType | undefined
	isLogin: boolean
}

/**
 * @package
 */
export default function HeaderPresentation({ user, isLogin }: Props) {
	return (
		<DPHeader user={user} isLogin={isLogin}>
			{user && (
				<Link href={`/account/setting/works/add`} className="outline-none">
					<DPButton colormode="primary">
						<TextWithIcon icon={<RiPencilLine />}>作品を投稿する</TextWithIcon>
					</DPButton>
				</Link>
			)}
		</DPHeader>
	)
}
