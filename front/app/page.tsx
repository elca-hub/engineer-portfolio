'use client'

import LoginButtons from '@/app/_parts/loginbuttons'
import HomeSection, { ButtonProps } from '@/app/_parts/section'
import HeadContent from '@/components/layout/headContent'
import { motion, useAnimate } from 'motion/react'
import { Pacifico } from 'next/font/google'
import Image from 'next/image'
import { useEffect, useRef, useState } from 'react'
import { RiArrowDownSLine } from 'react-icons/ri'

const pacifico = Pacifico({
	weight: '400',
	subsets: ['latin'],
})

const sections: ButtonProps[] = [
	{
		title: 'What is DevPort?',
		contents: `DevPortは将来エンジニアになりたいと努力している学生、転職を考えている社会人、自身のスキルをアピールしたい人など、エンジニアリングに関わる全ての人のために作られた、ポートフォリオ作成サイトです。`,
		imagePath: '/dummy.png',
		imageAlt: 'dummy',
		imageWidth: 1920,
		imageHeight: 1080,
	},
	{
		title: '自信のある作品を魅力ある形で。',
		contents: `作品一つひとつに物語はあります。そのストーリーを余すことなく書くことができるようにmarkdown記法を用いた自由な文章作成を行えます。QiitaやZenn、noteの記事やGitHubのリポジトリも一緒に紹介できます。`,
		imagePath: '/dummy.png',
		imageAlt: 'dummy',
		imageWidth: 1920,
		imageHeight: 1080,
	},
	{
		title: '今までの努力をアピール。',
		contents: `経歴やスキル、保有している資格など、ポートフォリオに必要な項目を用意しました。物足りないと感じた方は、markdown記法で記述できる自己紹介で思う存分アピールしましょう。`,
		imagePath: '/dummy.png',
		imageAlt: 'dummy',
		imageWidth: 1920,
		imageHeight: 1080,
	},
]

const AnimateTime = 0.8

const TopPage = () => {
	const targetRef = useRef<HTMLDivElement>(null)
	const topViewButtonRef = useRef<HTMLDivElement>(null)
	const [scope, animate] = useAnimate()

	const [isStaticButtonView, setIsStaticButtonView] = useState(false)

	const scrollTarget = () => {
		targetRef.current?.scrollIntoView({ behavior: 'smooth' })
	}

	useEffect(() => {
		const height = window.innerHeight

		const handleScroll = () => {
			setIsStaticButtonView(targetRef.current !== null && window.scrollY >= height)
		}

		window.addEventListener('scroll', handleScroll)

		handleScroll()
	}, [topViewButtonRef, animate])

	return (
		<div>
			<HeadContent
				title="Home"
				des="DevPortは全てのエンジニアのためのポートフォリオサイトです。学生から社会人まで、幅広い層の方にご利用いただけます。"
			/>
			<div className="w-100vw flex h-screen flex-col items-center justify-center">
				<header
					className="
          flex size-full flex-col items-center justify-center gap-20 bg-gradient-to-b from-lightblue to-background
          md:flex-row
        "
				>
					<motion.div initial={{ x: 50, opacity: 0 }} animate={{ x: 0, opacity: 1 }} transition={{ duration: AnimateTime }}>
						<Image src="/logo.webp" alt="DevPort" width={200} height={200} />
					</motion.div>
					<motion.div
						className="text-center"
						initial={{ x: -50, opacity: 0 }}
						animate={{ x: 0, opacity: 1 }}
						transition={{ duration: AnimateTime, delay: 0.4 }}
					>
						<h1 className={`${pacifico.className} mb-8 text-4xl md:text-8xl`}>
							<span className="text-primary">Dev</span>
							<span className="text-secondary">Port</span>
						</h1>
						<p className={`${pacifico.className} text-xl text-subtext md:text-4xl`}>
							<span className="text-primary">Dev</span>eloper <span className="text-secondary">Port</span>folio, Redefined.
						</p>
					</motion.div>
				</header>

				<motion.div
					initial={{ y: 50, opacity: 0 }}
					animate={{ y: 0, opacity: 1 }}
					transition={{ duration: AnimateTime, delay: 1 }}
					className="mb-16"
					ref={topViewButtonRef}
				>
					<LoginButtons />
				</motion.div>

				<motion.div
					className="mb-10 flex cursor-pointer items-center justify-center gap-4 text-xl text-subtext"
					initial={{ y: 50, opacity: 0 }}
					animate={{ y: 0, opacity: 1 }}
					transition={{ duration: AnimateTime, delay: 1.6 }}
					onClick={scrollTarget}
				>
					<RiArrowDownSLine></RiArrowDownSLine>
					<p>What is DevPort?</p>
				</motion.div>
			</div>

			<div
				ref={scope}
				className={`
          fixed z-10 w-full py-4 transition-all duration-300
          ${isStaticButtonView ? 'top-0' : '-top-20'}
        `}
			>
				<div>
					<LoginButtons />
				</div>
			</div>

			<main ref={targetRef}>
				{sections.map((section, index) => (
					<HomeSection {...section} key={`section-${index}`} />
				))}
			</main>
		</div>
	)
}

export default TopPage
