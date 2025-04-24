import { ExternalServiceUrlType } from '@/action/type/externalServiceUrl'
import createExternalServiceUrl from '@/action/usecase/externalServiceUrl/create'
import updateExternalServiceUrl from '@/action/usecase/externalServiceUrl/update'
import { CalloutContext } from '@/app/state'
import InputField from '@/components/layout/input/inputField'
import DPButton from '@/components/ui/button/button'
import TextWithIcon from '@/components/ui/text/textWithIcon'
import { getSessionToken } from '@/lib/access'
import {
	ConvertToExternalServiceUrl,
	ConvertUrlToServiceUserName, MAX_SERVICE_TYPE_LEN,
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
import {
	RiArrowDownWideLine, RiAtLine,
	RiBuilding2Line,
	RiLink,
	RiPencilLine,
	RiUser2Fill,
	RiUser2Line,
	RiUserLine,
} from 'react-icons/ri'
import { before } from 'node:test'

type LinkFormType = {
	values: {url: string}[]
}

type Props = {
	externalServiceUrls: ExternalServiceUrlType[]
}

export default function LinksSettingTabPresentation({ externalServiceUrls }: Props) {
	const router = useRouter()

	const { callout, setCallout } = useContext(CalloutContext)

	function initServiceValue(): {url: string}[] {
		const serviceValue: {url: string}[] = []
		for (let i = 0; i <= MAX_SERVICE_TYPE_LEN; i++) {
			const service = externalServiceUrls.find((val) => val.service_type === i)
			serviceValue[i] = {url: service ? ConvertUrlToServiceUserName(service.url) : ''}
		}

		return serviceValue
	}

	const { control, watch, handleSubmit } = useForm<LinkFormType>({
		defaultValues: {
			values: initServiceValue()
		},
	})

	const { fields } = useFieldArray<LinkFormType>({
		name: "values",
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

		const updateData: {serviceType: number, url: string}[] = []

		for (let i = 0; i <= MAX_SERVICE_TYPE_LEN; i++) {
			updateData[i] = {
				serviceType: i,
				url: ServiceTypeToString(i) === 'other' ? watch().values[i].url : ConvertToExternalServiceUrl(i, watch().values[i].url),
			}
		}

		// 既に同じサービスが存在する場合は更新
		const changeFlow = async (updateData: {serviceType: number, url: string}[]) => {
			const token = await getSessionToken()

			if (!token) {
				setCallout([...callout, { content: 'ログインしてください', type: 'error' }])
				return
			}

			let isError = false
			for (let i = 0; i < updateData.length; i++) {
				const userSelectedData = updateData[i]
				const isExist = externalServiceUrls.some((url) => url.service_type === userSelectedData.serviceType)


				if (isExist) {
					const updateTarget = externalServiceUrls.find((val) => val.service_type === userSelectedData.serviceType)

					if (updateTarget === undefined) {
						setCallout([...callout, { content: 'データ変換中にエラーが発生しました', type: 'error' }])
						isError = true
						break
					}

					const res = await updateExternalServiceUrl(token, {
						url: userSelectedData.url,
						service_type: userSelectedData.serviceType,
						id: updateTarget.id,
					})

					if (!res) {
						setCallout([...callout, { content: 'データ更新に失敗しました', type: 'error' }])
						isError = true
						break
					}
				} else {
					const res = await createExternalServiceUrl(token,userSelectedData.serviceType, userSelectedData.url)

					if (!res) {
						setCallout([...callout, { content: 'データ登録に失敗しました', type: 'error' }])
						isError = true
						break
					}

				}
			}

			if (isError) return

			setCallout([...callout, { content: '外部サービスリンクのデータ更新に成功しました!', type: 'success' }])
			router.refresh()
		}

		changeFlow(updateData)

	}, [isSubmit])

	return (
		<form onSubmit={handleSubmit(() => setIsSubmit(true))}>
			<div className="flex mt-4 justify-center gap-6">
				<MySelect
					className="flex flex-col gap-y-2 my-4"
					label="サービス"
					items={serviceItems}
					selectedKey={service}
					onSelectionChange={service => setService(service)}
				>
					{(item) => (
						<MyListBoxItem id={item.name} textValue={item.name}>
							<MyListContent imageName={item.imageName} content={item.content}></MyListContent>
						</MyListBoxItem>
					)}
				</MySelect>
				<div>
					{fields.map((field, index) => (
						index === StringToServiceType(service.toString()) ?
						<div key={index}>
							<Controller
								name={`values.${index}.url`}
								control={control}
								render={({field, fieldState}) => (
									service === 'other' ?
										<InputField
											title="URL"
											type="url"
											field={field}
											fieldState={fieldState}
											icon={<RiLink />}
											helperText=""
										></InputField>
										:
										<InputField
											title="ユーザ名"
											type="text"
											field={field}
											fieldState={fieldState}
											helperText="@マークなどは削除してください"
											icon={<RiAtLine />}
										></InputField>
								)}
							></Controller>
						</div> : <div key={index}></div>
					))}
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
