import SettingPageMainTitle from '@/components/ui/text/setting-page/mainTitle'
import Link from 'next/link'
import { RiArrowRightSLine, RiLink, RiPencilLine, RiUserLine } from 'react-icons/ri'

/**
 * @package
 */
export default function ProfileSettingPresentation() {
	const settingItems = [
		{
			icon: <RiUserLine />,
			title: '基本設定',
			description: 'プロフィール画像、名前、組織などの基本情報を設定',
			href: '/account/setting/profile/basic'
		},
		{
			icon: <RiPencilLine />,
			title: '自己紹介',
			description: 'あなたの自己紹介文を設定',
			href: '/account/setting/profile/introduction'
		},
		{
			icon: <RiLink />,
			title: 'リンク設定',
			description: 'SNSやWebサイトなどの外部リンクを設定',
			href: '/account/setting/profile/link'
		}
	]

	return (
		<>
			<SettingPageMainTitle icon={<RiUserLine />}>プロフィール設定</SettingPageMainTitle>
			<div className="space-y-3">
				{settingItems.map((item) => (
					<Link key={item.href} href={item.href}>
						<div className="bg-white border border-border rounded-lg p-4 hover:bg-gray-50 transition-colors cursor-pointer">
							<div className="flex items-center justify-between">
								<div className="flex items-center gap-3">
									<div className="text-primary text-xl">{item.icon}</div>
									<div>
										<h3 className="font-medium text-foreground">{item.title}</h3>
										<p className="text-sm text-subtext">{item.description}</p>
									</div>
								</div>
								<RiArrowRightSLine className="text-subtext text-xl" />
							</div>
						</div>
					</Link>
				))}
			</div>
		</>
	)
}
