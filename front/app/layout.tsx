import CalloutGroup from '@/components/layout/calloutGroup'
import '@radix-ui/themes/styles.css'
import { Noto_Sans_JP } from 'next/font/google'
import './globals.css'
import { MainProvider } from './state'

const notoSansJP = Noto_Sans_JP({
	variable: '--font-noto-sans-jp',
	subsets: ['latin'],
})

export default function RootLayout({
	children,
}: Readonly<{
	children: React.ReactNode
}>) {
	return (
		<html lang="ja" suppressHydrationWarning>
			<body className={`${notoSansJP.variable} bg-background antialiased`}>
				<MainProvider>
					{children}
					<CalloutGroup />
				</MainProvider>
			</body>
		</html>
	)
}
