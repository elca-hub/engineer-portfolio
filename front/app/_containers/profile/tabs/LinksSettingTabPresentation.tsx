import { ExternalServiceUrlType } from '@/action/type/externalServiceUrl'
import createExternalServiceUrl from '@/action/usecase/externalServiceUrl/create'
import deleteExternalServiceUrl from '@/action/usecase/externalServiceUrl/delete'
import updateExternalServiceUrl from '@/action/usecase/externalServiceUrl/update'
import { CalloutContext } from '@/app/state'
import InputField from '@/components/layout/input/inputField'
import DPButton from '@/components/ui/button/button'
import TextWithIcon from '@/components/ui/text/textWithIcon'
import { getSessionToken } from '@/lib/access'
import {
	ConvertToExternalServiceUrl,
	ConvertUrlToServiceUserName,
	MAX_SERVICE_TYPE_LEN,
	ServiceTypeToString,
	StringToServiceType,
} from '@/lib/externalServiceUrl'
import Image from 'next/image'
import { useRouter } from 'next/navigation'
import { useContext, useEffect, useState } from 'react'
import {
	Button,
	FieldError,
	Key,
	Label,
	ListBox,
	ListBoxItem,
	ListBoxItemProps,
	Popover,
	Select,
	SelectProps,
	SelectValue,
	Text,
	ValidationResult,
} from 'react-aria-components'
import { Controller, useFieldArray, useForm } from 'react-hook-form'
import { RiArrowDownWideLine, RiAtLine, RiLink, RiPencilLine } from 'react-icons/ri'

type LinkFormType = {
	values: { url: string }[]
}

type Props = {
	externalServiceUrls: ExternalServiceUrlType[]
}

export default function LinksSettingTabPresentation({ externalServiceUrls }: Props) {
	const router = useRouter()

	const { callout, setCallout } = useContext(CalloutContext)

	function initServiceValue(): { url: string }[] {
		const serviceValue: { url: string }[] = []
		for (let i = 0; i <= MAX_SERVICE_TYPE_LEN; i++) {
			const service = externalServiceUrls.find((val) => val.service_type === i)
			serviceValue[i] = { url: service ? ConvertUrlToServiceUserName(service.url) : '' }
		}

		return serviceValue
	}

	const { control, watch, handleSubmit } = useForm<LinkFormType>({
		defaultValues: {
			values: initServiceValue(),
		},
	})

	const { fields } = useFieldArray<LinkFormType>({
		name: 'values',
		control,
	})

	const serviceItems: {
		imageName: string
		content: string
		name: string
	}[] = [
		{
			imageName: 'github.png',
			content: 'GitHub',
			name: 'github',
		},
		{
			imageName: 'x.png',
			content: 'X',
			name: 'x',
		},
		{
			imageName: 'qiita.png',
			content: 'Qiita',
			name: 'qiita',
		},
		{
			imageName: 'zenn.svg',
			content: 'Zenn',
			name: 'zenn',
		},
		{
			imageName: 'note.png',
			content: 'note',
			name: 'note',
		},
		{
			imageName: 'other.svg',
			content: 'other',
			name: 'other',
		},
	]

	const [service, setService] = useState<Key>('github')

	const [isSubmit, setIsSubmit] = useState(false)

	useEffect(() => {
		if (!isSubmit) return

		setIsSubmit(false)

		type updateType = { serviceType: number; url: string; id?: string; isDelete: boolean; isExist: boolean }

		const updateData: updateType[] = []

		for (let i = 0; i <= MAX_SERVICE_TYPE_LEN; i++) {
			const url = ServiceTypeToString(i) === 'other' ? watch().values[i].url : ConvertToExternalServiceUrl(i, watch().values[i].url)

			const isExist = externalServiceUrls.some((url) => url.service_type === i)
			const isDelete = url === ''

			if (isDelete && !isExist) continue

			let id = undefined

			if (isExist) {
				const tar = externalServiceUrls.find((val) => val.service_type === i)

				if (tar === undefined) {
					setCallout([...callout, { content: 'データ変換中にエラーが発生しました', type: 'error' }])
					return
				}

				id = tar.id
			}

			updateData[i] = {
				serviceType: i,
				url,
				isExist,
				id,
				isDelete,
			}
		}

		const changeFlow = async (updateData: updateType[]) => {
			const token = await getSessionToken()

			if (!token) {
				setCallout([...callout, { content: 'ログインしてください', type: 'error' }])
				return
			}

			const promiseList = updateData.map(async (item) => {
				let errors: string[] | undefined
				if (item.isExist && item.id) {
					if (item.isDelete) {
						const res = await deleteExternalServiceUrl(token, item.id)
						errors = res.errors
					} else {
						const res = await updateExternalServiceUrl(token, {
							id: item.id,
							service_type: item.serviceType,
							url: item.url,
						})
						errors = res.errors
					}
				} else {
					const res = await createExternalServiceUrl(token, item.serviceType, item.url)
					errors = res.errors
				}

				return errors
			})

			const resList = await Promise.all(promiseList)

			for (const err of resList) {
				if (err) {
					console.error(err)
				}
			}

			if (resList.filter((err) => !!err).length > 0) {
				setCallout([...callout, { content: 'エラーが発生しました。再度試してください。', type: 'error' }])
			} else {
				setCallout([...callout, { content: '外部サービスリンクのデータ更新に成功しました!', type: 'success' }])
			}

			router.refresh()
		}

		changeFlow(updateData)
	}, [isSubmit])

	return (
		<form onSubmit={handleSubmit(() => setIsSubmit(true))}>
			<div className="flex mt-4 justify-center md:gap-6 gap-2 flex-col md:flex-row items-center md:items-start">
				<MySelect
					className="flex flex-col gap-y-2 my-4"
					label="サービス"
					items={serviceItems}
					selectedKey={service}
					onSelectionChange={(service) => setService(service)}
				>
					{(item) => (
						<MyListBoxItem id={item.name} textValue={item.name}>
							<MyListContent imageName={item.imageName} content={item.content}></MyListContent>
						</MyListBoxItem>
					)}
				</MySelect>
				<div>
					{fields.map((field, index) =>
						index === StringToServiceType(service.toString()) ? (
							<div key={index}>
								<Controller
									name={`values.${index}.url`}
									control={control}
									render={({ field, fieldState }) =>
										service === 'other' ? (
											<InputField title="URL" type="url" field={field} fieldState={fieldState} icon={<RiLink />} helperText=""></InputField>
										) : (
											<InputField
												title="ユーザ名"
												type="text"
												field={field}
												fieldState={fieldState}
												helperText="@マークなどは削除してください"
												icon={<RiAtLine />}
											></InputField>
										)
									}
								></Controller>
							</div>
						) : (
							<div key={index}></div>
						),
					)}
				</div>
			</div>
			<div className="flex justify-center">
				<DPButton colormode="primary" className="mt-4" type="submit">
					<TextWithIcon icon={<RiLink></RiLink>}>変更する</TextWithIcon>
				</DPButton>
			</div>
		</form>
	)
}

