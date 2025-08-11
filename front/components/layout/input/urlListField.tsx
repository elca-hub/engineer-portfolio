import React, { useCallback, useState } from 'react'
import { FieldError, Input, Text, TextField, TextFieldProps } from 'react-aria-components'
import { FieldArrayWithId } from 'react-hook-form'
import DBLabel from '../../ui/input/label'
import { RiCloseLine } from 'react-icons/ri'
import { WorkFormType } from '@/app/_containers/account-setting/works/edit/presentation'
import { WorkUrlType } from '@/action/type/workUrl'

interface InputFieldProps extends TextFieldProps, React.RefAttributes<HTMLDivElement> {
	title: string
  append: (value: WorkUrlType) => void
  remove: (index: number) => void
	isRequired?: boolean
	helperText?: string
	icon?: React.ReactNode
	popoverContent?: string
	isLoading?: boolean
	placeholder?: string
  urls: FieldArrayWithId<WorkFormType, "externalServiceUrls", "id">[]
  maxUrls: number
  validation: (urls: { url: string }[]) => string | null
}

const UrlListField = ({ ...props }: InputFieldProps) => {
  const [tagInput, setTagInput] = useState('')
  const [tagError, setTagError] = useState<string | undefined>(undefined)

  const handleTagRemove = (index: number) => {
    props.remove(index)
  }

  const handleTagAdd = (url: string) => {
    props.append({url: url, title: ""})
  }

  const handleTagKeyDown = useCallback((e: React.KeyboardEvent) => {
    const trimmedTagInput = tagInput.trim()
    const appendTags = [...props.urls, { url: trimmedTagInput }]
    const validationRes = props.validation(appendTags)
    setTagError(validationRes || undefined)

		if (e.key === 'Enter' && !e.nativeEvent.isComposing) {
			e.preventDefault()
			if (trimmedTagInput && validationRes === null) {
				handleTagAdd(trimmedTagInput)
        setTagInput('')
			}
		}
    if (e.key === 'Escape') {
      setTagInput('')
    }
	}, [tagInput])

	return props.isLoading ? (
		<div className="animate-pulse">
			<div className="h-4 bg-gray-200 rounded w-3/4 mb-2"></div>
			<div className="h-4 bg-gray-200 rounded w-1/2 mb-2"></div>
			<div className="h-4 bg-gray-200 rounded w-2/3"></div>
		</div>
	) : (
		<TextField
      value={tagInput}
      onChange={setTagInput}
      isInvalid={tagError !== undefined}
      {...props}
      className={`my-4 flex flex-col gap-1 ${props.className}`}
    >
			<DBLabel
				title={props.title}
				icon={props.icon}
				popoverContent={props.popoverContent}
				isRequired={props.isRequired}
			/>
      <div className="flex flex-wrap gap-2 p-2 border border-subtext rounded min-h-[42px] bg-white">
        {props.urls.filter((url) => url.title !== "").map((url, index) => (
          <div
            key={url.id}
            className="flex items-center gap-1 px-2 py-1 bg-primary text-primary-foreground rounded text-sm"
          >
            <span>{url.title}</span>
            <button
              type="button"
              onClick={() => handleTagRemove(index)}
              className="ml-1 hover:bg-primary-foreground/20 rounded p-0.5"
            >
              <RiCloseLine size={14} />
            </button>
          </div>
        ))}
        {/* タグ入力フィールド */}
        <div className="flex-1 relative">
          <Input
            type="text"
            onKeyDown={handleTagKeyDown}
              placeholder={props.urls.length === 0 ? "URLを入力してEnterキーで追加..." : ""}
            className="w-full bg-white border-none outline-none text-foreground placeholder-subtext"
            disabled={props.urls.length >= props.maxUrls}
          />
        </div>
      </div>
			<Text slot="description" className="text-sm text-subtext">
				{props.helperText}
			</Text>
      <FieldError className="text-red-500">{tagError}</FieldError>
		</TextField>
	)
}

export default UrlListField
