import { ComponentPropsWithoutRef } from 'react'

export type CustomComponents = {
	p?: React.ComponentType<ComponentPropsWithoutRef<'p'>>
	a?: React.ComponentType<ComponentPropsWithoutRef<'a'>>
	h1?: React.ComponentType<ComponentPropsWithoutRef<'h1'>>
	h2?: React.ComponentType<ComponentPropsWithoutRef<'h2'>>
	h3?: React.ComponentType<ComponentPropsWithoutRef<'h3'>>
	h4?: React.ComponentType<ComponentPropsWithoutRef<'h4'>>
	h5?: React.ComponentType<ComponentPropsWithoutRef<'h5'>>
	h6?: React.ComponentType<ComponentPropsWithoutRef<'h6'>>
	ul?: React.ComponentType<ComponentPropsWithoutRef<'ul'>>
	ol?: React.ComponentType<ComponentPropsWithoutRef<'ol'>>
	li?: React.ComponentType<ComponentPropsWithoutRef<'li'>>
	blockquote?: React.ComponentType<ComponentPropsWithoutRef<'blockquote'>>
	code?: React.ComponentType<ComponentPropsWithoutRef<'code'>>
	pre?: React.ComponentType<ComponentPropsWithoutRef<'pre'>>
	img?: React.ComponentType<ComponentPropsWithoutRef<'img'>>
	hr?: React.ComponentType<ComponentPropsWithoutRef<'hr'>>
	table?: React.ComponentType<ComponentPropsWithoutRef<'table'>>
	thead?: React.ComponentType<ComponentPropsWithoutRef<'thead'>>
	tbody?: React.ComponentType<ComponentPropsWithoutRef<'tbody'>>
	tr?: React.ComponentType<ComponentPropsWithoutRef<'tr'>>
	th?: React.ComponentType<ComponentPropsWithoutRef<'th'>>
	td?: React.ComponentType<ComponentPropsWithoutRef<'td'>>
}
