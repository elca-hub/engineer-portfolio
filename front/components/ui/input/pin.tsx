import { Dispatch, SetStateAction, useEffect, useRef } from 'react'
import { Input } from 'react-aria-components'

export type PinCodeState = {
	code: string
	id: number
}

export default function Pin(props: { id: number; nowFocusId: number; code: Dispatch<SetStateAction<PinCodeState>> }) {
	const { code } = props

	const inputRef = useRef<HTMLInputElement | null>(null)

	useEffect(() => {
		if (props.nowFocusId === props.id) {
			inputRef.current?.focus()
		}
	}, [props.id, props.nowFocusId])

	return (
		<Input
			className="
      h-12 w-10 rounded
      border border-subtext
      p-2 text-center
      text-xl 
      text-foreground
      transition duration-200 ease-in-out focus:border-primary
      focus:outline-none focus:ring-1 focus:ring-primary"
			type="text"
			placeholder="-"
			onChange={(e) => {
				const value = parseInt(e.target.value)

				if (value >= 0 && value <= 9) {
					code({
						code: value.toString(),
						id: props.id,
					})
				} else {
					e.target.value = ''
					code({
						code: '',
						id: props.id,
					})
				}
			}}
			ref={inputRef}
			autoComplete="off"
			maxLength={1}
		/>
	)
}
