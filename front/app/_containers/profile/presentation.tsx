'use client'

import { ExternalServiceUrlType } from '@/action/type/externalServiceUrl'
import { UserType } from '@/action/type/user'
import CustomMarkdown from '@/components/layout/markdown/CustomMarkdown'
import UserImage from '@/components/ui/image/userImage'
import TextWithIcon from '@/components/ui/text/textWithIcon'
import { ImageNameByServiceType, ServiceTypeToServiceName } from '@/lib/externalServiceUrl'
import Image from 'next/image'
import Link from 'next/link'
import { RiBriefcaseLine, RiBuilding2Line, RiCheckboxCircleLine, RiMapPinLine } from 'react-icons/ri'

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

				<section>
					{user.bio ? (
						<div className="mt-4 bg-white rounded-md py-6 px-10">
							<h2 className="text-2xl font-bold mb-6">自己紹介</h2>
							<CustomMarkdown>{user.bio}</CustomMarkdown>
						</div>
					) : (
						<p className="text-subtext">自己紹介はありません</p>
					)}
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
