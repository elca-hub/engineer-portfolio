'use client'

import { UserType } from "@/action/type/user"
import DPHeader from "@/components/layout/header"
import SidebarTitle from "@/components/ui/sidebar/sidebarTitle"
import TextWithIcon from "@/components/ui/text/textWithIcon"
import { useContext } from "react"
import { Link } from "react-aria-components"
import { RiSettings4Line, RiUserLine, RiPencilLine, RiLink, RiAwardLine, RiMedal2Line, RiMedalLine } from "react-icons/ri"
import { HeaderButtonContext } from "./state"

type Props = {
	user: UserType
}

export default function SettingHeader({ user }: Props) {
  const contextObj = useContext(HeaderButtonContext)

  const headerButton = contextObj === undefined ? null : contextObj.headerButton

	return <DPHeader isLogin={true} user={user} sidebar={<SettingSidebarContent />}>{headerButton}</DPHeader>
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
			<ProfileEditAsideItem icon={<RiAwardLine />} text="スキル・資格設定" link="/account/setting/skills-and-certifications">
				<ProfileEditAsideItem icon={<RiMedal2Line />} text="スキルの管理" link="/account/setting/skills-and-certifications/skills" />
				<ProfileEditAsideItem icon={<RiMedalLine />} text="資格の管理" link="/account/setting/skills-and-certifications/certifications" />
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
