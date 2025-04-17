'use client'

import { ExternalServiceUrlType } from '@/action/type/externalServiceUrl'
import { UserType } from '@/action/type/user'
import BasicSettingTabPresentation from '@/app/_containers/profile/tabs/BasicSettingTabPresentation'
import LinksSettingTabPresentation from '@/app/_containers/profile/tabs/LinksSettingTabPresentation'
import ProfileImageTabPresentation from '@/app/_containers/profile/tabs/ProfileImageTabPresentation'
import DPModal from '@/components/layout/modal'
import DPButton from '@/components/ui/button/button'
import UserImage from '@/components/ui/image/userImage'
import TextWithIcon from '@/components/ui/text/textWithIcon'
import { DialogTrigger, Tab, TabList, TabPanel, Tabs } from 'react-aria-components'
import { RiBriefcaseLine, RiBuilding2Line, RiMapPinLine, RiUserLine } from 'react-icons/ri'

type Props = {
	header: React.ReactNode
	user: UserType
	isAuthUser: boolean
	externalServiceUrls: ExternalServiceUrlType[]
}

/**
 * @package
 */
export default function ProfilePresentation({ header, user, isAuthUser, externalServiceUrls }: Props) {
	const tabItemClassName = `cursor-pointer text-lg text-subtext outline-border outline-primary data-[selected]:border-b data-[selected]:border-primary data-[selected]:font-medium data-[selected]:text-foreground`

	return (
		<div className="flex flex-col h-screen">
			{header}
			<div className="relative">
				<div className="relative w-full h-60">
					<UserImage
						priority={true}
						imageType="header"
						fileName={user.header_icon_name}
						className="object-cover w-full h-full"
						width={1000}
						height={1000}
					/>
				</div>
				<div className="z-2 absolute -bottom-1/4 md:-bottom-1/2 left-2 md:left-16">
					<div className="flex items-end gap-4">
						<UserImage
							imageType="icon"
							fileName={user.icon_name}
							width="171"
							height="171"
							className="size-[120px] md:size-[170px] border-2 border-white shadow-md rounded-xl object-cover"
							priority
						/>
						<div className="flex flex-col mb-1 bg-white/70 rounded-lg p-2 md:bg-transparent">
							{isAuthUser && (
								<DialogTrigger>
									<DPButton colormode="primary">
										<p className="text-sm">
											<TextWithIcon icon={<RiUserLine />}>プロフィール編集</TextWithIcon>
										</p>
									</DPButton>
									<DPModal header={{ title: 'プロフィール編集', icon: <RiUserLine /> }}>
										<Tabs>
											<TabList aria-label="プロフィール編集" className="flex justify-center gap-x-4 mb-4 ">
												<Tab id="image" className={tabItemClassName}>
													画像
												</Tab>
												<Tab id="basic" className={tabItemClassName}>
													基本情報
												</Tab>
												<Tab id="links" className={tabItemClassName}>
													リンク
												</Tab>
											</TabList>
											<TabPanel id="image">
												<div className="relative mb-14">
													<ProfileImageTabPresentation user={user} />
												</div>
											</TabPanel>
											<TabPanel id="basic">
												<BasicSettingTabPresentation user={user} />
											</TabPanel>
											<TabPanel id="links">
												<LinksSettingTabPresentation externalServiceUrls={externalServiceUrls} />
											</TabPanel>
										</Tabs>
									</DPModal>
								</DialogTrigger>
							)}
							<h1 className="text-3xl font-bold tracking-wide text-foreground">{user.name}</h1>
							<p className="text-lg font-medium text-gray-600">@{user.user_id}</p>
						</div>
					</div>
				</div>
			</div>
			<main className="px-4 md:px-16 mt-[80px] md:mt-[130px]">
				<section className="flex justify-start gap-x-4 gap-y-2 flex-wrap">
					{user.organization_name && (
						<p className="text-subtext">
							<TextWithIcon icon={<RiBuilding2Line />}>{user.organization_name}</TextWithIcon>
						</p>
					)}
					{user.occupation_name && (
						<p className="text-subtext">
							<TextWithIcon icon={<RiBriefcaseLine />}>{user.occupation_name}</TextWithIcon>
						</p>
					)}
					{user.place && (
						<p className="text-subtext">
							<TextWithIcon icon={<RiMapPinLine />}>{user.place}</TextWithIcon>
						</p>
					)}
				</section>
			</main>
		</div>
	)
}
