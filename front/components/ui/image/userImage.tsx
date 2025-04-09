import { defaultUserHeaderImage, defaultUserIcon } from '@/constants/constant'
import Image, { ImageProps } from 'next/image'

interface Props extends Omit<ImageProps, 'src' | 'alt'> {
	imageType: 'icon' | 'header'
	fileName: string
}

export default function UserImage(props: Props) {
	const { imageType, fileName, ...rest } = props
	const alt = imageType === 'icon' ? 'ユーザアイコン' : 'ユーザヘッダー'

	if (fileName === '' || fileName === '') {
		return <Image {...rest} alt={alt} src={imageType === 'icon' ? defaultUserIcon : defaultUserHeaderImage} />
	}

	const host = process.env.NEXT_PUBLIC_S3_IMAGE_HOST
	const imagePath = `${host}/${imageType}/${fileName}`

	return <Image {...rest} alt={alt} src={imagePath} />
}
