import TextWithIcon from '@/components/ui/text/textWithIcon'

export default function SettingPageSubTitle({ icon, children }: { icon: React.ReactNode; children: React.ReactNode }) {
	return (
		<h2 className="text-2xl font-medium text-gray-800 mb-2">
			<TextWithIcon icon={icon}>{children}</TextWithIcon>
		</h2>
	)
}
