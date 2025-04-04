import type { NextConfig } from 'next'

const nextConfig: NextConfig = {
	experimental: {
		optimizePackageImports: ['@/components/ui'],
		serverActions: { bodySizeLimit: '50mb' },
	},
	images: {
		remotePatterns: [
			{
				protocol: 'http',
				hostname: 'minio',
				port: '9000',
			},
		],
	},
}

export default nextConfig
