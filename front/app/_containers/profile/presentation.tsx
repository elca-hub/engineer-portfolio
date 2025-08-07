'use client'

import { CertificationType } from '@/action/type/certification'
import { ExternalServiceUrlType } from '@/action/type/externalServiceUrl'
import { SkillType } from '@/action/type/skill'
import { UserType } from '@/action/type/user'
import CustomMarkdown from '@/components/layout/markdown/CustomMarkdown'
import DPButton from '@/components/ui/button/button'
import UserImage from '@/components/ui/image/userImage'
import TextWithIcon from '@/components/ui/text/textWithIcon'
import formatDate from '@/lib/date'
import { ImageNameByServiceType, ServiceTypeToServiceName } from '@/lib/externalServiceUrl'
import Image from 'next/image'
import Link from 'next/link'
import { useState } from 'react'
import {
	RiArrowDownSLine,
	RiBriefcaseLine,
	RiBuilding2Line,
	RiChat1Line,
	RiCheckboxCircleLine,
	RiEditLine,
	RiMapPinLine,
	RiMedal2Line,
	RiMedalLine,
} from 'react-icons/ri'

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
						<div className="flex flex-col mb-1 bg-white/70 rounded-lg p-2 md:bg-transparent mr-2 md:mr-0">
							{isAuthUser && (
								<p className="text-green-600 text-sm md:text-base">
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
							<div className="mt-4 bg-white rounded-md py-6 px-4 md:px-10">
								<CustomMarkdown>{user.bio}</CustomMarkdown>
							</div>
						</>
					) : (
						<p className="text-subtext">自己紹介の記載がありません...</p>
					)}
				</section>

				<ProfileSection isAuthUser={isAuthUser} user={user} profileCardType="skill" />

				<ProfileSection isAuthUser={isAuthUser} user={user} profileCardType="certification" />
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

type ProfileCardType = 'skill' | 'certification'

function ProfileSection({ isAuthUser, user, profileCardType }: { isAuthUser: boolean; user: UserType; profileCardType: ProfileCardType }) {
	const sectionConfig = {
		skill: {
			title: 'スキル',
			editPath: 'skills',
			item: user.skills,
			icon: <RiMedalLine />,
		},
		certification: {
			title: '資格',
			editPath: 'certifications',
			item: user.certifications,
			icon: <RiMedal2Line />,
		},
	}

	const sectionValue = sectionConfig[profileCardType]

	return (
		<section className="mt-10">
			<div className="flex justify-between items-center mb-6">
				<h2 className="text-2xl font-bold">
					<TextWithIcon icon={sectionValue.icon}>{sectionValue.title}</TextWithIcon>
				</h2>
				{isAuthUser && (
					<Link href={`/account/setting/skills-and-certifications/${sectionValue.editPath}`}>
						<DPButton colormode="primary" className="flex items-center gap-2 px-3 py-2 text-sm">
							<RiEditLine className="w-4 h-4" />
							編集
						</DPButton>
					</Link>
				)}
			</div>
			<div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4 items-start">
				{sectionValue.item.length === 0 ? (
					<p className="text-subtext">{sectionValue.title}の記載がありません...</p>
				) : (
					sectionValue.item.map((item) => <ProfileCard item={item} type={profileCardType} key={item.id} />)
				)}
			</div>
		</section>
	)
}

function ProfileCard({ item, type }: { item: SkillType | CertificationType; type: ProfileCardType }) {
	const [isCommentOpen, setIsCommentOpen] = useState(false)

	const cardConfig = {
		skill: {
			badgeColor: 'bg-blue-100 text-blue-800',
			badgeText: 'スキル',
			dateLabel: '学習開始',
		},
		certification: {
			badgeColor: 'bg-green-100 text-green-800',
			badgeText: '資格',
			dateLabel: '取得',
		},
	}

	const cardValue = cardConfig[type]

	return (
		<div className="bg-white rounded-lg shadow-md p-6 hover:shadow-lg transition-shadow">
			<div className="flex justify-between items-start mb-4">
				<span className={`inline-block px-3 py-1 rounded-full text-xs font-medium ${cardValue.badgeColor}`}>{cardValue.badgeText}</span>
			</div>

			<h3 className="text-lg font-bold mb-2 text-gray-900 line-clamp-2">{item.name}</h3>

			<div className="text-sm text-gray-600 mb-4">
				<span className="font-medium">{cardValue.dateLabel}年月日:</span> {formatDate(item.year)}
			</div>

			<div className="mt-4">
				{item.comment && item.comment.trim() ? (
					<>
						<button
							onClick={() => setIsCommentOpen(!isCommentOpen)}
							className="flex items-center gap-1 text-sm text-blue-600 hover:text-blue-800 font-medium transition-colors"
							aria-expanded={isCommentOpen}
							aria-controls={`comment-${item.id}`}
						>
							<span className={`transition-transform duration-200 ${isCommentOpen ? 'rotate-180' : ''}`}>
								<RiArrowDownSLine className="w-4 h-4" />
							</span>
							{isCommentOpen ? 'コメントを閉じる' : 'コメントを表示'}
						</button>
						<div
							className={`overflow-hidden transition-all duration-300 ease-in-out ${isCommentOpen ? 'max-h-96 opacity-100' : 'max-h-0 opacity-0'}`}
							id={`comment-${item.id}`}
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
