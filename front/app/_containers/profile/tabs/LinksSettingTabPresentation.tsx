import { ExternalServiceUrlType } from '@/action/type/externalServiceUrl'
import createExternalServiceUrl from '@/action/usecase/externalServiceUrl/create'
import updateExternalServiceUrl from '@/action/usecase/externalServiceUrl/update'
import { CalloutContext } from '@/app/state'
import InputField from '@/components/layout/input/inputField'
import DPButton from '@/components/ui/button/button'
import TextWithIcon from '@/components/ui/text/textWithIcon'
import { getSessionToken } from '@/lib/access'
import { ConvertToExternalServiceUrl, ServiceTypeToString, StringToServiceType } from '@/lib/externalServiceUrl'
import Image from 'next/image'
import { useRouter } from 'next/navigation'
import { useContext, useState } from 'react'
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
import { Controller, useForm } from 'react-hook-form'
import { RiArrowDownWideLine, RiLink, RiPencilLine, RiUserLine } from 'react-icons/ri'

type LinkFormType = {
	name: string
	url: string
}

type Props = {
	externalServiceUrls: ExternalServiceUrlType[]
}

export default function LinksSettingTabPresentation({ externalServiceUrls }: Props) {
	const router = useRouter()

	const { callout, setCallout } = useContext(CalloutContext)

	const { control, watch } = useForm<LinkFormType>({
		defaultValues: {
			name: '',
			url: '',
		},
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

	const submitFunction = async () => {
		const userSelectedData = {
			serviceType: StringToServiceType(watch().name),
			url: watch().name === 'other' ? watch().url : ConvertToExternalServiceUrl(StringToServiceType(watch().name)),
		}

		// 既に同じサービスが存在するか確認
		const isExist = externalServiceUrls.some((url) => url.service_type === userSelectedData.serviceType)

		// 既に同じサービスが存在する場合は更新
		const changeFlow = async () => {
			const token = await getSessionToken()

			if (!token) {
				setCallout([...callout, { content: 'ログインしてください', type: 'error' }])
				return
			}

			if (isExist) {
				const updateTarget = externalServiceUrls.find((val) => val.service_type === userSelectedData.serviceType)

				if (updateTarget === undefined) {
					setCallout([...callout, { content: 'データ変換中にエラーが発生しました', type: 'error' }])
					return
				}

				const res = await updateExternalServiceUrl(token, {
					url: userSelectedData.url,
					service_type: userSelectedData.serviceType,
					id: updateTarget.id,
				})

				if (!res) {
					setCallout([...callout, { content: 'データ更新に失敗しました', type: 'error' }])
					return
				}

				setCallout([...callout, { content: `${ServiceTypeToString(res.service_type)}のURLを変更しました!`, type: 'success' }])
				router.refresh()
			} else {
				const res = await createExternalServiceUrl(token, {
					externalServiceUrl: {
						service_type: userSelectedData.serviceType,
						url: userSelectedData.url,
					},
				})

				if (!res) {
					setCallout([...callout, { content: 'データ登録に失敗しました', type: 'error' }])
					return
				}

				setCallout([...callout, { content: `${ServiceTypeToString(res.service_type)}にURLを登録しました!`, type: 'success' }])
				router.refresh()
			}
		}

		changeFlow()
	}

	return (
		<div>
			<div className="flex mt-4 justify-center gap-6">
				<MySelect
					className="flex flex-col gap-y-2 my-4"
					label="サービス"
					items={serviceItems}
					selectedKey={service}
					onSelectionChange={(selected) => setService(selected)}
				>
					{(item) => (
						<MyListBoxItem id={item.name} textValue={item.name}>
							<MyListContent imageName={item.imageName} content={item.content}></MyListContent>
						</MyListBoxItem>
					)}
				</MySelect>
				<div>
					{service !== 'other' ? (
						<Controller
							name="name"
							control={control}
							rules={{
								required: 'ユーザ名を入力してください',
								max: {
									value: 50,
									message: 'ユーザ名は50文字以内で入力してください',
								},
							}}
							render={({ field, fieldState }) => (
								<InputField
									title="ユーザ名"
									type="text"
									field={field}
									fieldState={fieldState}
									isRequired
									icon={<RiUserLine />}
									autoComplete="off"
									helperText="@マークなどは除いてください"
								></InputField>
							)}
						></Controller>
					) : (
						<Controller
							name="url"
							control={control}
							rules={{
								required: 'URLを入力してください',
							}}
							render={({ field, fieldState }) => (
								<InputField title="URL" type="url" field={field} fieldState={fieldState} isRequired icon={<RiLink />} autoComplete="off"></InputField>
							)}
						></Controller>
					)}
				</div>
			</div>
			<div className="flex justify-center">
				<DPButton colormode="primary" className="mt-4" onPress={submitFunction}>
					<TextWithIcon icon={<RiLink></RiLink>}>変更する</TextWithIcon>
				</DPButton>
			</div>
		</div>
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
