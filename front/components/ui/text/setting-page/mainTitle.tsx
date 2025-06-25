import TextWithIcon from '@/components/ui/text/textWithIcon'

export default function SettingPageMainTitle({ icon, children }: { icon: React.ReactNode; children: React.ReactNode }) {
	return (
		<h1 className="text-3xl font-bold tracking-wide text-foreground mb-4">
			<TextWithIcon icon={icon}>{children}</TextWithIcon>
		</h1>
	)
}
