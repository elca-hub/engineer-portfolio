'use server'

import HeadContent from '@/components/layout/headContent'
import DPHeader from '@/components/layout/header'
import SidebarTitle from '@/components/ui/sidebar/sidebarTitle'
import TextWithIcon from '@/components/ui/text/textWithIcon'
import { getAuthUser, handleAuthRedirect } from '@/lib/access'
import Link from 'next/link'
import { redirect } from 'next/navigation'
import { RiLink, RiPencilLine, RiSettings4Line, RiUserLine } from 'react-icons/ri'

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
			<div className="min-h-screen">
				<DPHeader isLogin={true} user={user.data.user} sidebar={<SettingSidebarContent />}></DPHeader>
				<div className="flex flex-row">
					<main
						className="w-full bg-white border-r border-border p-4 flex flex-col gap-2"
						style={{ minHeight: 'calc(100vh - var(--header-height, 64px))' }}
					>
						{children}
					</main>
				</div>
			</div>
		</>
	)
}

function SettingSidebarContent() {
	return (
		<>
			<SidebarTitle icon={<RiSettings4Line />}>設定</SidebarTitle>
			<ProfileEditAsideItem icon={<RiUserLine />} text="プロフィール設定" link="/account/setting/profile">
				<ProfileEditAsideItem icon={<RiUserLine />} text="基本設定" link="/account/setting/profile/basic" />
				<ProfileEditAsideItem icon={<RiPencilLine />} text="自己紹介" link="/account/setting/profile/introduction" />
				<ProfileEditAsideItem icon={<RiLink />} text="リンク設定" link="/account/setting/profile/link" />
			</ProfileEditAsideItem>
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
