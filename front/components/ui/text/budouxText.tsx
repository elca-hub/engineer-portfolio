import { jaModel, Parser } from 'budoux'

export const BudouXText: React.FC<{ text: string }> = ({ text }) => {
	const parser = new Parser(jaModel)
	const segments = parser.parse(text)

	return segments.map((s, index) => (
		<span className="inline-block" key={index}>
			{s}
		</span>
	))
}
