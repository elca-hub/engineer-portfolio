'use client'

import { ExternalServiceUrlType } from '@/action/type/externalServiceUrl'
import createExternalServiceUrl from '@/action/usecase/externalServiceUrl/create'
import deleteExternalServiceUrl from '@/action/usecase/externalServiceUrl/delete'
import updateExternalServiceUrl from '@/action/usecase/externalServiceUrl/update'
import { CalloutContext } from '@/app/state'
import InputField from '@/components/layout/input/inputField'
import DPButton from '@/components/ui/button/button'
import { BudouXText } from '@/components/ui/text/budouxText'
import TextWithIcon from '@/components/ui/text/textWithIcon'
import { getSessionToken } from '@/lib/access'
import {
	ConvertToExternalServiceUrl,
	ConvertUrlToServiceUserName,
	ImageNameByServiceType,
	MAX_SERVICE_TYPE_LEN,
	ServiceTypeToServiceName,
	ServiceTypeToString,
} from '@/lib/externalServiceUrl'
import { motion } from 'framer-motion'
import Image from 'next/image'
import { useRouter } from 'next/navigation'
import React, { useContext, useEffect, useState } from 'react'
import { Cell, Checkbox, CheckboxProps, Column, Row, Table, TableBody, TableHeader } from 'react-aria-components'
import { Controller, useFieldArray, useForm } from 'react-hook-form'
import { RiAtLine, RiCheckboxBlankLine, RiCheckboxFill, RiCheckboxIndeterminateLine, RiLink, RiPencilLine } from 'react-icons/ri'
import { useListData } from 'react-stately'

type LinkFormType = {
	values: { url: string }[]
}

type Props = {
	externalServiceUrls: ExternalServiceUrlType[]
}

export default function ProfileLinkSettingPresentation({ externalServiceUrls }: Props) {
	const router = useRouter()

	const { callout, setCallout } = useContext(CalloutContext)

	function makeInitialList(): { id: number; userId: string; url: string; serviceName: string; isDisabledUserId: boolean }[] {
		const res: { id: number; userId: string; url: string; serviceName: string; isDisabledUserId: boolean }[] = []
		for (let i = 0; i <= MAX_SERVICE_TYPE_LEN; i++) {
			const serviceObj = externalServiceUrls.find((val) => val.service_type === i)
			if (serviceObj) {
				res.push({
					id: i,
					userId: i === 0 ? '' : ConvertUrlToServiceUserName(serviceObj.url),
					url: serviceObj.url,
					serviceName: ServiceTypeToServiceName(i),
					isDisabledUserId: i === 0,
				})
			} else {
				res.push({
					id: i,
					userId: '',
					url: '',
					serviceName: ServiceTypeToServiceName(i),
					isDisabledUserId: i === 0,
				})
			}
		}

		return res
	}

	const list = useListData({
		initialItems: makeInitialList(),
	})

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

	const [editTargetServiceId, setEditTargetServiceId] = useState(-1)

	return (
		<>
			<form onSubmit={handleSubmit(() => setIsSubmit(true))}>
				<div className="flex md:flex-row flex-col md:justify-evenly">
					<Table
						onSelectionChange={(key) => {
							const tar = Array.from(key)
							setEditTargetServiceId(tar.length ? Number(tar[0]) : -1)
						}}
						className="flex justify-center items-center flex-col md:w-[30%]"
						selectionMode="single"
						aria-label="external-service-urls"
					>
						<TableHeader>
							<Column>
								<RiPencilLine></RiPencilLine>
							</Column>
							<Column isRowHeader>サービス名</Column>
						</TableHeader>
						<TableBody items={list.items}>
							{(item) => (
								<Row className="data-[dragging]:opacity-30">
									<Cell>
										<MyCheckbox slot="selection"></MyCheckbox>
									</Cell>
									<Cell>
										<div className="text-[1.2rem] flex justify-center items-center gap-2">
											<Image
												src={`/external_service/${ImageNameByServiceType(item.id)}`}
												width={16}
												height={16}
												alt={`logo by ${item.serviceName}`}
												className="size-[1.2rem]"
											></Image>
											{item.serviceName}
										</div>
									</Cell>
								</Row>
							)}
						</TableBody>
					</Table>
					<div className="md:w-[50%] flex mt-4 justify-center md:gap-6 gap-2 flex-col md:flex-row items-center">
						<div>
							{fields.map((field, index) =>
								index === editTargetServiceId ? (
									<div key={index}>
										<Controller
											name={`values.${index}.url`}
											control={control}
											render={({ field, fieldState }) =>
												editTargetServiceId === 0 ? (
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
							{editTargetServiceId === -1 && (
								<>
									<p className="text-subtext">
										<BudouXText text="表からサービスを選択することで、ユーザ名やURLを編集することができます。"></BudouXText>
									</p>
								</>
							)}
						</div>
					</div>
				</div>
				<div className="flex justify-center">
					<DPButton colormode="primary" className="mt-4" type="submit">
						<TextWithIcon icon={<RiLink></RiLink>}>変更する</TextWithIcon>
					</DPButton>
				</div>
			</form>
		</>
	)
}

function MyCheckbox({
	children,
	...props
}: Omit<CheckboxProps, 'children'> & {
	children?: React.ReactNode
}) {
	return (
		<Checkbox {...props}>
			{({ isIndeterminate, isSelected }) => (
				<>
					<div className="checkbox flex items-center justify-center">
						{isIndeterminate ? (
							<RiCheckboxIndeterminateLine className="text-primary size-[1.25rem]" />
						) : isSelected ? (
							<motion.div
								className="flex items-center justify-center"
								initial={{ scale: 0 }}
								animate={{ scale: 1 }}
								transition={{ type: 'spring', stiffness: 500, damping: 15 }}
							>
								<RiCheckboxFill className="text-secondary size-[1.25rem]" />
							</motion.div>
						) : (
							<RiCheckboxBlankLine className="text-subtext size-[1.25rem]" />
						)}
					</div>
					{children}
				</>
			)}
		</Checkbox>
	)
}