function MyListBoxItem(props: ListBoxItemProps) {
	return (
		<ListBoxItem
			{...props}
			className="cursor-pointer group flex w-full items-center rounded-md px-3 py-2 box-border outline-none cursor-default text-foreground focus:bg-primary focus:text-white"
		/>
	)
}

interface MySelectProps<T extends object> extends Omit<SelectProps<T>, 'children'> {
	label?: string
	description?: string
	errorMessage?: string | ((validation: ValidationResult) => string)
	items?: Iterable<T>
	children: React.ReactNode | ((item: T) => React.ReactNode)
}

function MySelect<T extends object>({ label, description, errorMessage, children, items, ...props }: MySelectProps<T>) {
	return (
		<Select {...props}>
			<Label className="text-gray-700">
				<TextWithIcon icon={<RiPencilLine></RiPencilLine>}>{label}</TextWithIcon>
			</Label>
			<Button className="h-fit w-[10rem] flex flex-row items-center justify-between gap-x-2 rounded border border-subtext p-2 text-foreground transition duration-200 ease-in-out focus:border-primary focus:outline-none focus:ring-1 focus:ring-primary">
				<SelectValue />
				<RiArrowDownWideLine />
			</Button>
			{description && <Text slot="description">{description}</Text>}
			<FieldError>{errorMessage}</FieldError>
			<Popover className="bg-white ring-1 ring-primary p-2 rounded entering:animate-in entering:fade-in entering:placement-bottom:slide-in-from-top-1 entering:placement-top:slide-in-from-bottom-1 exiting:animate-out exiting:fade-out exiting:placement-bottom:slide-out-to-top-1 exiting:placement-top:slide-out-to-bottom-1 fill-mode-forwards origin-top-left">
				<ListBox items={items} className="outline-none">
					{children}
				</ListBox>
			</Popover>
		</Select>
	)
}

function MyListContent(props: { imageName: string; content: string }) {
	return (
		<span className="flex items-center gap-x-2 p-0.5">
			<Image src={`/external_service/${props.imageName}`} width={16} height={16} alt={`images by ${props.content}`} className="size-[1rem]"></Image>
			<p className="leading-tight text-[1rem]">{props.content}</p>
		</span>
	)
}
