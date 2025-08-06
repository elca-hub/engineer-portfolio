'use client'

import { CertificationType } from '@/action/type/certification'
import { ExternalServiceUrlType } from '@/action/type/externalServiceUrl'
import { SkillType } from '@/action/type/skill'
import { UserType } from '@/action/type/user'
import CustomMarkdown from '@/components/layout/markdown/CustomMarkdown'
import UserImage from '@/components/ui/image/userImage'
import DPButton from '@/components/ui/button/button'
import TextWithIcon from '@/components/ui/text/textWithIcon'
import { ImageNameByServiceType, ServiceTypeToServiceName } from '@/lib/externalServiceUrl'
import Image from 'next/image'
import Link from 'next/link'
import { useState } from 'react'
import { RiBriefcaseLine, RiBuilding2Line, RiCheckboxCircleLine, RiMapPinLine, RiMedal2Line, RiMedalLine, RiArrowDownSLine, RiEditLine, RiChat1Line } from 'react-icons/ri'

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
								<p className="text-green-600">
									<TextWithIcon icon={<RiCheckboxCircleLine />}>あなたのアカウントです！</TextWithIcon>
								</p>
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
				<section className="flex justify-start gap-x-4 gap-y-2 flex-wrap">
					{externalServiceUrls.map((url) => (
						<ProfileLinkComponent url={url.url} serviceType={url.service_type} key={url.id} />
					))}
				</section>

				<section className="mt-10">
					{user.bio ? (
						<>
							<h2 className="text-2xl font-bold mb-6">
								<TextWithIcon icon={<RiChat1Line />}>自己紹介</TextWithIcon>
							</h2>
							<div className="mt-4 bg-white rounded-md py-6 px-10">
								<CustomMarkdown>{user.bio}</CustomMarkdown>
							</div>
						</>
					) : (
						<p className="text-subtext">自己紹介の記載がありません...</p>
					)}
				</section>

				<section className="mt-10">
					<div className="flex justify-between items-center mb-6">
						<h2 className="text-2xl font-bold">
							<TextWithIcon icon={<RiMedal2Line />}>スキル</TextWithIcon>
						</h2>
						{isAuthUser && (
							<Link href="/account/setting/skills-and-certifications/skills">
								<DPButton
									colormode="primary"
									className="flex items-center gap-2 px-3 py-2 text-sm"
								>
									<RiEditLine className="w-4 h-4" />
									編集
								</DPButton>
							</Link>
						)}
					</div>
					<div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4 items-start">
						{user.skills.length === 0 ? (
							<p className="text-subtext">スキルの記載がありません...</p>
						) : (
							user.skills.map((skill) => (
								<ProfileCard item={skill} type="skill" key={skill.id} />
							))
						)}
					</div>
				</section>

				<section className="mt-10">
					<div className="flex justify-between items-center mb-6">
						<h2 className="text-2xl font-bold">
							<TextWithIcon icon={<RiMedalLine />}>資格</TextWithIcon>
						</h2>
						{isAuthUser && (
							<Link href="/account/setting/skills-and-certifications/certifications">
								<DPButton
									colormode="primary"
									className="flex items-center gap-2 px-3 py-2 text-sm"
								>
									<RiEditLine className="w-4 h-4" />
									編集
								</DPButton>
							</Link>
						)}
					</div>
					<div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4 items-start">
						{user.certifications.length === 0 ? (
							<p className="text-subtext">資格の記載がありません...</p>
						) : (
							user.certifications.map((certification) => (
								<ProfileCard item={certification} type="certification" key={certification.id} />
							))
						)}
					</div>
				</section>
			</main>
		</div>
	)
}

function ProfileLinkComponent({ url, serviceType }: { url: string; serviceType: number }) {
	let text = ServiceTypeToServiceName(serviceType) as string
	if (text === 'other') text = url

	return (
		<Link href={url} className="flex items-center flex-row gap-2 text-subtext" target="_blank" rel="noreferrer">
			<Image
				src={`/external_service/${ImageNameByServiceType(serviceType)}`}
				alt={`icon of ${text}`}
				width={16}
				height={16}
				className="size-[16px]"
			/>
			<p>{text}</p>
		</Link>
	)
}

type ProfileCardItem = SkillType | CertificationType
type ProfileCardType = 'skill' | 'certification'

function ProfileCard({ item, type }: { item: ProfileCardItem; type: ProfileCardType }) {
	const [isCommentOpen, setIsCommentOpen] = useState(false)

	const badgeColor = type === 'skill' ? 'bg-blue-100 text-blue-800' : 'bg-green-100 text-green-800'
	const badgeText = type === 'skill' ? 'スキル' : '資格'
	const dateLabel = type === 'skill' ? '学習開始' : '取得'

	return (
		<div className="bg-white rounded-lg shadow-md p-6 hover:shadow-lg transition-shadow">
			<div className="flex justify-between items-start mb-4">
				<span className={`inline-block px-3 py-1 rounded-full text-xs font-medium ${badgeColor}`}>
					{badgeText}
				</span>
			</div>

			<h3 className="text-lg font-bold mb-2 text-gray-900 line-clamp-2">{item.name}</h3>

			<div className="text-sm text-gray-600 mb-4">
				<span className="font-medium">{dateLabel}年月日:</span> {item.year}
			</div>

			<div className="mt-4">
				{item.comment && item.comment.trim() ? (
					<>
						<button
							onClick={() => setIsCommentOpen(!isCommentOpen)}
							className="flex items-center gap-1 text-sm text-blue-600 hover:text-blue-800 font-medium transition-colors"
						>
							<span className={`transition-transform duration-200 ${isCommentOpen ? 'rotate-180' : ''}`}>
								<RiArrowDownSLine className="w-4 h-4" />
							</span>
							{isCommentOpen ? 'コメントを閉じる' : 'コメントを表示'}
						</button>
						<div
							className={`overflow-hidden transition-all duration-300 ease-in-out ${
								isCommentOpen ? 'max-h-96 opacity-100' : 'max-h-0 opacity-0'
							}`}
						>
							<div className="mt-3 p-3 bg-gray-50 rounded-md">
								<CustomMarkdown>{item.comment}</CustomMarkdown>
							</div>
						</div>
					</>
				) : (
					<p className="text-sm text-gray-500">コメントはありません</p>
				)}
			</div>
		</div>
	)
}