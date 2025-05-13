'use client'

import Image from 'next/image'
import { CustomComponents } from './types'

export const defaultComponents: CustomComponents = {
	p: ({ children, ...props }) => (
		<p className="mb-4" {...props}>
			{children}
		</p>
	),
	a: ({ children, href, ...props }) => (
		<a href={href} className="text-blue-600 hover:text-blue-800 underline" target="_blank" rel="noopener noreferrer" {...props}>
			{children}
		</a>
	),
	h1: ({ children, ...props }) => (
		<h1 className="text-3xl font-bold mb-4" {...props}>
			{children}
		</h1>
	),
	h2: ({ children, ...props }) => (
		<h2 className="text-2xl font-bold mb-3" {...props}>
			{children}
		</h2>
	),
	h3: ({ children, ...props }) => (
		<h3 className="text-xl font-bold mb-2" {...props}>
			{children}
		</h3>
	),
	ul: ({ children, ...props }) => (
		<ul className="list-disc list-inside mb-4 pl-4 [&>li>ul]:pl-4" {...props}>
			{children}
		</ul>
	),
	ol: ({ children, ...props }) => (
		<ol className="list-decimal list-inside mb-4" {...props}>
			{children}
		</ol>
	),
	li: ({ children, ...props }) => (
		<li className="mb-1" {...props}>
			{children}
		</li>
	),
	blockquote: ({ children, ...props }) => (
		<blockquote className="border-l-4 border-gray-300 pl-4 italic mb-4" {...props}>
			{children}
		</blockquote>
	),
	code: ({ children, ...props }) => (
		<code className="bg-gray-200 rounded px-1 py-0.5" {...props}>
			{children}
		</code>
	),
	pre: ({ children, ...props }) => (
		<pre className="bg-gray-200 rounded p-4 mb-4 overflow-x-auto" {...props}>
			{children}
		</pre>
	),
	img: ({ src, alt }) => (
		<Image
			src={src || ''}
			alt={alt || ''}
			width={400}
			height={250}
			className="select-none pointer-events-none h-[200px] max-w-3xl mx-auto mb-4 rounded-lg object-contain"
			unoptimized
		/>
	),
	hr: () => (
		<div className="relative my-8">
			<div className="absolute inset-0 flex items-center">
				<div className="w-full border-t border-gray-300"></div>
			</div>
			<div className="relative flex justify-center">
				<div className="bg-white px-4">
					<div className="w-4 h-4 border-2 border-gray-300 rounded-full animate-[slide_2s_ease-in-out_infinite]"></div>
				</div>
			</div>
		</div>
	),
}
