import type { KnowledgePage, KnowledgeParams } from '@/types/consult'
import { request } from '@/utils/request'

// 文章列表
export const getKnowledgePage = (params: KnowledgeParams) =>
  request<KnowledgePage>('patient/home/knowledge', 'GET', params)
