import { UserType } from '@/action/type/user'
import uploadImage from '@/action/usecase/user/uploadImage'
import { CalloutContext } from '@/app/state'
import UserImage from '@/components/ui/image/userImage'
import { getSessionToken } from '@/lib/access'
import { useRouter } from 'next/navigation'
import { useContext, useEffect, useState } from 'react'
import { Button, DropZone, FileTrigger } from 'react-aria-components'
import { RiImageAddLine } from 'react-icons/ri'

type Props = {
	user: UserType
}

function UserIconComponent({
	user,
	type,
	setIconFile,
}: {
	user: UserType
	type: 'icon' | 'header'
	setIconFile: (file: { file: File; type: 'icon' | 'header' }) => void
}) {
	const iconClass = type === 'icon' ? 'size-[6rem] md:size-[8rem]' : 'w-full md:h-40 h-[8rem]'

	const iconOnlyClass = type === 'icon' ? 'absolute -bottom-10 left-4 z-10' : 'relative'

	return (
		<div className={`${iconOnlyClass}`}>
			<DropZone
				onDrop={(e) => {
					const targetFile = e.items[0]
					if (targetFile.kind !== 'file') return
					targetFile.getFile().then((f) => {
						setIconFile({ file: f, type: type })
					})
				}}
				className={`mx-auto relative rounded-xl data-[drop-target]:ring-2 data-[drop-target]:ring-primary data-[drop-target]:ring-offset-2 data-[drop-target]:ring-offset-background`}
			>
				<UserImage
					imageType={type}
					fileName={type === 'icon' ? user.icon_name : user.header_icon_name}
					width={type === 'icon' ? 500 : 1000}
					height={type === 'icon' ? 500 : 1000}
					className={`${iconClass} ${type === 'icon' ? 'border-2 border-white shadow-md rounded-xl object-cover' : ''} rounded-xl object-cover brightness-50`}
				/>
				<FileTrigger
					onSelect={(e) => {
						if (!e) return
						const file = e.item(0)
						if (file === null) return
						setIconFile({ file: file, type: type })
					}}
					acceptedFileTypes={['image/*']}
				>
					<Button
						className={`absolute top-1/2 ${type === 'icon' ? 'left-1/2 -translate-x-1/2' : 'right-1/3 md:right-[40%]'} -translate-y-1/2 rounded-full bg-white/80 p-3 shadow-md hover:scale-95 transition-all duration-200`}
					>
						<RiImageAddLine className="text-xl md:text-2xl" />
					</Button>
				</FileTrigger>
			</DropZone>
		</div>
	)
}

export default function ProfileImageTabPresentation({ user }: Props) {
	const [iconFile, setIconFile] = useState<{ file: File; type: 'icon' | 'header' }>()

	const { callout, setCallout } = useContext(CalloutContext)

	const router = useRouter()

	useEffect(() => {
		if (iconFile) {
			const { file, type } = iconFile
			if (file.type !== 'image/png' && file.type !== 'image/jpeg') {
				setCallout([...callout, { content: 'PNGまたはJPEG形式の画像を選択してください', type: 'error' }])
				setIconFile(undefined)
				return
			}

			if (file.size > 1024 * 1024 * 50) {
				setCallout([...callout, { content: '画像のサイズは50MB以下にしてください', type: 'error' }])
				setIconFile(undefined)
				return
			}

			const iconUploadFlow = async () => {
				const token = await getSessionToken()
				if (!token) {
					setCallout([...callout, { content: 'ログインしてください', type: 'error' }])
					return
				}

				const iconUploadRes = type === 'icon' ? await uploadImage(token, { icon: file }) : await uploadImage(token, { headerIcon: file })

				if (iconUploadRes.errors) {
					for (const error of iconUploadRes.errors) {
						setCallout([...callout, { content: error, type: 'error' }])
					}
					return
				}

				if (iconUploadRes.data) {
					setCallout([...callout, { content: '変更しました', type: 'info' }])
					router.refresh()
				}
			}

			iconUploadFlow()

			setIconFile(undefined)
		}
	}, [iconFile])

	return (
		<div className="relative mb-14">
			<UserIconComponent user={user} type="icon" setIconFile={setIconFile} />
			<UserIconComponent user={user} type="header" setIconFile={setIconFile} />
		</div>
	)
}
