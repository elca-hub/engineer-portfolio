import Image, { ImageProps } from 'next/image'

interface Props extends Omit<ImageProps, 'src' | 'alt'> {
	imageType: 'icon' | 'header'
	fileName: string
}

export default function UserIconImage(props: Props) {
	const { imageType, fileName, ...rest } = props
	const host = process.env.NEXT_PUBLIC_S3_IMAGE_HOST
	const imagePath = `${host}/${imageType}/${fileName}`
	const alt = imageType === 'icon' ? 'ユーザアイコン' : 'ユーザヘッダー'

	return <Image {...rest} alt={alt} src={imagePath} />
}
