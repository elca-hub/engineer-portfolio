'use client'

import DPButton from '@/components/ui/button/button'
import TextWithIcon from '@/components/ui/text/textWithIcon'
import { motion } from 'framer-motion'
import Link from 'next/link'
import { RiAddLine, RiLoginBoxLine } from 'react-icons/ri'

const LoginButtons = () => {
	return (
		<motion.div className="flex items-center justify-center gap-10 md:gap-28">
			<Link href="/login">
				<DPButton colormode="primary">
					<TextWithIcon icon={<RiLoginBoxLine />}>ログイン</TextWithIcon>
				</DPButton>
			</Link>

			<Link href="/register">
				<DPButton colormode="secondary">
					<TextWithIcon icon={<RiAddLine />}>新規登録</TextWithIcon>
				</DPButton>
			</Link>
		</motion.div>
	)
}

export default LoginButtons
