import TextWithIcon from '@/components/ui/text/textWithIcon'
import { Button, Dialog, Heading, Modal } from 'react-aria-components'
import { RiCloseLine } from 'react-icons/ri'

type Props = {
	header: {
		title: string
		icon?: React.ReactNode
	}
	children: React.ReactNode
}

export default function DPModal({ children, header }: Props) {
	return (
		<Modal className="fixed inset-0 z-50 flex items-center justify-center bg-black/30 entering:animate-in entering:fade-in entering:duration-200 exiting:animate-out exiting:fade-out exiting:duration-200">
			<Dialog className="outline-none bg-white rounded-lg shadow-lg p-6 w-full mx-6 md:w-1/2 md:mx-auto relative">
				<Button slot="close" className="absolute top-4 right-4 outline-block outline-primary">
					<RiCloseLine className="w-8 h-8 hover:scale-90 transition-all duration-200" />
				</Button>
				<Heading slot="title" className="text-2xl font-bold text-foreground mb-4">
					{header.icon ? <TextWithIcon icon={header.icon}>{header.title}</TextWithIcon> : <>{header.title}</>}
				</Heading>

				{children}
			</Dialog>
		</Modal>
	)
}
