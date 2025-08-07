import { WorkTagType } from "./workTag"
import { WorkUrlType } from "./workUrl"

export type WorkType = {
  id: string
  title: string
  content: string
  githubRepositoryUrl: string
  externalServiceUrls: WorkUrlType[]
  tags: WorkTagType[]
  isDraft: boolean
}
