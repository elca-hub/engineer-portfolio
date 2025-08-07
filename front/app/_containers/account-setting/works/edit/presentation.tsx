/**
 * @package
 */

import { WorkType } from '@/action/type/work'
import { UserType } from '@/action/type/user'

type EditWorkPresentationProps = {
	work: WorkType
	user: UserType
}

export default function EditWorkPresentation({ work, user }: EditWorkPresentationProps) {
	return (
		<div className="flex w-full h-full justify-center items-center">
			<p className="text-subtext">左のサイドメニューから設定項目を選択することができます。</p>
		</div>
	)
}
