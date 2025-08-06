import SettingPageMainTitle from '@/components/ui/text/setting-page/mainTitle'
import Link from 'next/link'
import { RiArrowRightSLine, RiPencilLine, RiUserLine } from 'react-icons/ri'

/**
 * @package
 */
export default function SkillsAndCertificationsSettingPresentation() {
	const settingItems = [
		{
			icon: <RiUserLine />,
			title: 'スキルの管理',
			description: 'プログラミングスキルを公開して他のユーザにアピールしましょう！',
			href: '/account/setting/skills-and-certifications/skills',
		},
		{
			icon: <RiPencilLine />,
			title: '資格の管理',
			description: '苦労して取得した資格を公開！',
			href: '/account/setting/skills-and-certifications/certifications',
		},
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
