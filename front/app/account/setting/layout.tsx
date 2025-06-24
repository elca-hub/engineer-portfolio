'use server'

import HeadContent from '@/components/layout/headContent'
import DPHeader from '@/components/layout/header'
import TextWithIcon from '@/components/ui/text/textWithIcon'
import { getAuthUser, handleAuthRedirect } from '@/lib/access'
import Link from 'next/link'
import { redirect } from 'next/navigation'
import { RiLink, RiPencilLine, RiUserLine } from 'react-icons/ri'

type Props = {
	children: React.ReactNode
}
export default async function SettingLayout({ children }: Props) {
	const authRedirectPath = await handleAuthRedirect('auth')
	if (authRedirectPath) redirect(authRedirectPath)

	const user = await getAuthUser()

	if (user.errors || !user.data) {
		redirect('/')
	}

	return (
		<>
			<HeadContent
				title={'プロフィール編集'}
				des="DevPortは全てのエンジニアのためのポートフォリオサイトです。学生から社会人まで、幅広い層の方にご利用いただけます。"
			/>
			<div className="flex flex-col h-full">
				<DPHeader user={user.data.user} isLogin={true}></DPHeader>
				<div className="flex flex-row h-full">
					<aside className="fixed top-0 z-10 pt-20 left-0 w-1/6 h-full bg-white border-r border-border p-4 flex flex-col gap-2">
						<h3 className="text-lg font-medium mb-3">設定</h3>
						<ProfileEditAsideItem icon={<RiUserLine />} text="プロフィール設定" link="/account/setting/profile">
							<ProfileEditAsideItem icon={<RiUserLine />} text="基本設定" link="/account/setting/profile/basic" />
							<ProfileEditAsideItem icon={<RiPencilLine />} text="自己紹介" link="/account/setting/profile/introduction" />
							<ProfileEditAsideItem icon={<RiLink />} text="リンク設定" link="/account/setting/profile/link" />
						</ProfileEditAsideItem>
					</aside>

					<main className="ml-[16.666667%] w-5/6 h-full bg-white border-r border-border p-4 flex flex-col gap-2">{children}</main>
				</div>
			</div>
		</>
	)
}

function ProfileEditAsideItem({ icon, text, link, children }: { icon: React.ReactNode; text: string; link: string; children?: React.ReactNode }) {
	return (
		<div className="text-md text-subtext hover:text-foreground cursor-pointer animate-in duration-300">
			<Link href={link}>
				<TextWithIcon icon={icon}>{text}</TextWithIcon>
			</Link>
			<div className="ml-4 flex flex-col mt-2">{children}</div>
		</div>
	)
}
