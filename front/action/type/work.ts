import { WorkTagType } from "./workTag"
import { WorkUrlType } from "./workUrl"

export type WorkType = {
  id: string
  title: string
  content: string
  github_repository_url: string
  external_service_urls: WorkUrlType[]
  tags: WorkTagType[]
  is_draft: boolean
  thumbnail_image_url: string | null
  publish_status: PublishStatus
}

export type PublishStatus = 'draft' | 'private' | 'limited' | 'public'
